// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/inverted/span_expression.proto

package inverted

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SetOperator is an operator on sets.
type SetOperator int32

const (
	// None is used in an expression node with no children.
	None SetOperator = 0
	// SetUnion unions the children.
	SetUnion SetOperator = 1
	// SetIntersection intersects the children.
	SetIntersection SetOperator = 2
)

var SetOperator_name = map[int32]string{
	0: "None",
	1: "SetUnion",
	2: "SetIntersection",
}

var SetOperator_value = map[string]int32{
	"None":            0,
	"SetUnion":        1,
	"SetIntersection": 2,
}

func (x SetOperator) String() string {
	return proto.EnumName(SetOperator_name, int32(x))
}

func (SetOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3007edc9c4a12a18, []int{0}
}

// SpanExpressionProto is a proto representation of an inverted.Expression
// tree consisting only of SpanExpressions. It is intended for use in
// expression execution.
type SpanExpressionProto struct {
	SpansToRead []SpanExpressionProto_Span `protobuf:"bytes,1,rep,name=spans_to_read,json=spansToRead,proto3" json:"spans_to_read"`
	Node        SpanExpressionProto_Node   `protobuf:"bytes,2,opt,name=node,proto3" json:"node"`
}

func (m *SpanExpressionProto) Reset()         { *m = SpanExpressionProto{} }
func (m *SpanExpressionProto) String() string { return proto.CompactTextString(m) }
func (*SpanExpressionProto) ProtoMessage()    {}
func (*SpanExpressionProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3007edc9c4a12a18, []int{0}
}
func (m *SpanExpressionProto) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpanExpressionProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SpanExpressionProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanExpressionProto.Merge(m, src)
}
func (m *SpanExpressionProto) XXX_Size() int {
	return m.Size()
}
func (m *SpanExpressionProto) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanExpressionProto.DiscardUnknown(m)
}

var xxx_messageInfo_SpanExpressionProto proto.InternalMessageInfo

// Span is a span of the inverted index. Represents [start, end).
type SpanExpressionProto_Span struct {
	Start []byte `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   []byte `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (m *SpanExpressionProto_Span) Reset()         { *m = SpanExpressionProto_Span{} }
func (m *SpanExpressionProto_Span) String() string { return proto.CompactTextString(m) }
func (*SpanExpressionProto_Span) ProtoMessage()    {}
func (*SpanExpressionProto_Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_3007edc9c4a12a18, []int{0, 0}
}
func (m *SpanExpressionProto_Span) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpanExpressionProto_Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SpanExpressionProto_Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanExpressionProto_Span.Merge(m, src)
}
func (m *SpanExpressionProto_Span) XXX_Size() int {
	return m.Size()
}
func (m *SpanExpressionProto_Span) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanExpressionProto_Span.DiscardUnknown(m)
}

var xxx_messageInfo_SpanExpressionProto_Span proto.InternalMessageInfo

type SpanExpressionProto_Node struct {
	FactoredUnionSpans []SpanExpressionProto_Span `protobuf:"bytes,1,rep,name=factored_union_spans,json=factoredUnionSpans,proto3" json:"factored_union_spans"`
	Operator           SetOperator                `protobuf:"varint,2,opt,name=operator,proto3,enum=cockroach.sql.inverted.SetOperator" json:"operator,omitempty"`
	Left               *SpanExpressionProto_Node  `protobuf:"bytes,3,opt,name=left,proto3" json:"left,omitempty"`
	Right              *SpanExpressionProto_Node  `protobuf:"bytes,4,opt,name=right,proto3" json:"right,omitempty"`
}

func (m *SpanExpressionProto_Node) Reset()         { *m = SpanExpressionProto_Node{} }
func (m *SpanExpressionProto_Node) String() string { return proto.CompactTextString(m) }
func (*SpanExpressionProto_Node) ProtoMessage()    {}
func (*SpanExpressionProto_Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_3007edc9c4a12a18, []int{0, 1}
}
func (m *SpanExpressionProto_Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpanExpressionProto_Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *SpanExpressionProto_Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanExpressionProto_Node.Merge(m, src)
}
func (m *SpanExpressionProto_Node) XXX_Size() int {
	return m.Size()
}
func (m *SpanExpressionProto_Node) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanExpressionProto_Node.DiscardUnknown(m)
}

var xxx_messageInfo_SpanExpressionProto_Node proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cockroach.sql.inverted.SetOperator", SetOperator_name, SetOperator_value)
	proto.RegisterType((*SpanExpressionProto)(nil), "cockroach.sql.inverted.SpanExpressionProto")
	proto.RegisterType((*SpanExpressionProto_Span)(nil), "cockroach.sql.inverted.SpanExpressionProto.Span")
	proto.RegisterType((*SpanExpressionProto_Node)(nil), "cockroach.sql.inverted.SpanExpressionProto.Node")
}

func init() {
	proto.RegisterFile("sql/inverted/span_expression.proto", fileDescriptor_3007edc9c4a12a18)
}

