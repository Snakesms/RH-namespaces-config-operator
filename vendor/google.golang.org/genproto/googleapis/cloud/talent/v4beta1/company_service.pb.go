// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/company_service.proto

package talent // import "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The Request of the CreateCompany method.
type CreateCompanyRequest struct {
	// Required.
	//
	// Resource name of the tenant under which the company is created.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/api-test-project/tenant/foo".
	//
	// Tenant id is optional and a default tenant is created if unspecified, for
	// example, "projects/api-test-project".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required.
	//
	// The company to be created.
	Company              *Company `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCompanyRequest) Reset()         { *m = CreateCompanyRequest{} }
func (m *CreateCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCompanyRequest) ProtoMessage()    {}
func (*CreateCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_company_service_1d879da651d4da28, []int{0}
}
func (m *CreateCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCompanyRequest.Unmarshal(m, b)
}
func (m *CreateCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCompanyRequest.Marshal(b, m, deterministic)
}
func (dst *CreateCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCompanyRequest.Merge(dst, src)
}
func (m *CreateCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCompanyRequest.Size(m)
}
func (m *CreateCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCompanyRequest proto.InternalMessageInfo

func (m *CreateCompanyRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateCompanyRequest) GetCompany() *Company {
	if m != nil {
		return m.Company
	}
	return nil
}

// Request for getting a company by name.
type GetCompanyRequest struct {
	// Required.
	//
	// The resource name of the company to be retrieved.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for
	// example, "projects/api-test-project/tenants/foo/companies/bar".
	//
	// Tenant id is optional and the default tenant is used if unspecified, for
	// example, "projects/api-test-project/companies/bar".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCompanyRequest) Reset()         { *m = GetCompanyRequest{} }
func (m *GetCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*GetCompanyRequest) ProtoMessage()    {}
func (*GetCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_company_service_1d879da651d4da28, []int{1}
}
func (m *GetCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCompanyRequest.Unmarshal(m, b)
}
func (m *GetCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCompanyRequest.Marshal(b, m, deterministic)
}
func (dst *GetCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCompanyRequest.Merge(dst, src)
}
func (m *GetCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_GetCompanyRequest.Size(m)
}
func (m *GetCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCompanyRequest proto.InternalMessageInfo

func (m *GetCompanyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request for updating a specified company.
type UpdateCompanyRequest struct {
	// Required.
	//
	// The company resource to replace the current resource in the system.
	Company *Company `protobuf:"bytes,1,opt,name=company,proto3" json:"company,omitempty"`
	// Optional but strongly recommended for the best service
	// experience.
	//
	// If [update_mask][google.cloud.talent.v4beta1.UpdateCompanyRequest.update_mask] is provided, only the specified fields in
	// [company][google.cloud.talent.v4beta1.UpdateCompanyRequest.company] are updated. Otherwise all the fields are updated.
	//
	// A field mask to specify the company fields to be updated. Only
	// top level fields of [Company][google.cloud.talent.v4beta1.Company] are supported.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateCompanyRequest) Reset()         { *m = UpdateCompanyRequest{} }
func (m *UpdateCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCompanyRequest) ProtoMessage()    {}
func (*UpdateCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_company_service_1d879da651d4da28, []int{2}
}
func (m *UpdateCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCompanyRequest.Unmarshal(m, b)
}
func (m *UpdateCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCompanyRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCompanyRequest.Merge(dst, src)
}
func (m *UpdateCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCompanyRequest.Size(m)
}
func (m *UpdateCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCompanyRequest proto.InternalMessageInfo

func (m *UpdateCompanyRequest) GetCompany() *Company {
	if m != nil {
		return m.Company
	}
	return nil
}

func (m *UpdateCompanyRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Request to delete a company.
type DeleteCompanyRequest struct {
	// Required.
	//
	// The resource name of the company to be deleted.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for
	// example, "projects/api-test-project/tenants/foo/companies/bar".
	//
	// Tenant id is optional and the default tenant is used if unspecified, for
	// example, "projects/api-test-project/companies/bar".
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCompanyRequest) Reset()         { *m = DeleteCompanyRequest{} }
func (m *DeleteCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCompanyRequest) ProtoMessage()    {}
func (*DeleteCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_company_service_1d879da651d4da28, []int{3}
}
func (m *DeleteCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCompanyRequest.Unmarshal(m, b)
}
func (m *DeleteCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCompanyRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCompanyRequest.Merge(dst, src)
}
func (m *DeleteCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCompanyRequest.Size(m)
}
func (m *DeleteCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCompanyRequest proto.InternalMessageInfo

func (m *DeleteCompanyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// List companies for which the client has ACL visibility.
type ListCompaniesRequest struct {
	// Required.
	//
	// Resource name of the tenant under which the company is created.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/api-test-project/tenant/foo".
	//
	// Tenant id is optional and the default tenant is used if unspecified, for
	// example, "projects/api-test-project".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional.
	//
	// The starting indicator from which to return results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional.
	//
	// The maximum number of companies to be returned, at most 100.
	// Default is 100 if a non-positive number is provided.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional.
	//
	// Set to true if the companies requested must have open jobs.
	//
	// Defaults to false.
	//
	// If true, at most [page_size][google.cloud.talent.v4beta1.ListCompaniesRequest.page_size] of companies are fetched, among which
	// only those with open jobs are returned.
	RequireOpenJobs      bool     `protobuf:"varint,4,opt,name=require_open_jobs,json=requireOpenJobs,proto3" json:"require_open_jobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCompaniesRequest) Reset()         { *m = ListCompaniesRequest{} }
func (m *ListCompaniesRequest) String() string { return proto.CompactTextString(m) }
func (*ListCompaniesRequest) ProtoMessage()    {}
func (*ListCompaniesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_company_service_1d879da651d4da28, []int{4}
}
func (m *ListCompaniesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCompaniesRequest.Unmarshal(m, b)
}
func (m *ListCompaniesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCompaniesRequest.Marshal(b, m, deterministic)
}
func (dst *ListCompaniesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCompaniesRequest.Merge(dst, src)
}
func (m *ListCompaniesRequest) XXX_Size() int {
	return xxx_messageInfo_ListCompaniesRequest.Size(m)
}
func (m *ListCompaniesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCompaniesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCompaniesRequest proto.InternalMessageInfo

func (m *ListCompaniesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListCompaniesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListCompaniesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListCompaniesRequest) GetRequireOpenJobs() bool {
	if m != nil {
		return m.RequireOpenJobs
	}
	return false
}

// Output only.
//
// The List companies response object.
type ListCompaniesResponse struct {
	// Companies for the current client.
	Companies []*Company `protobuf:"bytes,1,rep,name=companies,proto3" json:"companies,omitempty"`
	// A token to retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// Additional information for the API invocation, such as the request
	// tracking id.
	Metadata             *ResponseMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListCompaniesResponse) Reset()         { *m = ListCompaniesResponse{} }
func (m *ListCompaniesResponse) String() string { return proto.CompactTextString(m) }
func (*ListCompaniesResponse) ProtoMessage()    {}
func (*ListCompaniesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_company_service_1d879da651d4da28, []int{5}
}
func (m *ListCompaniesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCompaniesResponse.Unmarshal(m, b)
}
func (m *ListCompaniesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCompaniesResponse.Marshal(b, m, deterministic)
}
func (dst *ListCompaniesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCompaniesResponse.Merge(dst, src)
}
func (m *ListCompaniesResponse) XXX_Size() int {
	return xxx_messageInfo_ListCompaniesResponse.Size(m)
}
func (m *ListCompaniesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCompaniesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCompaniesResponse proto.InternalMessageInfo

func (m *ListCompaniesResponse) GetCompanies() []*Company {
	if m != nil {
		return m.Companies
	}
	return nil
}

func (m *ListCompaniesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListCompaniesResponse) GetMetadata() *ResponseMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateCompanyRequest)(nil), "google.cloud.talent.v4beta1.CreateCompanyRequest")
	proto.RegisterType((*GetCompanyRequest)(nil), "google.cloud.talent.v4beta1.GetCompanyRequest")
	proto.RegisterType((*UpdateCompanyRequest)(nil), "google.cloud.talent.v4beta1.UpdateCompanyRequest")
	proto.RegisterType((*DeleteCompanyRequest)(nil), "google.cloud.talent.v4beta1.DeleteCompanyRequest")
	proto.RegisterType((*ListCompaniesRequest)(nil), "google.cloud.talent.v4beta1.ListCompaniesRequest")
	proto.RegisterType((*ListCompaniesResponse)(nil), "google.cloud.talent.v4beta1.ListCompaniesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompanyServiceClient interface {
	// Creates a new company entity.
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*Company, error)
	// Retrieves specified company.
	GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*Company, error)
	// Updates specified company.
	UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*Company, error)
	// Deletes specified company.
	// Prerequisite: The company has no jobs associated with it.
	DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists all companies associated with the project.
	ListCompanies(ctx context.Context, in *ListCompaniesRequest, opts ...grpc.CallOption) (*ListCompaniesResponse, error)
}

