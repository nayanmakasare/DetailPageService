// Code generated by protoc-gen-go. DO NOT EDIT.
// source: DetailPageService.proto

package DetailPageService

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Row Request
type TileType int32

const (
	TileType_ImageTile     TileType = 0
	TileType_VideoTile     TileType = 1
	TileType_FeatureTile   TileType = 2
	TileType_AdvertiseTile TileType = 3
	TileType_CarouselTile  TileType = 4
)

var TileType_name = map[int32]string{
	0: "ImageTile",
	1: "VideoTile",
	2: "FeatureTile",
	3: "AdvertiseTile",
	4: "CarouselTile",
}

var TileType_value = map[string]int32{
	"ImageTile":     0,
	"VideoTile":     1,
	"FeatureTile":   2,
	"AdvertiseTile": 3,
	"CarouselTile":  4,
}

func (x TileType) String() string {
	return proto.EnumName(TileType_name, int32(x))
}

func (TileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2aea33dfb16e7225, []int{0}
}

type TileInfoRequest struct {
	TileId               string   `protobuf:"bytes,1,opt,name=tileId,proto3" json:"tileId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *TileInfoRequest) Reset()         { *m = TileInfoRequest{} }
func (m *TileInfoRequest) String() string { return proto.CompactTextString(m) }
func (*TileInfoRequest) ProtoMessage()    {}
func (*TileInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aea33dfb16e7225, []int{0}
}

func (m *TileInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TileInfoRequest.Unmarshal(m, b)
}
func (m *TileInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TileInfoRequest.Marshal(b, m, deterministic)
}
func (m *TileInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TileInfoRequest.Merge(m, src)
}
func (m *TileInfoRequest) XXX_Size() int {
	return xxx_messageInfo_TileInfoRequest.Size(m)
}
func (m *TileInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TileInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TileInfoRequest proto.InternalMessageInfo

func (m *TileInfoRequest) GetTileId() string {
	if m != nil {
		return m.TileId
	}
	return ""
}

type DetailTileInfo struct {
	Title                string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Synopsis             string            `protobuf:"bytes,2,opt,name=synopsis,proto3" json:"synopsis,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Button               []*Button         `protobuf:"bytes,4,rep,name=button,proto3" json:"button,omitempty"`
	ContentTile          []*ContentTile    `protobuf:"bytes,5,rep,name=contentTile,proto3" json:"contentTile,omitempty"`
	BackDrop             string            `protobuf:"bytes,6,opt,name=backDrop,proto3" json:"backDrop,omitempty"`
	Poster               string            `protobuf:"bytes,7,opt,name=poster,proto3" json:"poster,omitempty"`
	Portrait             string            `protobuf:"bytes,8,opt,name=portrait,proto3" json:"portrait,omitempty"`
	Target               []string          `protobuf:"bytes,9,rep,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" bson:"-"`
	XXX_unrecognized     []byte            `json:"-" bson:"-"`
	XXX_sizecache        int32             `json:"-" bson:"-"`
}

func (m *DetailTileInfo) Reset()         { *m = DetailTileInfo{} }
func (m *DetailTileInfo) String() string { return proto.CompactTextString(m) }
func (*DetailTileInfo) ProtoMessage()    {}
func (*DetailTileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aea33dfb16e7225, []int{1}
}

func (m *DetailTileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailTileInfo.Unmarshal(m, b)
}
func (m *DetailTileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailTileInfo.Marshal(b, m, deterministic)
}
func (m *DetailTileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailTileInfo.Merge(m, src)
}
func (m *DetailTileInfo) XXX_Size() int {
	return xxx_messageInfo_DetailTileInfo.Size(m)
}
func (m *DetailTileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailTileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DetailTileInfo proto.InternalMessageInfo

func (m *DetailTileInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DetailTileInfo) GetSynopsis() string {
	if m != nil {
		return m.Synopsis
	}
	return ""
}

func (m *DetailTileInfo) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *DetailTileInfo) GetButton() []*Button {
	if m != nil {
		return m.Button
	}
	return nil
}

func (m *DetailTileInfo) GetContentTile() []*ContentTile {
	if m != nil {
		return m.ContentTile
	}
	return nil
}

func (m *DetailTileInfo) GetBackDrop() string {
	if m != nil {
		return m.BackDrop
	}
	return ""
}

func (m *DetailTileInfo) GetPoster() string {
	if m != nil {
		return m.Poster
	}
	return ""
}

func (m *DetailTileInfo) GetPortrait() string {
	if m != nil {
		return m.Portrait
	}
	return ""
}

func (m *DetailTileInfo) GetTarget() []string {
	if m != nil {
		return m.Target
	}
	return nil
}