var fileDescriptor_3007edc9c4a12a18 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0xae, 0x12, 0x31,
	0x14, 0xc6, 0xa7, 0x50, 0xcc, 0xa4, 0xa0, 0x92, 0x42, 0xcc, 0x64, 0x16, 0x95, 0xe0, 0x86, 0xb0,
	0x18, 0x0c, 0x3e, 0x80, 0x86, 0xa8, 0x89, 0x2e, 0xd4, 0x0c, 0xba, 0x61, 0x33, 0xa9, 0x33, 0x07,
	0x98, 0x48, 0xda, 0xa1, 0xad, 0xc6, 0x47, 0x70, 0xe9, 0x3b, 0xf8, 0x00, 0x3e, 0xc2, 0xdd, 0xb2,
	0x64, 0xc9, 0xea, 0xe6, 0xde, 0xe1, 0x45, 0x6e, 0x5a, 0xfe, 0xe4, 0x2e, 0xb8, 0x8b, 0xcb, 0xdd,
	0x9d, 0x9e, 0x7e, 0xdf, 0xaf, 0xdf, 0x39, 0x29, 0xe9, 0xea, 0xe5, 0x62, 0x90, 0x8b, 0x5f, 0xa0,
	0x0c, 0x64, 0x03, 0x5d, 0x70, 0x91, 0xc0, 0xef, 0x42, 0x81, 0xd6, 0xb9, 0x14, 0x51, 0xa1, 0xa4,
	0x91, 0xf4, 0x59, 0x2a, 0xd3, 0x1f, 0x4a, 0xf2, 0x74, 0x1e, 0xe9, 0xe5, 0x22, 0x3a, 0xa8, 0xc3,
	0xf6, 0x4c, 0xce, 0xa4, 0x93, 0x0c, 0x6c, 0xb5, 0x53, 0x77, 0xff, 0x63, 0xd2, 0x1a, 0x17, 0x5c,
	0xbc, 0x3b, 0x62, 0xbe, 0x38, 0xca, 0x84, 0x3c, 0xb6, 0x78, 0x9d, 0x18, 0x99, 0x28, 0xe0, 0x59,
	0x80, 0x3a, 0xd5, 0x5e, 0x7d, 0xf8, 0x32, 0x3a, 0x4d, 0x8f, 0x4e, 0x30, 0x5c, 0x6f, 0x84, 0x57,
	0x97, 0xcf, 0xbd, 0xb8, 0xee, 0x60, 0x5f, 0x65, 0x0c, 0x3c, 0xa3, 0x1f, 0x09, 0x16, 0x32, 0x83,
	0xa0, 0xd2, 0x41, 0xf7, 0x45, 0x7e, 0x92, 0x19, 0xec, 0x91, 0x8e, 0x11, 0x46, 0x04, 0x5b, 0x1d,
	0x6d, 0x93, 0x9a, 0x36, 0x5c, 0x99, 0x00, 0x75, 0x50, 0xaf, 0x11, 0xef, 0x0e, 0xb4, 0x49, 0xaa,
	0x20, 0x32, 0xf7, 0x50, 0x23, 0xb6, 0x65, 0x78, 0x51, 0x21, 0xd8, 0x42, 0xe8, 0x9c, 0xb4, 0xa7,
	0x3c, 0x35, 0x52, 0x41, 0x96, 0xfc, 0x14, 0xb9, 0x14, 0x89, 0x8b, 0xf8, 0xc0, 0x39, 0xe9, 0x81,
	0xf9, 0xcd, 0x22, 0xed, 0x85, 0xa6, 0xaf, 0x89, 0x2f, 0x0b, 0x50, 0xdc, 0x48, 0xe5, 0x92, 0x3c,
	0x19, 0xbe, 0xb8, 0x93, 0x0e, 0xe6, 0xf3, 0x5e, 0x1a, 0x1f, 0x4d, 0xf4, 0x2d, 0xc1, 0x0b, 0x98,
	0x9a, 0xa0, 0x7a, 0xde, 0xbe, 0x62, 0xe7, 0xa6, 0xef, 0x49, 0x4d, 0xe5, 0xb3, 0xb9, 0x09, 0xf0,
	0x99, 0x98, 0x9d, 0xbd, 0xff, 0x86, 0xd4, 0x6f, 0xc5, 0xa4, 0xbe, 0xdd, 0xa7, 0x80, 0xa6, 0x47,
	0x1b, 0xc4, 0x1f, 0x83, 0x71, 0x83, 0x37, 0x11, 0x6d, 0x91, 0xa7, 0x63, 0x30, 0x1f, 0x84, 0x01,
	0xa5, 0x21, 0x35, 0xb6, 0x59, 0x09, 0xf1, 0x9f, 0x7f, 0xcc, 0x1b, 0xf5, 0x57, 0xd7, 0xcc, 0x5b,
	0x95, 0x0c, 0xad, 0x4b, 0x86, 0x36, 0x25, 0x43, 0x57, 0x25, 0x43, 0x7f, 0xb7, 0xcc, 0x5b, 0x6f,
	0x99, 0xb7, 0xd9, 0x32, 0x6f, 0xe2, 0x1f, 0xd2, 0x7c, 0x7f, 0xe4, 0xbe, 0xe9, 0xab, 0x9b, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa8, 0x6e, 0xf1, 0xa7, 0xfa, 0x02, 0x00, 0x00,
}

