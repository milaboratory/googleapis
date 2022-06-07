// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/type/quaternion.proto

package _type

import (
	bytes "bytes"
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// A quaternion is defined as the quotient of two directed lines in a
// three-dimensional space or equivalently as the quotient of two Euclidean
// vectors (https://en.wikipedia.org/wiki/Quaternion).
//
// Quaternions are often used in calculations involving three-dimensional
// rotations (https://en.wikipedia.org/wiki/Quaternions_and_spatial_rotation),
// as they provide greater mathematical robustness by avoiding the gimbal lock
// problems that can be encountered when using Euler angles
// (https://en.wikipedia.org/wiki/Gimbal_lock).
//
// Quaternions are generally represented in this form:
//
//     w + xi + yj + zk
//
// where x, y, z, and w are real numbers, and i, j, and k are three imaginary
// numbers.
//
// Our naming choice `(x, y, z, w)` comes from the desire to avoid confusion for
// those interested in the geometric properties of the quaternion in the 3D
// Cartesian space. Other texts often use alternative names or subscripts, such
// as `(a, b, c, d)`, `(1, i, j, k)`, or `(0, 1, 2, 3)`, which are perhaps
// better suited for mathematical interpretations.
//
// To avoid any confusion, as well as to maintain compatibility with a large
// number of software libraries, the quaternions represented using the protocol
// buffer below *must* follow the Hamilton convention, which defines `ij = k`
// (i.e. a right-handed algebra), and therefore:
//
//     i^2 = j^2 = k^2 = ijk = −1
//     ij = −ji = k
//     jk = −kj = i
//     ki = −ik = j
//
// Please DO NOT use this to represent quaternions that follow the JPL
// convention, or any of the other quaternion flavors out there.
//
// Definitions:
//
//   - Quaternion norm (or magnitude): `sqrt(x^2 + y^2 + z^2 + w^2)`.
//   - Unit (or normalized) quaternion: a quaternion whose norm is 1.
//   - Pure quaternion: a quaternion whose scalar component (`w`) is 0.
//   - Rotation quaternion: a unit quaternion used to represent rotation.
//   - Orientation quaternion: a unit quaternion used to represent orientation.
//
// A quaternion can be normalized by dividing it by its norm. The resulting
// quaternion maintains the same direction, but has a norm of 1, i.e. it moves
// on the unit sphere. This is generally necessary for rotation and orientation
// quaternions, to avoid rounding errors:
// https://en.wikipedia.org/wiki/Rotation_formalisms_in_three_dimensions
//
// Note that `(x, y, z, w)` and `(-x, -y, -z, -w)` represent the same rotation,
// but normalization would be even more useful, e.g. for comparison purposes, if
// it would produce a unique representation. It is thus recommended that `w` be
// kept positive, which can be achieved by changing all the signs when `w` is
// negative.
//
type Quaternion struct {
	// The x component.
	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	// The y component.
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	// The z component.
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
	// The scalar component.
	W                    float64  `protobuf:"fixed64,4,opt,name=w,proto3" json:"w,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Quaternion) Reset()      { *m = Quaternion{} }
func (*Quaternion) ProtoMessage() {}
func (*Quaternion) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b847e1bdd83ff5e, []int{0}
}
func (m *Quaternion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Quaternion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Quaternion.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Quaternion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quaternion.Merge(m, src)
}
func (m *Quaternion) XXX_Size() int {
	return m.Size()
}
func (m *Quaternion) XXX_DiscardUnknown() {
	xxx_messageInfo_Quaternion.DiscardUnknown(m)
}

var xxx_messageInfo_Quaternion proto.InternalMessageInfo

func (m *Quaternion) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Quaternion) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Quaternion) GetZ() float64 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *Quaternion) GetW() float64 {
	if m != nil {
		return m.W
	}
	return 0
}

func (*Quaternion) XXX_MessageName() string {
	return "google.type.Quaternion"
}
func init() {
	proto.RegisterType((*Quaternion)(nil), "google.type.Quaternion")
}

func init() { proto.RegisterFile("google/type/quaternion.proto", fileDescriptor_6b847e1bdd83ff5e) }

var fileDescriptor_6b847e1bdd83ff5e = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x2c, 0x4d, 0x2c, 0x49, 0x2d, 0xca, 0xcb,
	0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0xc8, 0xea, 0x81, 0x64, 0x95,
	0x5c, 0xb8, 0xb8, 0x02, 0xe1, 0x0a, 0x84, 0x78, 0xb8, 0x18, 0x2b, 0x24, 0x18, 0x15, 0x18, 0x35,
	0x18, 0x83, 0x18, 0x2b, 0x40, 0xbc, 0x4a, 0x09, 0x26, 0x08, 0xaf, 0x12, 0xc4, 0xab, 0x92, 0x60,
	0x86, 0xf0, 0xaa, 0x40, 0xbc, 0x72, 0x09, 0x16, 0x08, 0xaf, 0xdc, 0xa9, 0x8b, 0xf1, 0xc6, 0x43,
	0x39, 0x86, 0x0f, 0x0f, 0xe5, 0x18, 0x7f, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8, 0xe2,
	0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0xf8,
	0xe2, 0x91, 0x1c, 0xc3, 0x07, 0x90, 0xf8, 0x63, 0x39, 0xc6, 0x13, 0x8f, 0xe5, 0x18, 0xb9, 0xf8,
	0x93, 0xf3, 0x73, 0xf5, 0x90, 0x5c, 0xe2, 0xc4, 0x8f, 0x70, 0x47, 0x00, 0xc8, 0x9d, 0x01, 0x8c,
	0x51, 0xda, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9,
	0xf9, 0xfa, 0x10, 0xf5, 0x89, 0x05, 0x99, 0xc5, 0xfa, 0x48, 0x5e, 0xb4, 0x06, 0x11, 0x3f, 0x18,
	0x19, 0x17, 0x31, 0x31, 0xbb, 0x87, 0x04, 0x24, 0xb1, 0x81, 0xbd, 0x69, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xbd, 0xde, 0x79, 0xb4, 0x06, 0x01, 0x00, 0x00,
}

func (this *Quaternion) Compare(that interface{}) int {
	if that == nil {
		if this == nil {
			return 0
		}
		return 1
	}

	that1, ok := that.(*Quaternion)
	if !ok {
		that2, ok := that.(Quaternion)
		if ok {
			that1 = &that2
		} else {
			return 1
		}
	}
	if that1 == nil {
		if this == nil {
			return 0
		}
		return 1
	} else if this == nil {
		return -1
	}
	if this.X != that1.X {
		if this.X < that1.X {
			return -1
		}
		return 1
	}
	if this.Y != that1.Y {
		if this.Y < that1.Y {
			return -1
		}
		return 1
	}
	if this.Z != that1.Z {
		if this.Z < that1.Z {
			return -1
		}
		return 1
	}
	if this.W != that1.W {
		if this.W < that1.W {
			return -1
		}
		return 1
	}
	if c := bytes.Compare(this.XXX_unrecognized, that1.XXX_unrecognized); c != 0 {
		return c
	}
	return 0
}
func (this *Quaternion) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Quaternion)
	if !ok {
		that2, ok := that.(Quaternion)
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
	if this.X != that1.X {
		return false
	}
	if this.Y != that1.Y {
		return false
	}
	if this.Z != that1.Z {
		return false
	}
	if this.W != that1.W {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Quaternion) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&_type.Quaternion{")
	s = append(s, "X: "+fmt.Sprintf("%#v", this.X)+",\n")
	s = append(s, "Y: "+fmt.Sprintf("%#v", this.Y)+",\n")
	s = append(s, "Z: "+fmt.Sprintf("%#v", this.Z)+",\n")
	s = append(s, "W: "+fmt.Sprintf("%#v", this.W)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringQuaternion(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Quaternion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Quaternion) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Quaternion) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.W != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.W))))
		i--
		dAtA[i] = 0x21
	}
	if m.Z != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Z))))
		i--
		dAtA[i] = 0x19
	}
	if m.Y != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Y))))
		i--
		dAtA[i] = 0x11
	}
	if m.X != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.X))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuaternion(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuaternion(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedQuaternion(r randyQuaternion, easy bool) *Quaternion {
	this := &Quaternion{}
	this.X = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.X *= -1
	}
	this.Y = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Y *= -1
	}
	this.Z = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.Z *= -1
	}
	this.W = float64(r.Float64())
	if r.Intn(2) == 0 {
		this.W *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedQuaternion(r, 5)
	}
	return this
}

type randyQuaternion interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneQuaternion(r randyQuaternion) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringQuaternion(r randyQuaternion) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneQuaternion(r)
	}
	return string(tmps)
}
func randUnrecognizedQuaternion(r randyQuaternion, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldQuaternion(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldQuaternion(dAtA []byte, r randyQuaternion, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateQuaternion(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateQuaternion(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateQuaternion(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateQuaternion(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateQuaternion(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateQuaternion(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateQuaternion(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Quaternion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.X != 0 {
		n += 9
	}
	if m.Y != 0 {
		n += 9
	}
	if m.Z != 0 {
		n += 9
	}
	if m.W != 0 {
		n += 9
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovQuaternion(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuaternion(x uint64) (n int) {
	return sovQuaternion(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Quaternion) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Quaternion{`,
		`X:` + fmt.Sprintf("%v", this.X) + `,`,
		`Y:` + fmt.Sprintf("%v", this.Y) + `,`,
		`Z:` + fmt.Sprintf("%v", this.Z) + `,`,
		`W:` + fmt.Sprintf("%v", this.W) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringQuaternion(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Quaternion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuaternion
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
			return fmt.Errorf("proto: Quaternion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Quaternion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.X = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Y = float64(math.Float64frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Z", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Z = float64(math.Float64frombits(v))
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field W", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.W = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipQuaternion(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuaternion
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuaternion(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuaternion
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
					return 0, ErrIntOverflowQuaternion
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
					return 0, ErrIntOverflowQuaternion
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
				return 0, ErrInvalidLengthQuaternion
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuaternion
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuaternion
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuaternion        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuaternion          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuaternion = fmt.Errorf("proto: unexpected end of group")
)
