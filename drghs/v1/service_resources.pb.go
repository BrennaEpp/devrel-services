// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_resources.proto

package drghs_v1

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Request message for [SampleService.ListRepositories][].
type ListRepositoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the owner associated with the
	// [Repositories][Repository], in the format `owners/*`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. Limit the number of [Repositories][Repository] to include in the
	// response. Fewer repositories than requested might be returned.
	//
	// The maximum page size is `100`. If unspecified, the page size will be the
	// maximum. Further [Repositories][Repository] can subsequently be obtained
	// by including the [ListRepositoriesResponse.next_page_token][] in a
	// subsequent request.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. To request the first page of results, `page_token` must be empty.
	// To request the next page of results, page_token must be the value of
	// [ListRepositoriesResponse.next_page_token][] returned by a previous call to
	// [Repositorieservice.ListRepositories][].
	//
	// The page token is valid for only 2 hours.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. Filter expression used to include in the response only those
	// resources that match the filter. Filter must be in following the format:
	//
	//     field1=123
	//     field2="Foo bar"
	//     field3 IN ["value3", "value4"]
	//
	// Valid filter fields are: `repo` and `owner`.
	//
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Specify how the results should be sorted. The fields supported
	// for sorting are `name` and `size`.
	// The default ordering is by `name`. Prefix with `-` to specify
	// descending order, e.g. `-name`.
	OrderBy string `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
}

func (x *ListRepositoriesRequest) Reset() {
	*x = ListRepositoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepositoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepositoriesRequest) ProtoMessage() {}

func (x *ListRepositoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepositoriesRequest.ProtoReflect.Descriptor instead.
func (*ListRepositoriesRequest) Descriptor() ([]byte, []int) {
	return file_service_resources_proto_rawDescGZIP(), []int{0}
}

func (x *ListRepositoriesRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListRepositoriesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRepositoriesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListRepositoriesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *ListRepositoriesRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

// Response message for [SampleService.ListRepositories][].
type ListRepositoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of [Repositories][Repository].
	Repositories []*Repository `protobuf:"bytes,1,rep,name=repositories,proto3" json:"repositories,omitempty"`
	// A token to retrieve the next page of results, or empty if there are no
	// more results in the list. Pass this value in
	// [ListRepositoriesRequest.page_token][] to retrieve the next page of
	// results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The total number of repositories that matched the query.
	Total int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListRepositoriesResponse) Reset() {
	*x = ListRepositoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepositoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepositoriesResponse) ProtoMessage() {}

func (x *ListRepositoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepositoriesResponse.ProtoReflect.Descriptor instead.
func (*ListRepositoriesResponse) Descriptor() ([]byte, []int) {
	return file_service_resources_proto_rawDescGZIP(), []int{1}
}

func (x *ListRepositoriesResponse) GetRepositories() []*Repository {
	if x != nil {
		return x.Repositories
	}
	return nil
}

func (x *ListRepositoriesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListRepositoriesResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_service_resources_proto protoreflect.FileDescriptor

var file_service_resources_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x72, 0x67, 0x68, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x22, 0x92, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x72, 0x67,
	0x68, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_resources_proto_rawDescOnce sync.Once
	file_service_resources_proto_rawDescData = file_service_resources_proto_rawDesc
)

func file_service_resources_proto_rawDescGZIP() []byte {
	file_service_resources_proto_rawDescOnce.Do(func() {
		file_service_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_resources_proto_rawDescData)
	})
	return file_service_resources_proto_rawDescData
}

var file_service_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_resources_proto_goTypes = []interface{}{
	(*ListRepositoriesRequest)(nil),  // 0: drghs.v1.ListRepositoriesRequest
	(*ListRepositoriesResponse)(nil), // 1: drghs.v1.ListRepositoriesResponse
	(*Repository)(nil),               // 2: drghs.v1.Repository
}
var file_service_resources_proto_depIdxs = []int32{
	2, // 0: drghs.v1.ListRepositoriesResponse.repositories:type_name -> drghs.v1.Repository
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_resources_proto_init() }
func file_service_resources_proto_init() {
	if File_service_resources_proto != nil {
		return
	}
	file_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRepositoriesRequest); i {
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
		file_service_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRepositoriesResponse); i {
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
			RawDescriptor: file_service_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_resources_proto_goTypes,
		DependencyIndexes: file_service_resources_proto_depIdxs,
		MessageInfos:      file_service_resources_proto_msgTypes,
	}.Build()
	File_service_resources_proto = out.File
	file_service_resources_proto_rawDesc = nil
	file_service_resources_proto_goTypes = nil
	file_service_resources_proto_depIdxs = nil
}