type companyServiceClient struct {
	cc *grpc.ClientConn
}

func NewCompanyServiceClient(cc *grpc.ClientConn) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/CreateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/GetCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*Company, error) {
	out := new(Company)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/UpdateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/DeleteCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) ListCompanies(ctx context.Context, in *ListCompaniesRequest, opts ...grpc.CallOption) (*ListCompaniesResponse, error) {
	out := new(ListCompaniesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.CompanyService/ListCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServiceServer is the server API for CompanyService service.
type CompanyServiceServer interface {
	// Creates a new company entity.
	CreateCompany(context.Context, *CreateCompanyRequest) (*Company, error)
	// Retrieves specified company.
	GetCompany(context.Context, *GetCompanyRequest) (*Company, error)
	// Updates specified company.
	UpdateCompany(context.Context, *UpdateCompanyRequest) (*Company, error)
	// Deletes specified company.
	// Prerequisite: The company has no jobs associated with it.
	DeleteCompany(context.Context, *DeleteCompanyRequest) (*empty.Empty, error)
	// Lists all companies associated with the project.
	ListCompanies(context.Context, *ListCompaniesRequest) (*ListCompaniesResponse, error)
}

func RegisterCompanyServiceServer(s *grpc.Server, srv CompanyServiceServer) {
	s.RegisterService(&_CompanyService_serviceDesc, srv)
}

func _CompanyService_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_GetCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/GetCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetCompany(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/UpdateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, req.(*UpdateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_DeleteCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).DeleteCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/DeleteCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).DeleteCompany(ctx, req.(*DeleteCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_ListCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCompaniesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).ListCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.CompanyService/ListCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).ListCompanies(ctx, req.(*ListCompaniesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompanyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.talent.v4beta1.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompany",
			Handler:    _CompanyService_CreateCompany_Handler,
		},
		{
			MethodName: "GetCompany",
			Handler:    _CompanyService_GetCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _CompanyService_UpdateCompany_Handler,
		},
		{
			MethodName: "DeleteCompany",
			Handler:    _CompanyService_DeleteCompany_Handler,
		},
		{
			MethodName: "ListCompanies",
			Handler:    _CompanyService_ListCompanies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/talent/v4beta1/company_service.proto",
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/company_service.proto", fileDescriptor_company_service_1d879da651d4da28)
}

var fileDescriptor_company_service_1d879da651d4da28 = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6f, 0xd3, 0x4a,
	0x10, 0xd7, 0xb6, 0x7d, 0x7d, 0xc9, 0x56, 0x79, 0x55, 0xf7, 0xe5, 0x55, 0x51, 0xfa, 0x10, 0x91,
	0x85, 0x4a, 0x08, 0xc2, 0xa6, 0x29, 0x07, 0x44, 0x05, 0x12, 0x2d, 0x7f, 0x04, 0xa2, 0xa2, 0x72,
	0xcb, 0xa5, 0x42, 0xb2, 0xd6, 0xc9, 0xd4, 0xb8, 0x8d, 0x77, 0x5d, 0xef, 0xa6, 0x2a, 0x45, 0x3d,
	0xc0, 0x57, 0xa8, 0xc4, 0x07, 0xe0, 0x1b, 0xf0, 0x05, 0x38, 0x72, 0xe4, 0x00, 0x37, 0x8e, 0x88,
	0x0f, 0x82, 0x6c, 0xaf, 0x93, 0x3a, 0x31, 0x8e, 0xe9, 0xcd, 0x3b, 0x3b, 0xbf, 0x99, 0xdf, 0x6f,
	0x66, 0x67, 0x8c, 0x57, 0x1c, 0xce, 0x9d, 0x1e, 0x18, 0x9d, 0x1e, 0xef, 0x77, 0x0d, 0x49, 0x7b,
	0xc0, 0xa4, 0x71, 0x74, 0xcb, 0x06, 0x49, 0x57, 0x8c, 0x0e, 0xf7, 0x7c, 0xca, 0x5e, 0x5b, 0x02,
	0x82, 0x23, 0xb7, 0x03, 0xba, 0x1f, 0x70, 0xc9, 0xc9, 0x52, 0x0c, 0xd1, 0x23, 0x88, 0x1e, 0x43,
	0x74, 0x05, 0xa9, 0xff, 0xaf, 0xe2, 0x51, 0xdf, 0x35, 0x28, 0x63, 0x5c, 0x52, 0xe9, 0x72, 0x26,
	0x62, 0x68, 0xbd, 0x39, 0x21, 0x9b, 0xc7, 0x99, 0xf2, 0xbc, 0x56, 0x80, 0x97, 0x72, 0x55, 0x7c,
	0x8c, 0xe8, 0x64, 0xf7, 0xf7, 0x0c, 0xf0, 0x7c, 0x99, 0x5c, 0x36, 0x46, 0x2f, 0xf7, 0x5c, 0xe8,
	0x75, 0x2d, 0x8f, 0x8a, 0x83, 0xd8, 0x43, 0x63, 0xb8, 0xba, 0x11, 0x00, 0x95, 0xb0, 0x11, 0x47,
	0x35, 0xe1, 0xb0, 0x0f, 0x42, 0x92, 0x45, 0x3c, 0xeb, 0xd3, 0x00, 0x98, 0xac, 0xa1, 0x06, 0x6a,
	0x96, 0x4d, 0x75, 0x22, 0xf7, 0xf0, 0xdf, 0x2a, 0x7f, 0x6d, 0xaa, 0x81, 0x9a, 0x73, 0xed, 0x2b,
	0x7a, 0x4e, 0x41, 0xf4, 0x24, 0x6a, 0x02, 0xd2, 0xae, 0xe2, 0x85, 0xc7, 0x20, 0x47, 0x92, 0x11,
	0x3c, 0xc3, 0xa8, 0x07, 0x2a, 0x55, 0xf4, 0xad, 0x9d, 0x21, 0x5c, 0x7d, 0xe1, 0x77, 0xc7, 0x99,
	0x9d, 0x63, 0x80, 0x2e, 0xc0, 0x80, 0xac, 0xe1, 0xb9, 0x7e, 0x14, 0x37, 0x2a, 0x83, 0x52, 0x51,
	0x4f, 0x62, 0x24, 0x95, 0xd2, 0x1f, 0x85, 0x95, 0xda, 0xa4, 0xe2, 0xc0, 0xc4, 0xb1, 0x7b, 0xf8,
	0xad, 0xb5, 0x70, 0xf5, 0x01, 0xf4, 0x60, 0x8c, 0x54, 0x96, 0x82, 0xf7, 0x08, 0x57, 0x9f, 0xb9,
	0x42, 0x89, 0x75, 0x41, 0x4c, 0xaa, 0xed, 0x25, 0x8c, 0x7d, 0xea, 0x80, 0x25, 0xf9, 0x01, 0xb0,
	0x88, 0x58, 0xd9, 0x2c, 0x87, 0x96, 0x9d, 0xd0, 0x40, 0x96, 0x70, 0x74, 0xb0, 0x84, 0x7b, 0x02,
	0xb5, 0xe9, 0x06, 0x6a, 0xfe, 0x65, 0x96, 0x42, 0xc3, 0xb6, 0x7b, 0x02, 0xa4, 0x85, 0x17, 0x02,
	0x38, 0xec, 0xbb, 0x01, 0x58, 0xdc, 0x07, 0x66, 0xed, 0x73, 0x5b, 0xd4, 0x66, 0x1a, 0xa8, 0x59,
	0x32, 0xe7, 0xd5, 0xc5, 0x73, 0x1f, 0xd8, 0x53, 0x6e, 0x0b, 0xed, 0x0b, 0xc2, 0xff, 0x8d, 0x10,
	0x13, 0x3e, 0x67, 0x02, 0xc8, 0x3a, 0x2e, 0x77, 0x12, 0x63, 0x0d, 0x35, 0xa6, 0x0b, 0x57, 0x77,
	0x08, 0x23, 0xcb, 0x78, 0x9e, 0xc1, 0xb1, 0xb4, 0xc6, 0xa4, 0x54, 0x42, 0xf3, 0xd6, 0x40, 0xce,
	0x13, 0x5c, 0xf2, 0x40, 0xd2, 0x2e, 0x95, 0x34, 0x52, 0x33, 0xd7, 0xbe, 0x91, 0x9b, 0x2a, 0x21,
	0xb9, 0xa9, 0x40, 0xe6, 0x00, 0xde, 0xfe, 0x58, 0xc2, 0xff, 0x28, 0x26, 0xdb, 0xf1, 0xb0, 0x92,
	0xaf, 0x08, 0x57, 0x52, 0x0f, 0x9b, 0xac, 0xe4, 0x0b, 0xc9, 0x18, 0x82, 0x7a, 0x21, 0xed, 0xda,
	0xab, 0x77, 0xdf, 0x7e, 0x9e, 0x4d, 0xd9, 0xda, 0xcd, 0xc1, 0x84, 0xbe, 0x89, 0x1b, 0x7a, 0xd7,
	0x0f, 0xf8, 0x3e, 0x74, 0xa4, 0x30, 0x5a, 0x86, 0x04, 0x46, 0x59, 0xf8, 0x75, 0x6a, 0x0c, 0x4a,
	0x75, 0x07, 0xb5, 0x76, 0xaf, 0x6b, 0xcb, 0x39, 0xb0, 0xb4, 0x33, 0xf9, 0x8c, 0x30, 0x1e, 0x4e,
	0x0f, 0xd1, 0x73, 0xe9, 0x8d, 0x8d, 0x59, 0x41, 0x39, 0x76, 0x24, 0xe7, 0x25, 0x39, 0x27, 0x27,
	0x7c, 0xce, 0x99, 0x62, 0x86, 0xf4, 0x8c, 0xd6, 0xe9, 0x6e, 0x93, 0x2c, 0xff, 0x1e, 0x73, 0xde,
	0x93, 0xfc, 0x40, 0xb8, 0x92, 0x1a, 0xee, 0x09, 0xdd, 0xc9, 0x5a, 0x04, 0x05, 0xe5, 0x1c, 0x47,
	0x72, 0x82, 0xf6, 0xed, 0x21, 0xb5, 0x64, 0x81, 0x16, 0x93, 0x15, 0x76, 0x69, 0xb5, 0xad, 0x4f,
	0x86, 0x8f, 0x80, 0xc8, 0x27, 0x84, 0x2b, 0xa9, 0x65, 0x31, 0x41, 0x64, 0xd6, 0x62, 0xa9, 0x2f,
	0x8e, 0x2d, 0xa6, 0x87, 0xe1, 0x7e, 0x4f, 0xba, 0xd4, 0xba, 0x40, 0x97, 0x5a, 0x45, 0xbb, 0xf4,
	0x1d, 0xe1, 0x4a, 0x6a, 0x4f, 0x4c, 0x10, 0x90, 0xb5, 0xec, 0xea, 0xed, 0x3f, 0x81, 0xc4, 0x13,
	0x9e, 0xf5, 0x04, 0x8b, 0x4d, 0x54, 0xfa, 0x09, 0xe6, 0x8d, 0xd3, 0xfa, 0x5b, 0x84, 0x2f, 0x77,
	0xb8, 0x97, 0xc7, 0x6e, 0xfd, 0xdf, 0xf4, 0x52, 0xd9, 0x0a, 0x5b, 0xb0, 0x85, 0x76, 0xef, 0x2b,
	0x8c, 0xc3, 0x7b, 0x94, 0x39, 0x3a, 0x0f, 0x1c, 0xc3, 0x01, 0x16, 0x35, 0xc8, 0x88, 0xaf, 0xa8,
	0xef, 0x8a, 0xcc, 0x9f, 0xf7, 0x5a, 0x7c, 0xfc, 0x30, 0x35, 0xbd, 0xb1, 0xb3, 0x6d, 0xcf, 0x46,
	0x98, 0xd5, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x8c, 0xa4, 0x83, 0x87, 0x08, 0x00, 0x00,
}
