// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: type/v1beta1/selector.proto

package v1beta1

import (
	fmt "fmt"
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

// WorkloadSelector specifies the criteria used to determine if a policy can be applied
// to a proxy. The matching criteria includes the metadata associated with a proxy,
// workload instance info such as labels attached to the pod/VM, or any other info
// that the proxy provides to Istio during the initial handshake. If multiple conditions are
// specified, all conditions need to match in order for the workload instance to be
// selected. Currently, only label based selection mechanism is supported.
type WorkloadSelector struct {
	// REQUIRED: One or more labels that indicate a specific set of pods/VMs
	// on which a policy should be applied. The scope of label search is restricted to
	// the configuration namespace in which the resource is present.
	MatchLabels          map[string]string `protobuf:"bytes,1,rep,name=match_labels,json=matchLabels,proto3" json:"match_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WorkloadSelector) Reset()         { *m = WorkloadSelector{} }
func (m *WorkloadSelector) String() string { return proto.CompactTextString(m) }
func (*WorkloadSelector) ProtoMessage()    {}
func (*WorkloadSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd464a999980808d, []int{0}
}
func (m *WorkloadSelector) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkloadSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkloadSelector.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkloadSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadSelector.Merge(m, src)
}
func (m *WorkloadSelector) XXX_Size() int {
	return m.Size()
}
func (m *WorkloadSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadSelector.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadSelector proto.InternalMessageInfo

func (m *WorkloadSelector) GetMatchLabels() map[string]string {
	if m != nil {
		return m.MatchLabels
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkloadSelector)(nil), "istio.type.v1beta1.WorkloadSelector")
	proto.RegisterMapType((map[string]string)(nil), "istio.type.v1beta1.WorkloadSelector.MatchLabelsEntry")
}

func init() { proto.RegisterFile("type/v1beta1/selector.proto", fileDescriptor_dd464a999980808d) }

var fileDescriptor_dd464a999980808d = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0xa9, 0x2c, 0x48,
	0xd5, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x4e, 0xcd, 0x49, 0x4d, 0x2e, 0xc9,
	0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xca, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7, 0x03,
	0x29, 0xd1, 0x83, 0x2a, 0x51, 0x5a, 0xc3, 0xc8, 0x25, 0x10, 0x9e, 0x5f, 0x94, 0x9d, 0x93, 0x9f,
	0x98, 0x12, 0x0c, 0x55, 0x2e, 0x14, 0xc1, 0xc5, 0x93, 0x9b, 0x58, 0x92, 0x9c, 0x11, 0x9f, 0x93,
	0x98, 0x94, 0x9a, 0x53, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0x64, 0xaa, 0x87, 0xa9, 0x5f,
	0x0f, 0x5d, 0xaf, 0x9e, 0x2f, 0x48, 0xa3, 0x0f, 0x58, 0x9f, 0x6b, 0x5e, 0x49, 0x51, 0x65, 0x10,
	0x77, 0x2e, 0x42, 0x44, 0xca, 0x8e, 0x4b, 0x00, 0x5d, 0x81, 0x90, 0x00, 0x17, 0x73, 0x76, 0x6a,
	0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98,
	0x53, 0x9a, 0x2a, 0xc1, 0x04, 0x16, 0x83, 0x70, 0xac, 0x98, 0x2c, 0x18, 0x9d, 0xd4, 0x4f, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x28, 0x49, 0x88, 0x83, 0x32,
	0xf3, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x91, 0xbd, 0x9e, 0xc4, 0x06, 0xf6, 0xb2, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xbd, 0x9c, 0xce, 0x50, 0x11, 0x01, 0x00, 0x00,
}

func (m *WorkloadSelector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkloadSelector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkloadSelector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.MatchLabels) > 0 {
		for k := range m.MatchLabels {
			v := m.MatchLabels[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintSelector(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintSelector(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintSelector(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintSelector(dAtA []byte, offset int, v uint64) int {
	offset -= sovSelector(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WorkloadSelector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MatchLabels) > 0 {
		for k, v := range m.MatchLabels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovSelector(uint64(len(k))) + 1 + len(v) + sovSelector(uint64(len(v)))
			n += mapEntrySize + 1 + sovSelector(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSelector(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSelector(x uint64) (n int) {
	return sovSelector(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WorkloadSelector) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSelector
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
			return fmt.Errorf("proto: WorkloadSelector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkloadSelector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MatchLabels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSelector
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
				return ErrInvalidLengthSelector
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSelector
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MatchLabels == nil {
				m.MatchLabels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSelector
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSelector
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthSelector
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthSelector
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSelector
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthSelector
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthSelector
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipSelector(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthSelector
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MatchLabels[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSelector(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSelector
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSelector
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
func skipSelector(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSelector
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
					return 0, ErrIntOverflowSelector
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
					return 0, ErrIntOverflowSelector
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
				return 0, ErrInvalidLengthSelector
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSelector
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSelector
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
				next, err := skipSelector(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSelector
				}
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
	ErrInvalidLengthSelector = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSelector   = fmt.Errorf("proto: integer overflow")
)
