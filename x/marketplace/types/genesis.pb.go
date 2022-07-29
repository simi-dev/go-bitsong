// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitsong/marketplace/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the auction module's genesis state
type GenesisState struct {
	Params         Params           `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Auctions       []Auction        `protobuf:"bytes,2,rep,name=auctions,proto3" json:"auctions"`
	Bids           []Bid            `protobuf:"bytes,4,rep,name=bids,proto3" json:"bids"`
	BidderMetadata []BidderMetadata `protobuf:"bytes,5,rep,name=bidder_metadata,json=bidderMetadata,proto3" json:"bidder_metadata"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1429e2d6f9d81cf, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "bitsong.marketplace.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("bitsong/marketplace/genesis.proto", fileDescriptor_e1429e2d6f9d81cf) }

var fileDescriptor_e1429e2d6f9d81cf = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbd, 0x4a, 0x03, 0x41,
	0x10, 0xc7, 0xef, 0x62, 0x0c, 0x72, 0x11, 0x85, 0xc3, 0xe2, 0x88, 0xb0, 0xc6, 0x8f, 0x22, 0x20,
	0xde, 0x92, 0x58, 0x69, 0x97, 0x14, 0x5a, 0x09, 0xa2, 0x85, 0x90, 0x46, 0x66, 0x3f, 0xb2, 0x2e,
	0xe6, 0x32, 0xc7, 0xed, 0x46, 0xf4, 0x2d, 0x7c, 0x15, 0xdf, 0x22, 0x65, 0x4a, 0x2b, 0xd1, 0xe4,
	0x45, 0xc4, 0xcd, 0x2a, 0x06, 0xe4, 0xec, 0x86, 0xe1, 0xff, 0xfb, 0xcd, 0xc0, 0x3f, 0xda, 0x65,
	0xda, 0x1a, 0x1c, 0x29, 0x9a, 0x41, 0x71, 0x2f, 0x6d, 0x3e, 0x04, 0x2e, 0xa9, 0x92, 0x23, 0x69,
	0xb4, 0x49, 0xf3, 0x02, 0x2d, 0xc6, 0xdb, 0x3e, 0x92, 0xfe, 0x8a, 0xa4, 0x0f, 0x6d, 0x26, 0x2d,
	0xb4, 0x1b, 0x84, 0xa3, 0xc9, 0xd0, 0x50, 0x06, 0x46, 0x52, 0xbf, 0xa4, 0x1c, 0xf5, 0x68, 0x01,
	0x37, 0xb6, 0x14, 0x2a, 0x74, 0x23, 0xfd, 0x9a, 0xfc, 0xf6, 0xcf, 0xab, 0x30, 0xe6, 0x56, 0xa3,
	0x07, 0xf7, 0x5e, 0x2a, 0xd1, 0xfa, 0xf9, 0xe2, 0x8f, 0x6b, 0x0b, 0x56, 0xc6, 0xdd, 0xa8, 0x96,
	0x43, 0x01, 0x99, 0x49, 0xc2, 0x66, 0xd8, 0xaa, 0x77, 0xf6, 0xd3, 0x92, 0xbf, 0xd2, 0x4b, 0x17,
	0xed, 0x55, 0x27, 0x6f, 0x3b, 0xc1, 0x95, 0x07, 0xe3, 0xb3, 0x68, 0xcd, 0x1f, 0x31, 0x49, 0xa5,
	0xb9, 0xd2, 0xaa, 0x77, 0x0e, 0x4a, 0x25, 0xdd, 0x45, 0xd8, 0x5b, 0x7e, 0xd8, 0xf8, 0x34, 0xaa,
	0x32, 0x2d, 0x4c, 0x52, 0x75, 0x8e, 0x66, 0xa9, 0xa3, 0xa7, 0x85, 0xe7, 0x1d, 0x13, 0xf7, 0xa3,
	0x4d, 0xa6, 0x85, 0x90, 0xc5, 0x6d, 0x26, 0x2d, 0x08, 0xb0, 0x90, 0xac, 0x3a, 0xcd, 0xe1, 0x7f,
	0x1a, 0x21, 0x8b, 0x0b, 0x8f, 0x78, 0xe3, 0x06, 0x5b, 0xde, 0xde, 0x4c, 0x3e, 0x48, 0x30, 0x99,
	0x91, 0x70, 0x3a, 0x23, 0xe1, 0xfb, 0x8c, 0x84, 0xcf, 0x73, 0x12, 0x4c, 0xe7, 0x24, 0x78, 0x9d,
	0x93, 0xa0, 0x7f, 0xa2, 0xb4, 0xbd, 0x1b, 0xb3, 0x94, 0x63, 0x46, 0xfd, 0x29, 0x1c, 0x0c, 0x34,
	0xd7, 0x30, 0xa4, 0x0a, 0x8f, 0xbe, 0x2b, 0x79, 0x5c, 0x2a, 0xc5, 0x3e, 0xe5, 0xd2, 0xb0, 0x9a,
	0xeb, 0xe4, 0xf8, 0x33, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x64, 0xcf, 0xf1, 0x2e, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BidderMetadata) > 0 {
		for iNdEx := len(m.BidderMetadata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BidderMetadata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Bids) > 0 {
		for iNdEx := len(m.Bids) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Bids[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Auctions) > 0 {
		for iNdEx := len(m.Auctions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Auctions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Auctions) > 0 {
		for _, e := range m.Auctions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Bids) > 0 {
		for _, e := range m.Bids {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BidderMetadata) > 0 {
		for _, e := range m.BidderMetadata {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Auctions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Auctions = append(m.Auctions, Auction{})
			if err := m.Auctions[len(m.Auctions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bids", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bids = append(m.Bids, Bid{})
			if err := m.Bids[len(m.Bids)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidderMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BidderMetadata = append(m.BidderMetadata, BidderMetadata{})
			if err := m.BidderMetadata[len(m.BidderMetadata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