func (m *SpanExpressionProto) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpanExpressionProto) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpanExpressionProto) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Node.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSpanExpression(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.SpansToRead) > 0 {
		for iNdEx := len(m.SpansToRead) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpansToRead[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSpanExpression(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SpanExpressionProto_Span) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpanExpressionProto_Span) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpanExpressionProto_Span) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.End) > 0 {
		i -= len(m.End)
		copy(dAtA[i:], m.End)
		i = encodeVarintSpanExpression(dAtA, i, uint64(len(m.End)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Start) > 0 {
		i -= len(m.Start)
		copy(dAtA[i:], m.Start)
		i = encodeVarintSpanExpression(dAtA, i, uint64(len(m.Start)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SpanExpressionProto_Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpanExpressionProto_Node) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpanExpressionProto_Node) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Right != nil {
		{
			size, err := m.Right.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSpanExpression(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Left != nil {
		{
			size, err := m.Left.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSpanExpression(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Operator != 0 {
		i = encodeVarintSpanExpression(dAtA, i, uint64(m.Operator))
		i--
		dAtA[i] = 0x10
	}
	if len(m.FactoredUnionSpans) > 0 {
		for iNdEx := len(m.FactoredUnionSpans) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FactoredUnionSpans[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSpanExpression(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintSpanExpression(dAtA []byte, offset int, v uint64) int {
	offset -= sovSpanExpression(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SpanExpressionProto) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SpansToRead) > 0 {
		for _, e := range m.SpansToRead {
			l = e.Size()
			n += 1 + l + sovSpanExpression(uint64(l))
		}
	}
	l = m.Node.Size()
	n += 1 + l + sovSpanExpression(uint64(l))
	return n
}

func (m *SpanExpressionProto_Span) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Start)
	if l > 0 {
		n += 1 + l + sovSpanExpression(uint64(l))
	}
	l = len(m.End)
	if l > 0 {
		n += 1 + l + sovSpanExpression(uint64(l))
	}
	return n
}

func (m *SpanExpressionProto_Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FactoredUnionSpans) > 0 {
		for _, e := range m.FactoredUnionSpans {
			l = e.Size()
			n += 1 + l + sovSpanExpression(uint64(l))
		}
	}
	if m.Operator != 0 {
		n += 1 + sovSpanExpression(uint64(m.Operator))
	}
	if m.Left != nil {
		l = m.Left.Size()
		n += 1 + l + sovSpanExpression(uint64(l))
	}
	if m.Right != nil {
		l = m.Right.Size()
		n += 1 + l + sovSpanExpression(uint64(l))
	}
	return n
}

func sovSpanExpression(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSpanExpression(x uint64) (n int) {
	return sovSpanExpression(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SpanExpressionProto) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpanExpression
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SpanExpressionProto: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpanExpressionProto: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpansToRead", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpanExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSpanExpression
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpanExpression
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpansToRead = append(m.SpansToRead, SpanExpressionProto_Span{})
			if err := m.SpansToRead[len(m.SpansToRead)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpanExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSpanExpression
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpanExpression
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Node.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpanExpression(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpanExpression
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SpanExpressionProto_Span) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpanExpression
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Span: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Span: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpanExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSpanExpression
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpanExpression
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Start = append(m.Start[:0], dAtA[iNdEx:postIndex]...)
			if m.Start == nil {
				m.Start = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpanExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSpanExpression
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpanExpression
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.End = append(m.End[:0], dAtA[iNdEx:postIndex]...)
			if m.End == nil {
				m.End = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpanExpression(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpanExpression
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SpanExpressionProto_Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpanExpression
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FactoredUnionSpans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpanExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSpanExpression
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpanExpression
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FactoredUnionSpans = append(m.FactoredUnionSpans, SpanExpressionProto_Span{})
			if err := m.FactoredUnionSpans[len(m.FactoredUnionSpans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			m.Operator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpanExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Operator |= SetOperator(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Left", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpanExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSpanExpression
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpanExpression
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Left == nil {
				m.Left = &SpanExpressionProto_Node{}
			}
			if err := m.Left.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Right", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpanExpression
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSpanExpression
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSpanExpression
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Right == nil {
				m.Right = &SpanExpressionProto_Node{}
			}
			if err := m.Right.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpanExpression(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpanExpression
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSpanExpression(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpanExpression
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSpanExpression
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSpanExpression
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSpanExpression
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSpanExpression
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSpanExpression
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSpanExpression        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpanExpression          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSpanExpression = fmt.Errorf("proto: unexpected end of group")
)