// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: auction/auction.proto

package auction

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Auction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AuctionId     string                 `protobuf:"bytes,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	NftId         string                 `protobuf:"bytes,2,opt,name=nft_id,json=nftId,proto3" json:"nft_id,omitempty"`
	Collection    string                 `protobuf:"bytes,3,opt,name=collection,proto3" json:"collection,omitempty"`
	CurrentBid    string                 `protobuf:"bytes,4,opt,name=current_bid,json=currentBid,proto3" json:"current_bid,omitempty"`
	EndTime       string                 `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Auction) Reset() {
	*x = Auction{}
	mi := &file_auction_auction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auction) ProtoMessage() {}

func (x *Auction) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auction.ProtoReflect.Descriptor instead.
func (*Auction) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{0}
}

func (x *Auction) GetAuctionId() string {
	if x != nil {
		return x.AuctionId
	}
	return ""
}

func (x *Auction) GetNftId() string {
	if x != nil {
		return x.NftId
	}
	return ""
}

func (x *Auction) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *Auction) GetCurrentBid() string {
	if x != nil {
		return x.CurrentBid
	}
	return ""
}

func (x *Auction) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type AuctionsAvailableRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuctionsAvailableRequest) Reset() {
	*x = AuctionsAvailableRequest{}
	mi := &file_auction_auction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuctionsAvailableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionsAvailableRequest) ProtoMessage() {}

func (x *AuctionsAvailableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionsAvailableRequest.ProtoReflect.Descriptor instead.
func (*AuctionsAvailableRequest) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{1}
}

type AuctionsAvailableResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []*Auction             `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuctionsAvailableResponse) Reset() {
	*x = AuctionsAvailableResponse{}
	mi := &file_auction_auction_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuctionsAvailableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuctionsAvailableResponse) ProtoMessage() {}

func (x *AuctionsAvailableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auction_auction_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuctionsAvailableResponse.ProtoReflect.Descriptor instead.
func (*AuctionsAvailableResponse) Descriptor() ([]byte, []int) {
	return file_auction_auction_proto_rawDescGZIP(), []int{2}
}

func (x *AuctionsAvailableResponse) GetData() []*Auction {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_auction_auction_proto protoreflect.FileDescriptor

var file_auction_auction_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x9b, 0x01, 0x0a, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6e,
	0x66, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x66, 0x74,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x42, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x1a,
	0x0a, 0x18, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x19, 0x41, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x72, 0x0a,
	0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x41, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x19, 0x70, 0x68,
	0x6f, 0x65, 0x6e, 0x69, 0x78, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x07,
	0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0xca, 0x02, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0xe2, 0x02, 0x13, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_auction_auction_proto_rawDescOnce sync.Once
	file_auction_auction_proto_rawDescData []byte
)

func file_auction_auction_proto_rawDescGZIP() []byte {
	file_auction_auction_proto_rawDescOnce.Do(func() {
		file_auction_auction_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_auction_auction_proto_rawDesc), len(file_auction_auction_proto_rawDesc)))
	})
	return file_auction_auction_proto_rawDescData
}

var file_auction_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_auction_auction_proto_goTypes = []any{
	(*Auction)(nil),                   // 0: auction.Auction
	(*AuctionsAvailableRequest)(nil),  // 1: auction.AuctionsAvailableRequest
	(*AuctionsAvailableResponse)(nil), // 2: auction.AuctionsAvailableResponse
}
var file_auction_auction_proto_depIdxs = []int32{
	0, // 0: auction.AuctionsAvailableResponse.data:type_name -> auction.Auction
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auction_auction_proto_init() }
func file_auction_auction_proto_init() {
	if File_auction_auction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_auction_auction_proto_rawDesc), len(file_auction_auction_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auction_auction_proto_goTypes,
		DependencyIndexes: file_auction_auction_proto_depIdxs,
		MessageInfos:      file_auction_auction_proto_msgTypes,
	}.Build()
	File_auction_auction_proto = out.File
	file_auction_auction_proto_goTypes = nil
	file_auction_auction_proto_depIdxs = nil
}
