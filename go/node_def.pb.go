// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: node_def.proto

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
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

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

type NodeDef struct {
	Name   string                    `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Struct string                    `protobuf:"bytes,2,opt,name=Struct,proto3" json:"Struct,omitempty"`
	Inputs []string                  `protobuf:"bytes,3,rep,name=Inputs,proto3" json:"Inputs,omitempty"`
	Fields map[string]*FieldValueDef `protobuf:"bytes,5,rep,name=Fields,proto3" json:"Fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *NodeDef) Reset()      { *m = NodeDef{} }
func (*NodeDef) ProtoMessage() {}
func (*NodeDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_def_08600b1844377846, []int{0}
}
func (m *NodeDef) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeDef.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *NodeDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeDef.Merge(dst, src)
}
func (m *NodeDef) XXX_Size() int {
	return m.Size()
}
func (m *NodeDef) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeDef.DiscardUnknown(m)
}

var xxx_messageInfo_NodeDef proto.InternalMessageInfo

func (m *NodeDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeDef) GetStruct() string {
	if m != nil {
		return m.Struct
	}
	return ""
}

func (m *NodeDef) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *NodeDef) GetFields() map[string]*FieldValueDef {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeDef)(nil), "sol.NodeDef")
	proto.RegisterMapType((map[string]*FieldValueDef)(nil), "sol.NodeDef.FieldsEntry")
}
func (this *NodeDef) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*NodeDef)
	if !ok {
		that2, ok := that.(NodeDef)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *NodeDef")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NodeDef but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NodeDef but is not nil && this == nil")
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if this.Struct != that1.Struct {
		return fmt.Errorf("Struct this(%v) Not Equal that(%v)", this.Struct, that1.Struct)
	}
	if len(this.Inputs) != len(that1.Inputs) {
		return fmt.Errorf("Inputs this(%v) Not Equal that(%v)", len(this.Inputs), len(that1.Inputs))
	}
	for i := range this.Inputs {
		if this.Inputs[i] != that1.Inputs[i] {
			return fmt.Errorf("Inputs this[%v](%v) Not Equal that[%v](%v)", i, this.Inputs[i], i, that1.Inputs[i])
		}
	}
	if len(this.Fields) != len(that1.Fields) {
		return fmt.Errorf("Fields this(%v) Not Equal that(%v)", len(this.Fields), len(that1.Fields))
	}
	for i := range this.Fields {
		if !this.Fields[i].Equal(that1.Fields[i]) {
			return fmt.Errorf("Fields this[%v](%v) Not Equal that[%v](%v)", i, this.Fields[i], i, that1.Fields[i])
		}
	}
	return nil
}
func (this *NodeDef) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NodeDef)
	if !ok {
		that2, ok := that.(NodeDef)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Struct != that1.Struct {
		return false
	}
	if len(this.Inputs) != len(that1.Inputs) {
		return false
	}
	for i := range this.Inputs {
		if this.Inputs[i] != that1.Inputs[i] {
			return false
		}
	}
	if len(this.Fields) != len(that1.Fields) {
		return false
	}
	for i := range this.Fields {
		if !this.Fields[i].Equal(that1.Fields[i]) {
			return false
		}
	}
	return true
}
func (this *NodeDef) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&sol.NodeDef{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Struct: "+fmt.Sprintf("%#v", this.Struct)+",\n")
	s = append(s, "Inputs: "+fmt.Sprintf("%#v", this.Inputs)+",\n")
	keysForFields := make([]string, 0, len(this.Fields))
	for k, _ := range this.Fields {
		keysForFields = append(keysForFields, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForFields)
	mapStringForFields := "map[string]*FieldValueDef{"
	for _, k := range keysForFields {
		mapStringForFields += fmt.Sprintf("%#v: %#v,", k, this.Fields[k])
	}
	mapStringForFields += "}"
	if this.Fields != nil {
		s = append(s, "Fields: "+mapStringForFields+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringNodeDef(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *NodeDef) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeDef) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNodeDef(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Struct) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNodeDef(dAtA, i, uint64(len(m.Struct)))
		i += copy(dAtA[i:], m.Struct)
	}
	if len(m.Inputs) > 0 {
		for _, s := range m.Inputs {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Fields) > 0 {
		for k, _ := range m.Fields {
			dAtA[i] = 0x2a
			i++
			v := m.Fields[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovNodeDef(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovNodeDef(uint64(len(k))) + msgSize
			i = encodeVarintNodeDef(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintNodeDef(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintNodeDef(dAtA, i, uint64(v.Size()))
				n1, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n1
			}
		}
	}
	return i, nil
}

func encodeVarintNodeDef(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedNodeDef(r randyNodeDef, easy bool) *NodeDef {
	this := &NodeDef{}
	this.Name = string(randStringNodeDef(r))
	this.Struct = string(randStringNodeDef(r))
	v1 := r.Intn(10)
	this.Inputs = make([]string, v1)
	for i := 0; i < v1; i++ {
		this.Inputs[i] = string(randStringNodeDef(r))
	}
	if r.Intn(10) != 0 {
		v2 := r.Intn(10)
		this.Fields = make(map[string]*FieldValueDef)
		for i := 0; i < v2; i++ {
			this.Fields[randStringNodeDef(r)] = NewPopulatedFieldValueDef(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyNodeDef interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneNodeDef(r randyNodeDef) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringNodeDef(r randyNodeDef) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneNodeDef(r)
	}
	return string(tmps)
}
func randUnrecognizedNodeDef(r randyNodeDef, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldNodeDef(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldNodeDef(dAtA []byte, r randyNodeDef, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateNodeDef(dAtA, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		dAtA = encodeVarintPopulateNodeDef(dAtA, uint64(v4))
	case 1:
		dAtA = encodeVarintPopulateNodeDef(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateNodeDef(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateNodeDef(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateNodeDef(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateNodeDef(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *NodeDef) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNodeDef(uint64(l))
	}
	l = len(m.Struct)
	if l > 0 {
		n += 1 + l + sovNodeDef(uint64(l))
	}
	if len(m.Inputs) > 0 {
		for _, s := range m.Inputs {
			l = len(s)
			n += 1 + l + sovNodeDef(uint64(l))
		}
	}
	if len(m.Fields) > 0 {
		for k, v := range m.Fields {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovNodeDef(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovNodeDef(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovNodeDef(uint64(mapEntrySize))
		}
	}
	return n
}

func sovNodeDef(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNodeDef(x uint64) (n int) {
	return sovNodeDef(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *NodeDef) String() string {
	if this == nil {
		return "nil"
	}
	keysForFields := make([]string, 0, len(this.Fields))
	for k, _ := range this.Fields {
		keysForFields = append(keysForFields, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForFields)
	mapStringForFields := "map[string]*FieldValueDef{"
	for _, k := range keysForFields {
		mapStringForFields += fmt.Sprintf("%v: %v,", k, this.Fields[k])
	}
	mapStringForFields += "}"
	s := strings.Join([]string{`&NodeDef{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Struct:` + fmt.Sprintf("%v", this.Struct) + `,`,
		`Inputs:` + fmt.Sprintf("%v", this.Inputs) + `,`,
		`Fields:` + mapStringForFields + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringNodeDef(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *NodeDef) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNodeDef
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
			return fmt.Errorf("proto: NodeDef: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeDef: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeDef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodeDef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Struct", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeDef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodeDef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Struct = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeDef
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNodeDef
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inputs = append(m.Inputs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fields", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNodeDef
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
				return ErrInvalidLengthNodeDef
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fields == nil {
				m.Fields = make(map[string]*FieldValueDef)
			}
			var mapkey string
			var mapvalue *FieldValueDef
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNodeDef
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNodeDef
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthNodeDef
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNodeDef
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthNodeDef
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthNodeDef
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &FieldValueDef{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipNodeDef(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthNodeDef
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Fields[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNodeDef(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNodeDef
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
func skipNodeDef(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNodeDef
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
					return 0, ErrIntOverflowNodeDef
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
					return 0, ErrIntOverflowNodeDef
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
				return 0, ErrInvalidLengthNodeDef
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNodeDef
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
				next, err := skipNodeDef(dAtA[start:])
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
	ErrInvalidLengthNodeDef = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNodeDef   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("node_def.proto", fileDescriptor_node_def_08600b1844377846) }

var fileDescriptor_node_def_08600b1844377846 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcb, 0x4f, 0x49,
	0x8d, 0x4f, 0x49, 0x4d, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0xce, 0xcf, 0x91,
	0x12, 0x4d, 0xcb, 0x4c, 0xcd, 0x49, 0x89, 0x2f, 0x4b, 0xcc, 0x29, 0x45, 0x92, 0x93, 0x52, 0x4a,
	0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xb3, 0x93, 0x4a, 0xd3, 0xf4, 0x41, 0x3c, 0x30, 0x07, 0xcc, 0x82,
	0xa8, 0x51, 0x3a, 0xc7, 0xc8, 0xc5, 0xee, 0x97, 0x9f, 0x92, 0xea, 0x92, 0x9a, 0x26, 0x24, 0xc4,
	0xc5, 0xe2, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x89,
	0x71, 0xb1, 0x05, 0x97, 0x14, 0x95, 0x26, 0x97, 0x48, 0x30, 0x81, 0x45, 0xa1, 0x3c, 0x90, 0xb8,
	0x67, 0x5e, 0x41, 0x69, 0x49, 0xb1, 0x04, 0xb3, 0x02, 0x33, 0x48, 0x1c, 0xc2, 0x13, 0x32, 0xe0,
	0x62, 0x73, 0x03, 0x39, 0xa6, 0x58, 0x82, 0x55, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x42, 0xaf, 0x38,
	0x3f, 0x47, 0x0f, 0x6a, 0x83, 0x1e, 0x44, 0xca, 0x35, 0xaf, 0xa4, 0xa8, 0x32, 0x08, 0xaa, 0x4e,
	0xca, 0x97, 0x8b, 0x1b, 0x49, 0x58, 0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x12, 0xea, 0x06, 0x10,
	0x53, 0x48, 0x83, 0x8b, 0x15, 0xec, 0x33, 0xb0, 0x0b, 0xb8, 0x8d, 0x84, 0xc0, 0x26, 0x82, 0xb5,
	0x84, 0x81, 0x84, 0x5d, 0x52, 0xd3, 0x82, 0x20, 0x0a, 0xac, 0x98, 0x2c, 0x18, 0x9d, 0x4c, 0x6e,
	0x3c, 0x94, 0x63, 0x78, 0xf0, 0x50, 0x8e, 0xf1, 0xc3, 0x43, 0x39, 0xc6, 0x1f, 0x0f, 0xe5, 0x18,
	0x1b, 0x1e, 0xc9, 0x31, 0xae, 0x78, 0x24, 0xc7, 0xb8, 0xe3, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7,
	0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0x43, 0xc3, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xfc, 0x26, 0x18, 0x17, 0x5f, 0x01, 0x00, 0x00,
}
