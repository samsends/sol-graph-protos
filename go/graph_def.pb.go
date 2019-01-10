// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: graph_def.proto

package sol

/*
	Package Name:
*/

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type GraphDef struct {
	Node     []*NodeDef  `protobuf:"bytes,1,rep,name=node,proto3" json:"node,omitempty"`
	Versions *VersionDef `protobuf:"bytes,2,opt,name=versions,proto3" json:"versions,omitempty"`
}

func (m *GraphDef) Reset()      { *m = GraphDef{} }
func (*GraphDef) ProtoMessage() {}
func (*GraphDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_graph_def_d4c3ce8fb13e3479, []int{0}
}
func (m *GraphDef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GraphDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GraphDef.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GraphDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphDef.Merge(dst, src)
}
func (m *GraphDef) XXX_Size() int {
	return m.Size()
}
func (m *GraphDef) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphDef.DiscardUnknown(m)
}

var xxx_messageInfo_GraphDef proto.InternalMessageInfo

func (m *GraphDef) GetNode() []*NodeDef {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *GraphDef) GetVersions() *VersionDef {
	if m != nil {
		return m.Versions
	}
	return nil
}

func init() {
	proto.RegisterType((*GraphDef)(nil), "sol.GraphDef")
}
func (this *GraphDef) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*GraphDef)
	if !ok {
		that2, ok := that.(GraphDef)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *GraphDef")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *GraphDef but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *GraphDef but is not nil && this == nil")
	}
	if len(this.Node) != len(that1.Node) {
		return fmt.Errorf("Node this(%v) Not Equal that(%v)", len(this.Node), len(that1.Node))
	}
	for i := range this.Node {
		if !this.Node[i].Equal(that1.Node[i]) {
			return fmt.Errorf("Node this[%v](%v) Not Equal that[%v](%v)", i, this.Node[i], i, that1.Node[i])
		}
	}
	if !this.Versions.Equal(that1.Versions) {
		return fmt.Errorf("Versions this(%v) Not Equal that(%v)", this.Versions, that1.Versions)
	}
	return nil
}
func (this *GraphDef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GraphDef)
	if !ok {
		that2, ok := that.(GraphDef)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Node) != len(that1.Node) {
		return false
	}
	for i := range this.Node {
		if !this.Node[i].Equal(that1.Node[i]) {
			return false
		}
	}
	if !this.Versions.Equal(that1.Versions) {
		return false
	}
	return true
}
func (this *GraphDef) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&sol.GraphDef{")
	if this.Node != nil {
		s = append(s, "Node: "+fmt.Sprintf("%#v", this.Node)+",\n")
	}
	if this.Versions != nil {
		s = append(s, "Versions: "+fmt.Sprintf("%#v", this.Versions)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGraphDef(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *GraphDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GraphDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Node) > 0 {
		for _, msg := range m.Node {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGraphDef(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Versions != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGraphDef(dAtA, i, uint64(m.Versions.Size()))
		n1, err := m.Versions.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func encodeVarintGraphDef(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedGraphDef(r randyGraphDef, easy bool) *GraphDef {
	this := &GraphDef{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.Node = make([]*NodeDef, v1)
		for i := 0; i < v1; i++ {
			this.Node[i] = NewPopulatedNodeDef(r, easy)
		}
	}
	if r.Intn(10) != 0 {
		this.Versions = NewPopulatedVersionDef(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyGraphDef interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneGraphDef(r randyGraphDef) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringGraphDef(r randyGraphDef) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneGraphDef(r)
	}
	return string(tmps)
}
func randUnrecognizedGraphDef(r randyGraphDef, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldGraphDef(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldGraphDef(dAtA []byte, r randyGraphDef, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateGraphDef(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateGraphDef(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateGraphDef(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateGraphDef(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateGraphDef(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateGraphDef(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateGraphDef(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *GraphDef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Node) > 0 {
		for _, e := range m.Node {
			l = e.Size()
			n += 1 + l + sovGraphDef(uint64(l))
		}
	}
	if m.Versions != nil {
		l = m.Versions.Size()
		n += 1 + l + sovGraphDef(uint64(l))
	}
	return n
}

func sovGraphDef(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGraphDef(x uint64) (n int) {
	return sovGraphDef(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GraphDef) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GraphDef{`,
		`Node:` + strings.Replace(fmt.Sprintf("%v", this.Node), "NodeDef", "NodeDef", 1) + `,`,
		`Versions:` + strings.Replace(fmt.Sprintf("%v", this.Versions), "VersionDef", "VersionDef", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGraphDef(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GraphDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGraphDef
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GraphDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GraphDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGraphDef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGraphDef
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = append(m.Node, &NodeDef{})
			if err := m.Node[len(m.Node)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Versions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGraphDef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGraphDef
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Versions == nil {
				m.Versions = &VersionDef{}
			}
			if err := m.Versions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGraphDef(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGraphDef
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
func skipGraphDef(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGraphDef
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
					return 0, ErrIntOverflowGraphDef
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGraphDef
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGraphDef
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGraphDef
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGraphDef(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGraphDef = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGraphDef   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("graph_def.proto", fileDescriptor_graph_def_d4c3ce8fb13e3479) }

var fileDescriptor_graph_def_d4c3ce8fb13e3479 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2f, 0x4a, 0x2c,
	0xc8, 0x88, 0x4f, 0x49, 0x4d, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0xce, 0xcf,
	0x91, 0xe2, 0xcb, 0xcb, 0x4f, 0x49, 0x45, 0x08, 0x4a, 0x09, 0x96, 0xa5, 0x16, 0x15, 0x67, 0xe6,
	0xe7, 0x21, 0x09, 0x29, 0xa5, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0xd9, 0x49, 0xa5, 0x69, 0xfa, 0x20,
	0x1e, 0x98, 0x03, 0x66, 0x41, 0xd4, 0x28, 0x45, 0x72, 0x71, 0xb8, 0x83, 0x8c, 0x77, 0x49, 0x4d,
	0x13, 0x52, 0xe0, 0x62, 0x01, 0x19, 0x2a, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xc4, 0xa3, 0x57,
	0x9c, 0x9f, 0xa3, 0xe7, 0x97, 0x9f, 0x92, 0xea, 0x92, 0x9a, 0x16, 0x04, 0x96, 0x11, 0xd2, 0xe6,
	0xe2, 0x80, 0x5a, 0x53, 0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0xc4, 0x0f, 0x56, 0x15, 0x06,
	0x11, 0x04, 0x29, 0x84, 0x2b, 0x70, 0x32, 0xb9, 0xf1, 0x50, 0x8e, 0xe1, 0xc1, 0x43, 0x39, 0xc6,
	0x0f, 0x0f, 0xe5, 0x18, 0x7f, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8, 0xe2, 0x91, 0x1c,
	0xe3, 0x8e, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x90, 0xc4,
	0x06, 0x76, 0x97, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x54, 0xfb, 0x3c, 0xb9, 0xf6, 0x00, 0x00,
	0x00,
}
