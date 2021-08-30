// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: interfaces/test_server/test.proto

package generated

import (
	context "context"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/maanasasubrahmanyam-sd/go-proto-validators"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//protoc   -I ./interfaces/  -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0  --go_out="plugins=grpc:./generated"    --validate_out="lang=go:./generated" ./interfaces/test_server/*.proto
//protoc  --proto_path=${GOPATH}/src --proto_path=${GOPATH}/pkg/mod/github.com/gogo/protobuf@v1.3.2 --proto_path=${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0 --proto_path=${GOPATH}/pkg/mod/github.com/maanasasubrahmanyam-sd --proto_path=${GOPATH}/src/github.com/maanasasubrahmanyam-sd/customeValTest/interfaces/test_server  --validate_out="lang=go:./generated" --govalidators_out=gogoimport=true:. C:/Users/Rohit/go/src/github.com/maanasasubrahmanyam-sd/customeValTest/interfaces/test_server/test.proto
//protoc   -I ./interfaces/  -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.1.0 -I ${GOPATH}/src -I $GOPATH/src/github.com/google/protobuf/src/ --govalidators_out=./generated --go_out="plugins=grpc:./generated" --validate_out="lang=go:./generated"  --proto_path=C:/Users/Rohit/go/src/github.com/maanasasubrahmanyam-sd/customeValTest/interfaces/test_server/test.proto ./interfaces/test_server/*.proto
type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query   string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	EmailId string `protobuf:"bytes,2,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	//string id = 4 [(validator.field) = {alpha: true}];
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_test_server_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_test_server_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_interfaces_test_server_test_proto_rawDescGZIP(), []int{0}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchRequest) GetEmailId() string {
	if x != nil {
		return x.EmailId
	}
	return ""
}

func (x *SearchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

//* SearchResponse represents the response of the Search
type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchResponse string `protobuf:"bytes,1,opt,name=search_response,json=searchResponse,proto3" json:"search_response,omitempty"` // the search response
	PageNumber     int32  `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`            // the page number
	TotalPages     int32  `protobuf:"varint,3,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`            //total number of pages
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_test_server_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_test_server_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_interfaces_test_server_test_proto_rawDescGZIP(), []int{1}
}

func (x *SearchResponse) GetSearchResponse() string {
	if x != nil {
		return x.SearchResponse
	}
	return ""
}

func (x *SearchResponse) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *SearchResponse) GetTotalPages() int32 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

var File_interfaces_test_server_test_proto protoreflect.FileDescriptor

var file_interfaces_test_server_test_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x01,
	0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3a, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24,
	0xba, 0xe9, 0xc0, 0x03, 0x1f, 0x72, 0x1d, 0x28, 0x32, 0x32, 0x19, 0x28, 0x5b, 0x41, 0x2d, 0x5a,
	0x61, 0x2d, 0x7a, 0x5d, 0x2b, 0x29, 0x20, 0x28, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x5d,
	0x2b, 0x29, 0x2a, 0x24, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x08, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba,
	0xe9, 0xc0, 0x03, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x07, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xba, 0xe9, 0xc0, 0x03, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0xe2, 0xdf, 0x1f, 0x05, 0x58, 0x01, 0x90, 0x01, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x7b, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65,
	0x73, 0x32, 0x3c, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x0e, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interfaces_test_server_test_proto_rawDescOnce sync.Once
	file_interfaces_test_server_test_proto_rawDescData = file_interfaces_test_server_test_proto_rawDesc
)

func file_interfaces_test_server_test_proto_rawDescGZIP() []byte {
	file_interfaces_test_server_test_proto_rawDescOnce.Do(func() {
		file_interfaces_test_server_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_interfaces_test_server_test_proto_rawDescData)
	})
	return file_interfaces_test_server_test_proto_rawDescData
}

var file_interfaces_test_server_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_interfaces_test_server_test_proto_goTypes = []interface{}{
	(*SearchRequest)(nil),  // 0: SearchRequest
	(*SearchResponse)(nil), // 1: SearchResponse
}
var file_interfaces_test_server_test_proto_depIdxs = []int32{
	0, // 0: SearchService.Search:input_type -> SearchRequest
	1, // 1: SearchService.Search:output_type -> SearchResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_interfaces_test_server_test_proto_init() }
func file_interfaces_test_server_test_proto_init() {
	if File_interfaces_test_server_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interfaces_test_server_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_interfaces_test_server_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interfaces_test_server_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interfaces_test_server_test_proto_goTypes,
		DependencyIndexes: file_interfaces_test_server_test_proto_depIdxs,
		MessageInfos:      file_interfaces_test_server_test_proto_msgTypes,
	}.Build()
	File_interfaces_test_server_test_proto = out.File
	file_interfaces_test_server_test_proto_rawDesc = nil
	file_interfaces_test_server_test_proto_goTypes = nil
	file_interfaces_test_server_test_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchServiceClient interface {
	// Method to Search a test service
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
type SearchServiceServer interface {
	// Method to Search a test service
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

// UnimplementedSearchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (*UnimplementedSearchServiceServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interfaces/test_server/test.proto",
}
