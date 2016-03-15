// Code generated by protoc-gen-go.
// source: expression.proto
// DO NOT EDIT!

package tipb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ExpressionType int32

const (
	// children count 0.
	// values are in text format.
	ExpressionType_Null      ExpressionType = 0
	ExpressionType_Int64     ExpressionType = 1
	ExpressionType_Float64   ExpressionType = 2
	ExpressionType_String    ExpressionType = 3
	ExpressionType_ColumnRef ExpressionType = 21
	// children count 2.
	ExpressionType_AndAnd ExpressionType = 101
	ExpressionType_OrOr   ExpressionType = 102
	ExpressionType_Plus   ExpressionType = 103
	ExpressionType_Minus  ExpressionType = 104
	ExpressionType_EQ     ExpressionType = 105
	ExpressionType_GT     ExpressionType = 106
	ExpressionType_GE     ExpressionType = 107
	ExpressionType_LT     ExpressionType = 108
	ExpressionType_LE     ExpressionType = 109
)

var ExpressionType_name = map[int32]string{
	0:   "Null",
	1:   "Int64",
	2:   "Float64",
	3:   "String",
	21:  "ColumnRef",
	101: "AndAnd",
	102: "OrOr",
	103: "Plus",
	104: "Minus",
	105: "EQ",
	106: "GT",
	107: "GE",
	108: "LT",
	109: "LE",
}
var ExpressionType_value = map[string]int32{
	"Null":      0,
	"Int64":     1,
	"Float64":   2,
	"String":    3,
	"ColumnRef": 21,
	"AndAnd":    101,
	"OrOr":      102,
	"Plus":      103,
	"Minus":     104,
	"EQ":        105,
	"GT":        106,
	"GE":        107,
	"LT":        108,
	"LE":        109,
}

func (x ExpressionType) Enum() *ExpressionType {
	p := new(ExpressionType)
	*p = x
	return p
}
func (x ExpressionType) String() string {
	return proto.EnumName(ExpressionType_name, int32(x))
}
func (x *ExpressionType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ExpressionType_value, data, "ExpressionType")
	if err != nil {
		return err
	}
	*x = ExpressionType(value)
	return nil
}
func (ExpressionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Evaluators should implement evaluation functions for every expression type.
type Expression struct {
	Tp               *ExpressionType `protobuf:"varint,1,opt,name=tp,enum=tipb.ExpressionType" json:"tp,omitempty"`
	Val              []byte          `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
	Children         []*Expression   `protobuf:"bytes,3,rep,name=children" json:"children,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Expression) Reset()                    { *m = Expression{} }
func (m *Expression) String() string            { return proto.CompactTextString(m) }
func (*Expression) ProtoMessage()               {}
func (*Expression) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Expression) GetTp() ExpressionType {
	if m != nil && m.Tp != nil {
		return *m.Tp
	}
	return ExpressionType_Null
}

func (m *Expression) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *Expression) GetChildren() []*Expression {
	if m != nil {
		return m.Children
	}
	return nil
}

func init() {
	proto.RegisterType((*Expression)(nil), "tipb.Expression")
	proto.RegisterEnum("tipb.ExpressionType", ExpressionType_name, ExpressionType_value)
}

var fileDescriptor1 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x4d, 0x3b, 0xe7, 0xf6, 0xaa, 0xe3, 0x11, 0x1c, 0xf4, 0x58, 0x76, 0x90, 0xe1, 0xa1,
	0x87, 0x21, 0xde, 0x9d, 0x54, 0x11, 0xd4, 0x55, 0xd7, 0xbb, 0x74, 0x6d, 0xb6, 0x45, 0xd3, 0x24,
	0xa4, 0xa9, 0xe8, 0xb7, 0xf1, 0xa3, 0x9a, 0xf4, 0xa0, 0x30, 0x08, 0xfc, 0x5e, 0xf8, 0xbf, 0xfc,
	0xc2, 0x7b, 0x80, 0xec, 0x4b, 0x1b, 0xd6, 0xb6, 0x5c, 0xc9, 0x54, 0x1b, 0x65, 0x15, 0x1d, 0x58,
	0xae, 0x37, 0xb3, 0x37, 0x80, 0xec, 0x2f, 0xa1, 0x09, 0x04, 0x56, 0xc7, 0x24, 0x21, 0xf3, 0xc9,
	0xe2, 0x3c, 0xf5, 0x0d, 0xe9, 0x7f, 0x5a, 0x7c, 0x6b, 0x46, 0x23, 0x08, 0x3f, 0x4b, 0x11, 0x07,
	0xae, 0xe5, 0x94, 0xce, 0x60, 0x54, 0xed, 0xb9, 0xa8, 0x0d, 0x93, 0x71, 0x98, 0x84, 0xf3, 0x68,
	0x81, 0x87, 0x8f, 0x2e, 0x7f, 0x08, 0x4c, 0x0e, 0x1c, 0x23, 0x18, 0x3c, 0x77, 0x42, 0xe0, 0x11,
	0x1d, 0xc3, 0xf1, 0x83, 0xb4, 0xd7, 0x57, 0x48, 0x9c, 0xf8, 0xe4, 0x4e, 0xa8, 0xd2, 0x5f, 0x02,
	0x0a, 0x30, 0x5c, 0x5b, 0xc3, 0xe5, 0x0e, 0x43, 0x7a, 0x06, 0xe3, 0x5b, 0x25, 0xba, 0x46, 0xbe,
	0xb2, 0x2d, 0x4e, 0x7d, 0x74, 0x23, 0x6b, 0x77, 0xb0, 0x17, 0xad, 0xcc, 0xca, 0xe0, 0xd6, 0x57,
	0xb9, 0xe8, 0x5a, 0xdc, 0x79, 0xe5, 0x13, 0x97, 0xae, 0xdc, 0xd3, 0x21, 0x04, 0xd9, 0x0b, 0x72,
	0xcf, 0xfb, 0x02, 0xdf, 0x7b, 0x66, 0xf8, 0xe1, 0xf9, 0x58, 0xa0, 0xe8, 0x99, 0x61, 0xb3, 0xbc,
	0x80, 0x69, 0xa5, 0x9a, 0x54, 0xbb, 0xff, 0xaa, 0x52, 0xbb, 0x09, 0xea, 0x4d, 0x3f, 0xc6, 0x32,
	0x5a, 0x33, 0xc1, 0x2a, 0x9b, 0xfb, 0x7d, 0xe5, 0xe4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x93, 0x32,
	0x6c, 0xe2, 0x44, 0x01, 0x00, 0x00,
}