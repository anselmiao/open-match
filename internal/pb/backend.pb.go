// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/protobuf-spec/backend.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateMatchRequest struct {
	Match                *MatchObject `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateMatchRequest) Reset()         { *m = CreateMatchRequest{} }
func (m *CreateMatchRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMatchRequest) ProtoMessage()    {}
func (*CreateMatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92161ae1f6f50f7a, []int{0}
}

func (m *CreateMatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMatchRequest.Unmarshal(m, b)
}
func (m *CreateMatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMatchRequest.Marshal(b, m, deterministic)
}
func (m *CreateMatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMatchRequest.Merge(m, src)
}
func (m *CreateMatchRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMatchRequest.Size(m)
}
func (m *CreateMatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMatchRequest proto.InternalMessageInfo

func (m *CreateMatchRequest) GetMatch() *MatchObject {
	if m != nil {
		return m.Match
	}
	return nil
}

type CreateMatchResponse struct {
	Match                *MatchObject `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateMatchResponse) Reset()         { *m = CreateMatchResponse{} }
func (m *CreateMatchResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMatchResponse) ProtoMessage()    {}
func (*CreateMatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92161ae1f6f50f7a, []int{1}
}

func (m *CreateMatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMatchResponse.Unmarshal(m, b)
}
func (m *CreateMatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMatchResponse.Marshal(b, m, deterministic)
}
func (m *CreateMatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMatchResponse.Merge(m, src)
}
func (m *CreateMatchResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMatchResponse.Size(m)
}
func (m *CreateMatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMatchResponse proto.InternalMessageInfo

func (m *CreateMatchResponse) GetMatch() *MatchObject {
	if m != nil {
		return m.Match
	}
	return nil
}

type ListMatchesRequest struct {
	Match                *MatchObject `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListMatchesRequest) Reset()         { *m = ListMatchesRequest{} }
func (m *ListMatchesRequest) String() string { return proto.CompactTextString(m) }
func (*ListMatchesRequest) ProtoMessage()    {}
func (*ListMatchesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92161ae1f6f50f7a, []int{2}
}

func (m *ListMatchesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMatchesRequest.Unmarshal(m, b)
}
func (m *ListMatchesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMatchesRequest.Marshal(b, m, deterministic)
}
func (m *ListMatchesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMatchesRequest.Merge(m, src)
}
func (m *ListMatchesRequest) XXX_Size() int {
	return xxx_messageInfo_ListMatchesRequest.Size(m)
}
func (m *ListMatchesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMatchesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMatchesRequest proto.InternalMessageInfo

func (m *ListMatchesRequest) GetMatch() *MatchObject {
	if m != nil {
		return m.Match
	}
	return nil
}

type ListMatchesResponse struct {
	Match                *MatchObject `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListMatchesResponse) Reset()         { *m = ListMatchesResponse{} }
func (m *ListMatchesResponse) String() string { return proto.CompactTextString(m) }
func (*ListMatchesResponse) ProtoMessage()    {}
func (*ListMatchesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92161ae1f6f50f7a, []int{3}
}

func (m *ListMatchesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMatchesResponse.Unmarshal(m, b)
}
func (m *ListMatchesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMatchesResponse.Marshal(b, m, deterministic)
}
func (m *ListMatchesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMatchesResponse.Merge(m, src)
}
func (m *ListMatchesResponse) XXX_Size() int {
	return xxx_messageInfo_ListMatchesResponse.Size(m)
}
func (m *ListMatchesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMatchesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMatchesResponse proto.InternalMessageInfo

func (m *ListMatchesResponse) GetMatch() *MatchObject {
	if m != nil {
		return m.Match
	}
	return nil
}

type DeleteMatchRequest struct {
	Match                *MatchObject `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteMatchRequest) Reset()         { *m = DeleteMatchRequest{} }
func (m *DeleteMatchRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMatchRequest) ProtoMessage()    {}
func (*DeleteMatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92161ae1f6f50f7a, []int{4}
}

func (m *DeleteMatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMatchRequest.Unmarshal(m, b)
}
func (m *DeleteMatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMatchRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMatchRequest.Merge(m, src)
}
func (m *DeleteMatchRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMatchRequest.Size(m)
}
func (m *DeleteMatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMatchRequest proto.InternalMessageInfo

func (m *DeleteMatchRequest) GetMatch() *MatchObject {
	if m != nil {
		return m.Match
	}
	return nil
}

type DeleteMatchResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMatchResponse) Reset()         { *m = DeleteMatchResponse{} }
func (m *DeleteMatchResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteMatchResponse) ProtoMessage()    {}
func (*DeleteMatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92161ae1f6f50f7a, []int{5}
}

func (m *DeleteMatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMatchResponse.Unmarshal(m, b)
}
func (m *DeleteMatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMatchResponse.Marshal(b, m, deterministic)
}
func (m *DeleteMatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMatchResponse.Merge(m, src)
}
func (m *DeleteMatchResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteMatchResponse.Size(m)
}
func (m *DeleteMatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMatchResponse proto.InternalMessageInfo

type CreateAssignmentsRequest struct {
	Assignment           *Assignments `protobuf:"bytes,1,opt,name=assignment,proto3" json:"assignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateAssignmentsRequest) Reset()         { *m = CreateAssignmentsRequest{} }
func (m *CreateAssignmentsRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAssignmentsRequest) ProtoMessage()    {}
func (*CreateAssignmentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92161ae1f6f50f7a, []int{6}
}

func (m *CreateAssignmentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAssignmentsRequest.Unmarshal(m, b)
}
func (m *CreateAssignmentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAssignmentsRequest.Marshal(b, m, deterministic)
}
func (m *CreateAssignmentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAssignmentsRequest.Merge(m, src)
}
func (m *CreateAssignmentsRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAssignmentsRequest.Size(m)
}
func (m *CreateAssignmentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAssignmentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAssignmentsRequest proto.InternalMessageInfo

func (m *CreateAssignmentsRequest) GetAssignment() *Assignments {
	if m != nil {
		return m.Assignment
	}
	return nil
}

type CreateAssignmentsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAssignmentsResponse) Reset()         { *m = CreateAssignmentsResponse{} }
func (m *CreateAssignmentsResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAssignmentsResponse) ProtoMessage()    {}
func (*CreateAssignmentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92161ae1f6f50f7a, []int{7}
}

func (m *CreateAssignmentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAssignmentsResponse.Unmarshal(m, b)
}
func (m *CreateAssignmentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAssignmentsResponse.Marshal(b, m, deterministic)
}
func (m *CreateAssignmentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAssignmentsResponse.Merge(m, src)
}
func (m *CreateAssignmentsResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAssignmentsResponse.Size(m)
}
func (m *CreateAssignmentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAssignmentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAssignmentsResponse proto.InternalMessageInfo

type DeleteAssignmentsRequest struct {
	Roster               *Roster  `protobuf:"bytes,1,opt,name=roster,proto3" json:"roster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAssignmentsRequest) Reset()         { *m = DeleteAssignmentsRequest{} }
func (m *DeleteAssignmentsRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAssignmentsRequest) ProtoMessage()    {}
func (*DeleteAssignmentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_92161ae1f6f50f7a, []int{8}
}

func (m *DeleteAssignmentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAssignmentsRequest.Unmarshal(m, b)
}
func (m *DeleteAssignmentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAssignmentsRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAssignmentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAssignmentsRequest.Merge(m, src)
}
func (m *DeleteAssignmentsRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAssignmentsRequest.Size(m)
}
func (m *DeleteAssignmentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAssignmentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAssignmentsRequest proto.InternalMessageInfo

func (m *DeleteAssignmentsRequest) GetRoster() *Roster {
	if m != nil {
		return m.Roster
	}
	return nil
}

type DeleteAssignmentsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAssignmentsResponse) Reset()         { *m = DeleteAssignmentsResponse{} }
func (m *DeleteAssignmentsResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteAssignmentsResponse) ProtoMessage()    {}
func (*DeleteAssignmentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92161ae1f6f50f7a, []int{9}
}

func (m *DeleteAssignmentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAssignmentsResponse.Unmarshal(m, b)
}
func (m *DeleteAssignmentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAssignmentsResponse.Marshal(b, m, deterministic)
}
func (m *DeleteAssignmentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAssignmentsResponse.Merge(m, src)
}
func (m *DeleteAssignmentsResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteAssignmentsResponse.Size(m)
}
func (m *DeleteAssignmentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAssignmentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAssignmentsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateMatchRequest)(nil), "api.CreateMatchRequest")
	proto.RegisterType((*CreateMatchResponse)(nil), "api.CreateMatchResponse")
	proto.RegisterType((*ListMatchesRequest)(nil), "api.ListMatchesRequest")
	proto.RegisterType((*ListMatchesResponse)(nil), "api.ListMatchesResponse")
	proto.RegisterType((*DeleteMatchRequest)(nil), "api.DeleteMatchRequest")
	proto.RegisterType((*DeleteMatchResponse)(nil), "api.DeleteMatchResponse")
	proto.RegisterType((*CreateAssignmentsRequest)(nil), "api.CreateAssignmentsRequest")
	proto.RegisterType((*CreateAssignmentsResponse)(nil), "api.CreateAssignmentsResponse")
	proto.RegisterType((*DeleteAssignmentsRequest)(nil), "api.DeleteAssignmentsRequest")
	proto.RegisterType((*DeleteAssignmentsResponse)(nil), "api.DeleteAssignmentsResponse")
}

func init() { proto.RegisterFile("api/protobuf-spec/backend.proto", fileDescriptor_92161ae1f6f50f7a) }

var fileDescriptor_92161ae1f6f50f7a = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0x4b, 0x88, 0x98, 0x0c, 0x17, 0x5d, 0x42, 0xac, 0x18, 0x95, 0xf4, 0x44, 0x62, 0x68,
	0x0d, 0x86, 0x78, 0xa6, 0x90, 0x78, 0xd1, 0xa8, 0x8d, 0x27, 0x6f, 0xdb, 0x32, 0x94, 0x6a, 0xdb,
	0x5d, 0xbb, 0xdb, 0x5f, 0xe0, 0x1f, 0x37, 0xb4, 0x1b, 0x5c, 0xb2, 0xe5, 0x20, 0x1e, 0x99, 0x99,
	0xfd, 0xde, 0x63, 0xde, 0xa4, 0x70, 0x4d, 0x79, 0xe2, 0xf1, 0x82, 0x49, 0x16, 0x96, 0xab, 0xb1,
	0xe0, 0x18, 0x79, 0x21, 0x8d, 0x3e, 0x31, 0x5f, 0xba, 0x55, 0x95, 0xb4, 0x29, 0x4f, 0x06, 0x43,
	0x73, 0x2a, 0x43, 0x21, 0x68, 0x8c, 0xa2, 0x1e, 0x73, 0x66, 0x40, 0xe6, 0x05, 0x52, 0x89, 0x4f,
	0x54, 0x46, 0xeb, 0x00, 0xbf, 0x4a, 0x14, 0x92, 0xdc, 0xc0, 0x51, 0xb6, 0xf9, 0x6d, 0xb7, 0x86,
	0xad, 0x51, 0x77, 0xd2, 0x77, 0xb7, 0xaf, 0xaa, 0xb1, 0xe7, 0xf0, 0x03, 0x23, 0x19, 0xd4, 0x33,
	0x8e, 0x0f, 0xbd, 0x1d, 0x84, 0xe0, 0x2c, 0x17, 0xf8, 0x37, 0xc6, 0x0c, 0xc8, 0x63, 0x22, 0x64,
	0xd5, 0x41, 0x71, 0xa8, 0x8d, 0x1d, 0xc4, 0x81, 0x36, 0x16, 0x98, 0xe2, 0x7f, 0xb6, 0xd1, 0x87,
	0xde, 0x0e, 0xa2, 0xb6, 0xe1, 0xbc, 0x82, 0x5d, 0x2f, 0x69, 0x26, 0x44, 0x12, 0xe7, 0x19, 0xe6,
	0x72, 0xfb, 0x37, 0xa7, 0x00, 0x74, 0x5b, 0x35, 0x45, 0xf4, 0x17, 0xda, 0xa0, 0x73, 0x01, 0xe7,
	0x0d, 0x48, 0xa5, 0xb7, 0x00, 0xbb, 0xb6, 0xd1, 0xa0, 0x37, 0x82, 0x4e, 0xc1, 0x84, 0xc4, 0x42,
	0x69, 0x9d, 0xfc, 0x6a, 0x05, 0x55, 0x3d, 0x50, 0xfd, 0x8d, 0x44, 0x03, 0xa5, 0x96, 0x98, 0x7c,
	0xb7, 0xe1, 0xd8, 0xaf, 0x6f, 0x8e, 0xf8, 0xd0, 0xd5, 0x6e, 0x80, 0x9c, 0xb9, 0x94, 0x27, 0xae,
	0x79, 0x58, 0x03, 0xdb, 0x6c, 0x28, 0xc3, 0x16, 0x59, 0x40, 0x57, 0x0b, 0x50, 0x31, 0xcc, 0xab,
	0x50, 0x8c, 0x86, 0xac, 0x1d, 0xeb, 0xb6, 0xb5, 0x71, 0xa2, 0xed, 0x5f, 0x51, 0xcc, 0x50, 0x15,
	0xa5, 0x29, 0x2a, 0x8b, 0xbc, 0xc1, 0xa9, 0xb1, 0x59, 0x72, 0xa9, 0x59, 0x37, 0x97, 0x3a, 0xb8,
	0xda, 0xd7, 0xd6, 0xa9, 0xc6, 0x32, 0x15, 0x75, 0x5f, 0x54, 0x8a, 0xba, 0x37, 0x03, 0xc7, 0xf2,
	0xef, 0xdf, 0xa7, 0x71, 0x22, 0xd7, 0x65, 0xe8, 0x46, 0x2c, 0xf3, 0x1e, 0x18, 0x8b, 0x53, 0x9c,
	0xa7, 0xac, 0x5c, 0xbe, 0xa4, 0x54, 0xae, 0x58, 0x91, 0x79, 0x8c, 0x63, 0x3e, 0xae, 0x8e, 0xd3,
	0x4b, 0x72, 0x89, 0x45, 0x4e, 0x53, 0x8f, 0x87, 0x61, 0xa7, 0xfa, 0x00, 0xdc, 0xfd, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x5a, 0x8a, 0x1f, 0x10, 0x4a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BackendClient is the client API for Backend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackendClient interface {
	// Run MMF once.  Return a matchobject that fits this profile.
	// INPUT: MatchObject message with these fields populated:
	//  - id
	//  - properties
	//  - [optional] roster, any fields you fill are available to your MMF.
	//  - [optional] pools, any fields you fill are available to your MMF.
	// OUTPUT: MatchObject message with these fields populated:
	//  - id
	//  - properties
	//  - error. Empty if no error was encountered
	//  - rosters, if you choose to fill them in your MMF. (Recommended)
	//  - pools, if you used the MMLogicAPI in your MMF. (Recommended, and provides stats)
	CreateMatch(ctx context.Context, in *CreateMatchRequest, opts ...grpc.CallOption) (*CreateMatchResponse, error)
	// Continually run MMF and stream MatchObjects that fit this profile until
	// the backend client closes the connection.  Same inputs/outputs as CreateMatch.
	ListMatches(ctx context.Context, in *ListMatchesRequest, opts ...grpc.CallOption) (Backend_ListMatchesClient, error)
	// Delete a MatchObject from state storage manually. (MatchObjects in state
	// storage will also automatically expire after a while, defined in the config)
	// INPUT: MatchObject message with the 'id' field populated.
	// (All other fields are ignored.)
	DeleteMatch(ctx context.Context, in *DeleteMatchRequest, opts ...grpc.CallOption) (*DeleteMatchResponse, error)
	// Write the connection info for the list of players in the
	// Assignments.messages.Rosters to state storage.  The Frontend API is
	// responsible for sending anything sent here to the game clients.
	// Sending a player to this function kicks off a process that removes
	// the player from future matchmaking functions by adding them to the
	// 'deindexed' player list and then deleting their player ID from state storage
	// indexes.
	// INPUT: Assignments message with these fields populated:
	//  - assignment, anything you write to this string is sent to Frontend API
	//  - rosters. You can send any number of rosters, containing any number of
	//     player messages. All players from all rosters will be sent the assignment.
	//     The only field in the Roster's Player messages used by CreateAssignments is
	//     the id field.  All other fields in the Player messages are silently ignored.
	CreateAssignments(ctx context.Context, in *CreateAssignmentsRequest, opts ...grpc.CallOption) (*CreateAssignmentsResponse, error)
	// Remove DGS connection info from state storage for players.
	// INPUT: Roster message with the 'players' field populated.
	//    The only field in the Roster's Player messages used by
	//    DeleteAssignments is the 'id' field.  All others are silently ignored.  If
	//    you need to delete multiple rosters, make multiple calls.
	DeleteAssignments(ctx context.Context, in *DeleteAssignmentsRequest, opts ...grpc.CallOption) (*DeleteAssignmentsResponse, error)
}

type backendClient struct {
	cc *grpc.ClientConn
}

func NewBackendClient(cc *grpc.ClientConn) BackendClient {
	return &backendClient{cc}
}

func (c *backendClient) CreateMatch(ctx context.Context, in *CreateMatchRequest, opts ...grpc.CallOption) (*CreateMatchResponse, error) {
	out := new(CreateMatchResponse)
	err := c.cc.Invoke(ctx, "/api.Backend/CreateMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) ListMatches(ctx context.Context, in *ListMatchesRequest, opts ...grpc.CallOption) (Backend_ListMatchesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Backend_serviceDesc.Streams[0], "/api.Backend/ListMatches", opts...)
	if err != nil {
		return nil, err
	}
	x := &backendListMatchesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Backend_ListMatchesClient interface {
	Recv() (*ListMatchesResponse, error)
	grpc.ClientStream
}

type backendListMatchesClient struct {
	grpc.ClientStream
}

func (x *backendListMatchesClient) Recv() (*ListMatchesResponse, error) {
	m := new(ListMatchesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *backendClient) DeleteMatch(ctx context.Context, in *DeleteMatchRequest, opts ...grpc.CallOption) (*DeleteMatchResponse, error) {
	out := new(DeleteMatchResponse)
	err := c.cc.Invoke(ctx, "/api.Backend/DeleteMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) CreateAssignments(ctx context.Context, in *CreateAssignmentsRequest, opts ...grpc.CallOption) (*CreateAssignmentsResponse, error) {
	out := new(CreateAssignmentsResponse)
	err := c.cc.Invoke(ctx, "/api.Backend/CreateAssignments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) DeleteAssignments(ctx context.Context, in *DeleteAssignmentsRequest, opts ...grpc.CallOption) (*DeleteAssignmentsResponse, error) {
	out := new(DeleteAssignmentsResponse)
	err := c.cc.Invoke(ctx, "/api.Backend/DeleteAssignments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendServer is the server API for Backend service.
type BackendServer interface {
	// Run MMF once.  Return a matchobject that fits this profile.
	// INPUT: MatchObject message with these fields populated:
	//  - id
	//  - properties
	//  - [optional] roster, any fields you fill are available to your MMF.
	//  - [optional] pools, any fields you fill are available to your MMF.
	// OUTPUT: MatchObject message with these fields populated:
	//  - id
	//  - properties
	//  - error. Empty if no error was encountered
	//  - rosters, if you choose to fill them in your MMF. (Recommended)
	//  - pools, if you used the MMLogicAPI in your MMF. (Recommended, and provides stats)
	CreateMatch(context.Context, *CreateMatchRequest) (*CreateMatchResponse, error)
	// Continually run MMF and stream MatchObjects that fit this profile until
	// the backend client closes the connection.  Same inputs/outputs as CreateMatch.
	ListMatches(*ListMatchesRequest, Backend_ListMatchesServer) error
	// Delete a MatchObject from state storage manually. (MatchObjects in state
	// storage will also automatically expire after a while, defined in the config)
	// INPUT: MatchObject message with the 'id' field populated.
	// (All other fields are ignored.)
	DeleteMatch(context.Context, *DeleteMatchRequest) (*DeleteMatchResponse, error)
	// Write the connection info for the list of players in the
	// Assignments.messages.Rosters to state storage.  The Frontend API is
	// responsible for sending anything sent here to the game clients.
	// Sending a player to this function kicks off a process that removes
	// the player from future matchmaking functions by adding them to the
	// 'deindexed' player list and then deleting their player ID from state storage
	// indexes.
	// INPUT: Assignments message with these fields populated:
	//  - assignment, anything you write to this string is sent to Frontend API
	//  - rosters. You can send any number of rosters, containing any number of
	//     player messages. All players from all rosters will be sent the assignment.
	//     The only field in the Roster's Player messages used by CreateAssignments is
	//     the id field.  All other fields in the Player messages are silently ignored.
	CreateAssignments(context.Context, *CreateAssignmentsRequest) (*CreateAssignmentsResponse, error)
	// Remove DGS connection info from state storage for players.
	// INPUT: Roster message with the 'players' field populated.
	//    The only field in the Roster's Player messages used by
	//    DeleteAssignments is the 'id' field.  All others are silently ignored.  If
	//    you need to delete multiple rosters, make multiple calls.
	DeleteAssignments(context.Context, *DeleteAssignmentsRequest) (*DeleteAssignmentsResponse, error)
}

// UnimplementedBackendServer can be embedded to have forward compatible implementations.
type UnimplementedBackendServer struct {
}

func (*UnimplementedBackendServer) CreateMatch(ctx context.Context, req *CreateMatchRequest) (*CreateMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMatch not implemented")
}
func (*UnimplementedBackendServer) ListMatches(req *ListMatchesRequest, srv Backend_ListMatchesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListMatches not implemented")
}
func (*UnimplementedBackendServer) DeleteMatch(ctx context.Context, req *DeleteMatchRequest) (*DeleteMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMatch not implemented")
}
func (*UnimplementedBackendServer) CreateAssignments(ctx context.Context, req *CreateAssignmentsRequest) (*CreateAssignmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAssignments not implemented")
}
func (*UnimplementedBackendServer) DeleteAssignments(ctx context.Context, req *DeleteAssignmentsRequest) (*DeleteAssignmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAssignments not implemented")
}

func RegisterBackendServer(s *grpc.Server, srv BackendServer) {
	s.RegisterService(&_Backend_serviceDesc, srv)
}

func _Backend_CreateMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).CreateMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Backend/CreateMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).CreateMatch(ctx, req.(*CreateMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_ListMatches_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListMatchesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackendServer).ListMatches(m, &backendListMatchesServer{stream})
}

type Backend_ListMatchesServer interface {
	Send(*ListMatchesResponse) error
	grpc.ServerStream
}

type backendListMatchesServer struct {
	grpc.ServerStream
}

func (x *backendListMatchesServer) Send(m *ListMatchesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Backend_DeleteMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).DeleteMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Backend/DeleteMatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).DeleteMatch(ctx, req.(*DeleteMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_CreateAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAssignmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).CreateAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Backend/CreateAssignments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).CreateAssignments(ctx, req.(*CreateAssignmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_DeleteAssignments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAssignmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).DeleteAssignments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Backend/DeleteAssignments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).DeleteAssignments(ctx, req.(*DeleteAssignmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Backend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Backend",
	HandlerType: (*BackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMatch",
			Handler:    _Backend_CreateMatch_Handler,
		},
		{
			MethodName: "DeleteMatch",
			Handler:    _Backend_DeleteMatch_Handler,
		},
		{
			MethodName: "CreateAssignments",
			Handler:    _Backend_CreateAssignments_Handler,
		},
		{
			MethodName: "DeleteAssignments",
			Handler:    _Backend_DeleteAssignments_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListMatches",
			Handler:       _Backend_ListMatches_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/protobuf-spec/backend.proto",
}