type Button struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Icon                 string   `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty"`
	PackageName          string   `protobuf:"bytes,3,opt,name=packageName,proto3" json:"packageName,omitempty"`
	Action               string   `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Button) Reset()         { *m = Button{} }
func (m *Button) String() string { return proto.CompactTextString(m) }
func (*Button) ProtoMessage()    {}
func (*Button) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aea33dfb16e7225, []int{2}
}

func (m *Button) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Button.Unmarshal(m, b)
}
func (m *Button) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Button.Marshal(b, m, deterministic)
}
func (m *Button) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Button.Merge(m, src)
}
func (m *Button) XXX_Size() int {
	return xxx_messageInfo_Button.Size(m)
}
func (m *Button) XXX_DiscardUnknown() {
	xxx_messageInfo_Button.DiscardUnknown(m)
}

var xxx_messageInfo_Button proto.InternalMessageInfo

func (m *Button) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Button) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Button) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *Button) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type ContentTile struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	MediaUrl             string   `protobuf:"bytes,2,opt,name=mediaUrl,proto3" json:"mediaUrl,omitempty"`
	TileType             TileType `protobuf:"varint,3,opt,name=tileType,proto3,enum=DetailPageService.TileType" json:"tileType,omitempty"`
	Poster               string   `protobuf:"bytes,4,opt,name=poster,proto3" json:"poster,omitempty"`
	Portrait             string   `protobuf:"bytes,5,opt,name=portrait,proto3" json:"portrait,omitempty"`
	IsDetailPage         bool     `protobuf:"varint,6,opt,name=isDetailPage,proto3" json:"isDetailPage,omitempty"`
	PackageName          string   `protobuf:"bytes,7,opt,name=packageName,proto3" json:"packageName,omitempty"`
	ContentId            string   `protobuf:"bytes,8,opt,name=contentId,proto3" json:"contentId,omitempty"`
	Target               []string `protobuf:"bytes,9,rep,name=target,proto3" json:"target,omitempty"`
	RealeaseDate         string   `protobuf:"bytes,10,opt,name=realeaseDate,proto3" json:"realeaseDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *ContentTile) Reset()         { *m = ContentTile{} }
func (m *ContentTile) String() string { return proto.CompactTextString(m) }
func (*ContentTile) ProtoMessage()    {}
func (*ContentTile) Descriptor() ([]byte, []int) {
	return fileDescriptor_2aea33dfb16e7225, []int{3}
}

func (m *ContentTile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentTile.Unmarshal(m, b)
}
func (m *ContentTile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentTile.Marshal(b, m, deterministic)
}
func (m *ContentTile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentTile.Merge(m, src)
}
func (m *ContentTile) XXX_Size() int {
	return xxx_messageInfo_ContentTile.Size(m)
}
func (m *ContentTile) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentTile.DiscardUnknown(m)
}

var xxx_messageInfo_ContentTile proto.InternalMessageInfo

func (m *ContentTile) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ContentTile) GetMediaUrl() string {
	if m != nil {
		return m.MediaUrl
	}
	return ""
}

func (m *ContentTile) GetTileType() TileType {
	if m != nil {
		return m.TileType
	}
	return TileType_ImageTile
}

func (m *ContentTile) GetPoster() string {
	if m != nil {
		return m.Poster
	}
	return ""
}

func (m *ContentTile) GetPortrait() string {
	if m != nil {
		return m.Portrait
	}
	return ""
}

func (m *ContentTile) GetIsDetailPage() bool {
	if m != nil {
		return m.IsDetailPage
	}
	return false
}

func (m *ContentTile) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *ContentTile) GetContentId() string {
	if m != nil {
		return m.ContentId
	}
	return ""
}

func (m *ContentTile) GetTarget() []string {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *ContentTile) GetRealeaseDate() string {
	if m != nil {
		return m.RealeaseDate
	}
	return ""
}

func init() {
	proto.RegisterEnum("DetailPageService.TileType", TileType_name, TileType_value)
	proto.RegisterType((*TileInfoRequest)(nil), "DetailPageService.TileInfoRequest")
	proto.RegisterType((*DetailTileInfo)(nil), "DetailPageService.DetailTileInfo")
	proto.RegisterMapType((map[string]string)(nil), "DetailPageService.DetailTileInfo.MetadataEntry")
	proto.RegisterType((*Button)(nil), "DetailPageService.Button")
	proto.RegisterType((*ContentTile)(nil), "DetailPageService.ContentTile")
}

func init() { proto.RegisterFile("DetailPageService.proto", fileDescriptor_2aea33dfb16e7225) }

