package localstore

import (
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/juju/errors"
	"github.com/pingcap/tidb/kv"
	"github.com/pingcap/tidb/xapi/tipb"
)

type localClient struct {
	regionInfo []*regionInfo
}

func (c *localClient) Send(req *kv.Request) kv.ResponseIterator {
	it := c.buildRespIterator(req)
	it.run()
	return it
}

func (c *localClient) SupportRequestType(reqType, subType int64) bool {
	return false
}

func (c *localClient) buildRespIterator(req *kv.Request) *respIterator {
	it := &respIterator{
		client:      c,
		concurrency: req.Concurrency,
		taskChan:    make(chan *task, req.Concurrency),
		errChan:     make(chan error, req.Concurrency),
		respChan:    make(chan *regionResponse, req.Concurrency),
	}
	// TODO: Implement
	return it
}

func (c *localClient) updateRegionInfo() {
	c.regionInfo = pd.GetRegionInfo()
}

type localResponseReader struct {
	s []byte
	i int64
}

func (r *localResponseReader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func (r *localResponseReader) Close() error {
	r.i = int64(len(r.s))
	return nil
}

func (r *localResponseReader) rowToBytes(row *tipb.Row) ([]byte, error) {
	bs, err := proto.Marshal(row)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return bs, err
}

type respIterator struct {
	client      *localClient
	reqSent     int
	respGot     int
	concurrency int
	tasks       []*task
	responses   []*regionResponse
	taskChan    chan *task
	respChan    chan *regionResponse
	errChan     chan error
	finished    bool
}

type task struct {
	request *regionRequest
	region  *localRS
}

func (it *respIterator) Next() (resp io.ReadCloser, err error) {
	if it.finished {
		return nil, nil
	}
	var regionResp *regionResponse
	select {
	case regionResp = <-it.respChan:
	case err = <-it.errChan:
	}
	if err != nil {
		it.Close()
		return nil, err
	}
	if len(regionResp.newStartKey) != 0 {
		it.client.updateRegionInfo()
		retryTasks := it.createRetryTasks(regionResp)
		it.tasks = append(it.tasks, retryTasks...)
	}
	if it.reqSent < len(it.tasks) {
		it.taskChan <- it.tasks[it.reqSent]
		it.reqSent++
	}
	it.respGot++
	if it.reqSent == len(it.tasks) && it.respGot == it.reqSent {
		it.Close()
	}
	return &localResponseReader{s: regionResp.data}, nil
}

func (it *respIterator) createRetryTasks(resp *regionResponse) []*task {
	return nil
}

func (it *respIterator) Close() error {
	// Make goroutines quit.
	if it.finished {
		return nil
	}
	close(it.taskChan)
	it.finished = true
	return nil
}

func (it *respIterator) run() {
	for i := 0; i < it.concurrency; i++ {
		go func() {
			for task := range it.taskChan {
				resp, err := task.region.Handle(task.request)
				if err != nil {
					it.errChan <- err
					break
				}
				it.respChan <- resp
			}
		}()
		it.taskChan <- it.tasks[i]
	}
}