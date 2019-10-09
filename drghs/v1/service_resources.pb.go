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
// source: service_resources.proto

package drghs_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for [SampleService.ListRepositories][].
type ListRepositoriesRequest struct {
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
	//     field3 IN (value3, value4)
	//     field4 LIKE "%somestring%"
	//
	// Valid filter fields are: `name` and `owner`.
	//
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Specify how the results should be sorted. The fields supported
	// for sorting are `name` and `size`.
	// The default ordering is by `name`. Prefix with `-` to specify
	// descending order, e.g. `-name`.
	OrderBy              string   `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRepositoriesRequest) Reset()         { *m = ListRepositoriesRequest{} }
func (m *ListRepositoriesRequest) String() string { return proto.CompactTextString(m) }
func (*ListRepositoriesRequest) ProtoMessage()    {}
func (*ListRepositoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_761f26e4ca1188dd, []int{0}
}

func (m *ListRepositoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRepositoriesRequest.Unmarshal(m, b)
}
func (m *ListRepositoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRepositoriesRequest.Marshal(b, m, deterministic)
}
func (m *ListRepositoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRepositoriesRequest.Merge(m, src)
}
func (m *ListRepositoriesRequest) XXX_Size() int {
	return xxx_messageInfo_ListRepositoriesRequest.Size(m)
}
func (m *ListRepositoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRepositoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRepositoriesRequest proto.InternalMessageInfo

func (m *ListRepositoriesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListRepositoriesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListRepositoriesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListRepositoriesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListRepositoriesRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

//Response message for [SampleService.ListRepositories][].
type ListRepositoriesResponse struct {
	// The list of [Repositories][Repository].
	Repositories []*Repository `protobuf:"bytes,1,rep,name=repositories,proto3" json:"repositories,omitempty"`
	// A token to retrieve the next page of results, or empty if there are no
	// more results in the list. Pass this value in
	// [ListRepositoriesRequest.page_token][] to retrieve the next page of
	// results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The total number of repositories that matched the query.
	Total                int32    `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRepositoriesResponse) Reset()         { *m = ListRepositoriesResponse{} }
func (m *ListRepositoriesResponse) String() string { return proto.CompactTextString(m) }
func (*ListRepositoriesResponse) ProtoMessage()    {}
func (*ListRepositoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_761f26e4ca1188dd, []int{1}
}

func (m *ListRepositoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRepositoriesResponse.Unmarshal(m, b)
}
func (m *ListRepositoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRepositoriesResponse.Marshal(b, m, deterministic)
}
func (m *ListRepositoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRepositoriesResponse.Merge(m, src)
}
func (m *ListRepositoriesResponse) XXX_Size() int {
	return xxx_messageInfo_ListRepositoriesResponse.Size(m)
}
func (m *ListRepositoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRepositoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRepositoriesResponse proto.InternalMessageInfo

func (m *ListRepositoriesResponse) GetRepositories() []*Repository {
	if m != nil {
		return m.Repositories
	}
	return nil
}

func (m *ListRepositoriesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListRepositoriesResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*ListRepositoriesRequest)(nil), "drghs.v1.ListRepositoriesRequest")
	proto.RegisterType((*ListRepositoriesResponse)(nil), "drghs.v1.ListRepositoriesResponse")
}

func init() { proto.RegisterFile("service_resources.proto", fileDescriptor_761f26e4ca1188dd) }

var fileDescriptor_761f26e4ca1188dd = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4f, 0x02, 0x31,
	0x10, 0x85, 0x53, 0x70, 0x11, 0xaa, 0x86, 0xa4, 0x21, 0x52, 0x51, 0x93, 0x0d, 0x07, 0xb3, 0xa7,
	0x25, 0xea, 0xc5, 0xb3, 0x67, 0x0f, 0xa6, 0x7a, 0xdf, 0x14, 0x18, 0xd7, 0xc6, 0x4d, 0xa7, 0xb6,
	0x03, 0x71, 0xf9, 0x19, 0xfe, 0x02, 0x7f, 0xaa, 0xd9, 0x2e, 0x12, 0xd4, 0xe3, 0xfb, 0xe6, 0xbd,
	0xe9, 0xeb, 0xf0, 0x71, 0x00, 0xbf, 0x36, 0x0b, 0x28, 0x3c, 0x04, 0x5c, 0xf9, 0x05, 0x84, 0xdc,
	0x79, 0x24, 0x14, 0xfd, 0xa5, 0x2f, 0x5f, 0x43, 0xbe, 0xbe, 0x9e, 0x5c, 0x94, 0x88, 0x65, 0x05,
	0x33, 0xed, 0xcc, 0x4c, 0x5b, 0x8b, 0xa4, 0xc9, 0xa0, 0xdd, 0xfa, 0x26, 0xc3, 0x3f, 0xc1, 0xe9,
	0x17, 0xe3, 0xe3, 0x07, 0x13, 0x48, 0x81, 0xc3, 0x60, 0x08, 0xbd, 0x81, 0xa0, 0xe0, 0x7d, 0x05,
	0x81, 0xc4, 0x29, 0xef, 0x39, 0xed, 0xc1, 0x92, 0x64, 0x29, 0xcb, 0x06, 0x6a, 0xab, 0xc4, 0x39,
	0x1f, 0x38, 0x5d, 0x42, 0x11, 0xcc, 0x06, 0x64, 0x27, 0x65, 0x59, 0xa2, 0xfa, 0x0d, 0x78, 0x32,
	0x1b, 0x10, 0x97, 0x9c, 0xc7, 0x21, 0xe1, 0x1b, 0x58, 0xd9, 0x8d, 0xc1, 0x68, 0x7f, 0x6e, 0x40,
	0xb3, 0xf3, 0xc5, 0x54, 0x04, 0x5e, 0x1e, 0xb4, 0x3b, 0x5b, 0x25, 0xce, 0x78, 0x1f, 0xfd, 0x12,
	0x7c, 0x31, 0xaf, 0x65, 0x12, 0x27, 0x87, 0x51, 0xdf, 0xd7, 0xd3, 0x4f, 0xc6, 0xe5, 0xff, 0x8a,
	0xc1, 0xa1, 0x0d, 0x20, 0xee, 0xf8, 0xb1, 0xdf, 0xe3, 0x92, 0xa5, 0xdd, 0xec, 0xe8, 0x66, 0x94,
	0xff, 0xdc, 0x23, 0xdf, 0xa5, 0x6a, 0xf5, 0xcb, 0x29, 0xae, 0xf8, 0xd0, 0xc2, 0x07, 0x15, 0x7b,
	0x6d, 0x3b, 0xf1, 0xe1, 0x93, 0x06, 0x3f, 0xee, 0x1a, 0x8f, 0x78, 0x42, 0x48, 0xba, 0x8a, 0x7f,
	0x49, 0x54, 0x2b, 0xe6, 0xbd, 0x78, 0xbe, 0xdb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0xe4,
	0xc3, 0x3a, 0x92, 0x01, 0x00, 0x00,
}