var fileDescriptor_2aea33dfb16e7225 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xfe, 0x1d, 0xa7, 0xa9, 0x33, 0x69, 0xda, 0x74, 0xf5, 0x03, 0x4b, 0xa8, 0x50, 0xf0, 0xa9,
	0xf4, 0xd0, 0x88, 0x72, 0x00, 0x95, 0x0b, 0xd0, 0x00, 0x8a, 0x10, 0x08, 0x99, 0xc0, 0x11, 0x69,
	0x63, 0x0f, 0xd6, 0x2a, 0x8e, 0xd7, 0xac, 0x37, 0x11, 0xb9, 0x72, 0xe7, 0xc4, 0x53, 0xf0, 0x2e,
	0xdc, 0x78, 0x05, 0x1e, 0x04, 0xed, 0xae, 0x9d, 0x38, 0x4d, 0x22, 0x6e, 0xfe, 0xbe, 0x99, 0x6f,
	0xd6, 0xf3, 0xcd, 0xec, 0xc2, 0xad, 0x01, 0x2a, 0xc6, 0x93, 0x77, 0x2c, 0xc6, 0xf7, 0x28, 0xe7,
	0x3c, 0xc4, 0xf3, 0x4c, 0x0a, 0x25, 0xc8, 0xf1, 0x46, 0xa0, 0x7b, 0x12, 0x0b, 0x11, 0x27, 0xd8,
	0x67, 0x19, 0xef, 0xb3, 0x34, 0x15, 0x8a, 0x29, 0x2e, 0xd2, 0xdc, 0x0a, 0xfc, 0xfb, 0x70, 0x34,
	0xe2, 0x09, 0x0e, 0xd3, 0xcf, 0x22, 0xc0, 0x2f, 0x33, 0xcc, 0x15, 0xb9, 0x09, 0x0d, 0xa5, 0xa9,
	0x88, 0x3a, 0x3d, 0xe7, 0xb4, 0x19, 0x14, 0xc8, 0xff, 0xe9, 0xc2, 0xa1, 0x2d, 0x5f, 0x2a, 0xc8,
	0xff, 0xb0, 0xa7, 0xb8, 0x4a, 0xb0, 0xc8, 0xb4, 0x80, 0x74, 0xc1, 0xcb, 0x17, 0xa9, 0xc8, 0x72,
	0x9e, 0xd3, 0x9a, 0x09, 0x2c, 0x31, 0x79, 0x0d, 0xde, 0x14, 0x15, 0x8b, 0x98, 0x62, 0xd4, 0xed,
	0xb9, 0xa7, 0xad, 0x8b, 0xfe, 0xf9, 0x66, 0x33, 0xeb, 0xc7, 0x9c, 0xbf, 0x29, 0x14, 0x2f, 0x52,
	0x25, 0x17, 0xc1, 0xb2, 0x00, 0x79, 0x00, 0x8d, 0xf1, 0x4c, 0x29, 0x91, 0xd2, 0xba, 0x29, 0x75,
	0x7b, 0x4b, 0xa9, 0xe7, 0x26, 0x21, 0x28, 0x12, 0xc9, 0x53, 0x68, 0x85, 0x22, 0x55, 0x98, 0x2a,
	0x5d, 0x9d, 0xee, 0x19, 0xdd, 0xdd, 0x2d, 0xba, 0xab, 0x55, 0x56, 0x50, 0x95, 0xe8, 0xee, 0xc6,
	0x2c, 0x9c, 0x0c, 0xa4, 0xc8, 0x68, 0xc3, 0x76, 0x57, 0x62, 0x6d, 0x5d, 0x26, 0x72, 0x85, 0x92,
	0xee, 0x5b, 0xeb, 0x2c, 0xd2, 0x9a, 0x4c, 0x48, 0x25, 0x19, 0x57, 0xd4, 0xb3, 0x9a, 0x12, 0x1b,
	0xbb, 0x99, 0x8c, 0x51, 0xd1, 0x66, 0xcf, 0x35, 0x76, 0x1b, 0xd4, 0x7d, 0x02, 0xed, 0xb5, 0xbe,
	0x49, 0x07, 0xdc, 0x09, 0x2e, 0x0a, 0xab, 0xf5, 0xa7, 0xb6, 0x7f, 0xce, 0x92, 0x19, 0x16, 0x2e,
	0x5b, 0x70, 0x59, 0x7b, 0xec, 0xf8, 0x09, 0x34, 0x6c, 0xe3, 0x3b, 0x46, 0x44, 0xa0, 0xce, 0x43,
	0x91, 0x16, 0x42, 0xf3, 0x4d, 0x7a, 0xd0, 0xca, 0x58, 0x38, 0x61, 0x31, 0xbe, 0x65, 0x53, 0xa4,
	0xae, 0x09, 0x55, 0x29, 0xfd, 0xab, 0x2c, 0xd4, 0xdb, 0x43, 0xeb, 0xb6, 0x3d, 0x8b, 0xfc, 0x5f,
	0x35, 0x68, 0x55, 0xfc, 0xda, 0xbd, 0x16, 0x53, 0x8c, 0x38, 0xfb, 0x20, 0x93, 0x72, 0x2d, 0x4a,
	0x4c, 0x1e, 0x81, 0xa7, 0xb7, 0x6c, 0xb4, 0xc8, 0xec, 0xc1, 0x87, 0x17, 0x77, 0xb6, 0xcc, 0x64,
	0x54, 0xa4, 0x04, 0xcb, 0xe4, 0x8a, 0xe3, 0xf5, 0x9d, 0x8e, 0xef, 0x5d, 0x73, 0xdc, 0x87, 0x03,
	0x9e, 0xaf, 0xaa, 0x9b, 0x29, 0x7a, 0xc1, 0x1a, 0x77, 0xdd, 0x8c, 0xfd, 0x4d, 0x33, 0x4e, 0xa0,
	0x59, 0xac, 0xc5, 0x30, 0x2a, 0x86, 0xba, 0x22, 0x76, 0x4d, 0x55, 0x9f, 0x2d, 0x91, 0x25, 0xc8,
	0x72, 0x1c, 0x30, 0x85, 0x14, 0x8c, 0x70, 0x8d, 0x3b, 0xfb, 0x04, 0x5e, 0xd9, 0x29, 0x69, 0x43,
	0x73, 0x38, 0x65, 0x31, 0x6a, 0xa2, 0xf3, 0x9f, 0x86, 0x1f, 0x79, 0x84, 0xc2, 0x40, 0x87, 0x1c,
	0x41, 0xeb, 0x25, 0x32, 0x35, 0x93, 0x36, 0x5e, 0x23, 0xc7, 0xd0, 0x7e, 0x16, 0xcd, 0x51, 0x2a,
	0x9e, 0x5b, 0xca, 0x25, 0x1d, 0x38, 0xb8, 0x62, 0x52, 0xcc, 0x72, 0x34, 0x17, 0xaa, 0x53, 0xbf,
	0xf8, 0xee, 0xc0, 0xe6, 0x3b, 0x41, 0xbe, 0x42, 0xfb, 0x15, 0x2a, 0xcb, 0x9b, 0xcb, 0xed, 0xef,
	0x98, 0x40, 0xe5, 0xad, 0xe8, 0xde, 0xfb, 0xe7, 0xe5, 0xf5, 0x7b, 0xdf, 0x7e, 0xff, 0xf9, 0x51,
	0xeb, 0xfa, 0x37, 0xfa, 0x91, 0x09, 0xf4, 0xe3, 0xea, 0x29, 0x97, 0xce, 0xd9, 0xb8, 0x61, 0x9e,
	0xa2, 0x87, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x60, 0xfb, 0xd4, 0x8e, 0xd6, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DetailPageServiceClient is the client API for DetailPageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DetailPageServiceClient interface {
	GetDetailInfo(ctx context.Context, in *TileInfoRequest, opts ...grpc.CallOption) (*DetailTileInfo, error)
}

type detailPageServiceClient struct {
	cc *grpc.ClientConn
}

func NewDetailPageServiceClient(cc *grpc.ClientConn) DetailPageServiceClient {
	return &detailPageServiceClient{cc}
}

func (c *detailPageServiceClient) GetDetailInfo(ctx context.Context, in *TileInfoRequest, opts ...grpc.CallOption) (*DetailTileInfo, error) {
	out := new(DetailTileInfo)
	err := c.cc.Invoke(ctx, "/DetailPageService.DetailPageService/GetDetailInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetailPageServiceServer is the server API for DetailPageService service.
type DetailPageServiceServer interface {
	GetDetailInfo(context.Context, *TileInfoRequest) (*DetailTileInfo, error)
}

// UnimplementedDetailPageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDetailPageServiceServer struct {
}

func (*UnimplementedDetailPageServiceServer) GetDetailInfo(ctx context.Context, req *TileInfoRequest) (*DetailTileInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailInfo not implemented")
}

func RegisterDetailPageServiceServer(s *grpc.Server, srv DetailPageServiceServer) {
	s.RegisterService(&_DetailPageService_serviceDesc, srv)
}

func _DetailPageService_GetDetailInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TileInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetailPageServiceServer).GetDetailInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DetailPageService.DetailPageService/GetDetailInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetailPageServiceServer).GetDetailInfo(ctx, req.(*TileInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DetailPageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DetailPageService.DetailPageService",
	HandlerType: (*DetailPageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDetailInfo",
			Handler:    _DetailPageService_GetDetailInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "DetailPageService.proto",
}
