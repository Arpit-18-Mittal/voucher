// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/admin/v1beta2/operation.proto

package admin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Describes the state of the operation.
type OperationState int32

const (
	// Unspecified.
	OperationState_OPERATION_STATE_UNSPECIFIED OperationState = 0
	// Request is being prepared for processing.
	OperationState_INITIALIZING OperationState = 1
	// Request is actively being processed.
	OperationState_PROCESSING OperationState = 2
	// Request is in the process of being cancelled after user called
	// google.longrunning.Operations.CancelOperation on the operation.
	OperationState_CANCELLING OperationState = 3
	// Request has been processed and is in its finalization stage.
	OperationState_FINALIZING OperationState = 4
	// Request has completed successfully.
	OperationState_SUCCESSFUL OperationState = 5
	// Request has finished being processed, but encountered an error.
	OperationState_FAILED OperationState = 6
	// Request has finished being cancelled after user called
	// google.longrunning.Operations.CancelOperation.
	OperationState_CANCELLED OperationState = 7
)

var OperationState_name = map[int32]string{
	0: "OPERATION_STATE_UNSPECIFIED",
	1: "INITIALIZING",
	2: "PROCESSING",
	3: "CANCELLING",
	4: "FINALIZING",
	5: "SUCCESSFUL",
	6: "FAILED",
	7: "CANCELLED",
}

var OperationState_value = map[string]int32{
	"OPERATION_STATE_UNSPECIFIED": 0,
	"INITIALIZING":                1,
	"PROCESSING":                  2,
	"CANCELLING":                  3,
	"FINALIZING":                  4,
	"SUCCESSFUL":                  5,
	"FAILED":                      6,
	"CANCELLED":                   7,
}

func (x OperationState) String() string {
	return proto.EnumName(OperationState_name, int32(x))
}

func (OperationState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f946ae2a57533608, []int{0}
}

// Specifies how the index is changing.
type FieldOperationMetadata_IndexConfigDelta_ChangeType int32

const (
	// The type of change is not specified or known.
	FieldOperationMetadata_IndexConfigDelta_CHANGE_TYPE_UNSPECIFIED FieldOperationMetadata_IndexConfigDelta_ChangeType = 0
	// The single field index is being added.
	FieldOperationMetadata_IndexConfigDelta_ADD FieldOperationMetadata_IndexConfigDelta_ChangeType = 1
	// The single field index is being removed.
	FieldOperationMetadata_IndexConfigDelta_REMOVE FieldOperationMetadata_IndexConfigDelta_ChangeType = 2
)

var FieldOperationMetadata_IndexConfigDelta_ChangeType_name = map[int32]string{
	0: "CHANGE_TYPE_UNSPECIFIED",
	1: "ADD",
	2: "REMOVE",
}

var FieldOperationMetadata_IndexConfigDelta_ChangeType_value = map[string]int32{
	"CHANGE_TYPE_UNSPECIFIED": 0,
	"ADD":                     1,
	"REMOVE":                  2,
}

func (x FieldOperationMetadata_IndexConfigDelta_ChangeType) String() string {
	return proto.EnumName(FieldOperationMetadata_IndexConfigDelta_ChangeType_name, int32(x))
}

func (FieldOperationMetadata_IndexConfigDelta_ChangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f946ae2a57533608, []int{0, 0, 0}
}

// Metadata for [google.longrunning.Operation][google.longrunning.Operation] results from
// [FirestoreAdmin.UpdateField][google.firestore.admin.v1beta2.FirestoreAdmin.UpdateField].
type FieldOperationMetadata struct {
	// The time this operation started.
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The time this operation completed. Will be unset if operation still in
	// progress.
	EndTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// The field resource that this operation is acting on. For example:
	// `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/fields/{field_path}`
	Field string `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
	// A list of [IndexConfigDelta][google.firestore.admin.v1beta2.FieldOperationMetadata.IndexConfigDelta], which describe the intent of this
	// operation.
	IndexConfigDeltas []*FieldOperationMetadata_IndexConfigDelta `protobuf:"bytes,4,rep,name=index_config_deltas,json=indexConfigDeltas,proto3" json:"index_config_deltas,omitempty"`
	// The state of the operation.
	State OperationState `protobuf:"varint,5,opt,name=state,proto3,enum=google.firestore.admin.v1beta2.OperationState" json:"state,omitempty"`
	// The progress, in documents, of this operation.
	DocumentProgress *Progress `protobuf:"bytes,6,opt,name=document_progress,json=documentProgress,proto3" json:"document_progress,omitempty"`
	// The progress, in bytes, of this operation.
	BytesProgress        *Progress `protobuf:"bytes,7,opt,name=bytes_progress,json=bytesProgress,proto3" json:"bytes_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FieldOperationMetadata) Reset()         { *m = FieldOperationMetadata{} }
func (m *FieldOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*FieldOperationMetadata) ProtoMessage()    {}
func (*FieldOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f946ae2a57533608, []int{0}
}

func (m *FieldOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldOperationMetadata.Unmarshal(m, b)
}
func (m *FieldOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldOperationMetadata.Marshal(b, m, deterministic)
}
func (m *FieldOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldOperationMetadata.Merge(m, src)
}
func (m *FieldOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_FieldOperationMetadata.Size(m)
}
func (m *FieldOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_FieldOperationMetadata proto.InternalMessageInfo

func (m *FieldOperationMetadata) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *FieldOperationMetadata) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *FieldOperationMetadata) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *FieldOperationMetadata) GetIndexConfigDeltas() []*FieldOperationMetadata_IndexConfigDelta {
	if m != nil {
		return m.IndexConfigDeltas
	}
	return nil
}

func (m *FieldOperationMetadata) GetState() OperationState {
	if m != nil {
		return m.State
	}
	return OperationState_OPERATION_STATE_UNSPECIFIED
}

func (m *FieldOperationMetadata) GetDocumentProgress() *Progress {
	if m != nil {
		return m.DocumentProgress
	}
	return nil
}

func (m *FieldOperationMetadata) GetBytesProgress() *Progress {
	if m != nil {
		return m.BytesProgress
	}
	return nil
}

// Information about an index configuration change.
type FieldOperationMetadata_IndexConfigDelta struct {
	// Specifies how the index is changing.
	ChangeType FieldOperationMetadata_IndexConfigDelta_ChangeType `protobuf:"varint,1,opt,name=change_type,json=changeType,proto3,enum=google.firestore.admin.v1beta2.FieldOperationMetadata_IndexConfigDelta_ChangeType" json:"change_type,omitempty"`
	// The index being changed.
	Index                *Index   `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldOperationMetadata_IndexConfigDelta) Reset() {
	*m = FieldOperationMetadata_IndexConfigDelta{}
}
func (m *FieldOperationMetadata_IndexConfigDelta) String() string { return proto.CompactTextString(m) }
func (*FieldOperationMetadata_IndexConfigDelta) ProtoMessage()    {}
func (*FieldOperationMetadata_IndexConfigDelta) Descriptor() ([]byte, []int) {
	return fileDescriptor_f946ae2a57533608, []int{0, 0}
}

func (m *FieldOperationMetadata_IndexConfigDelta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldOperationMetadata_IndexConfigDelta.Unmarshal(m, b)
}
func (m *FieldOperationMetadata_IndexConfigDelta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldOperationMetadata_IndexConfigDelta.Marshal(b, m, deterministic)
}
func (m *FieldOperationMetadata_IndexConfigDelta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldOperationMetadata_IndexConfigDelta.Merge(m, src)
}
func (m *FieldOperationMetadata_IndexConfigDelta) XXX_Size() int {
	return xxx_messageInfo_FieldOperationMetadata_IndexConfigDelta.Size(m)
}
func (m *FieldOperationMetadata_IndexConfigDelta) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldOperationMetadata_IndexConfigDelta.DiscardUnknown(m)
}

var xxx_messageInfo_FieldOperationMetadata_IndexConfigDelta proto.InternalMessageInfo

func (m *FieldOperationMetadata_IndexConfigDelta) GetChangeType() FieldOperationMetadata_IndexConfigDelta_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return FieldOperationMetadata_IndexConfigDelta_CHANGE_TYPE_UNSPECIFIED
}

func (m *FieldOperationMetadata_IndexConfigDelta) GetIndex() *Index {
	if m != nil {
		return m.Index
	}
	return nil
}

// Describes the progress of the operation.
// Unit of work is generic and must be interpreted based on where [Progress][google.firestore.admin.v1beta2.Progress]
// is used.
type Progress struct {
	// The amount of work estimated.
	EstimatedWork int64 `protobuf:"varint,1,opt,name=estimated_work,json=estimatedWork,proto3" json:"estimated_work,omitempty"`
	// The amount of work completed.
	CompletedWork        int64    `protobuf:"varint,2,opt,name=completed_work,json=completedWork,proto3" json:"completed_work,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Progress) Reset()         { *m = Progress{} }
func (m *Progress) String() string { return proto.CompactTextString(m) }
func (*Progress) ProtoMessage()    {}
func (*Progress) Descriptor() ([]byte, []int) {
	return fileDescriptor_f946ae2a57533608, []int{1}
}

func (m *Progress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Progress.Unmarshal(m, b)
}
func (m *Progress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Progress.Marshal(b, m, deterministic)
}
func (m *Progress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Progress.Merge(m, src)
}
func (m *Progress) XXX_Size() int {
	return xxx_messageInfo_Progress.Size(m)
}
func (m *Progress) XXX_DiscardUnknown() {
	xxx_messageInfo_Progress.DiscardUnknown(m)
}

var xxx_messageInfo_Progress proto.InternalMessageInfo

func (m *Progress) GetEstimatedWork() int64 {
	if m != nil {
		return m.EstimatedWork
	}
	return 0
}

func (m *Progress) GetCompletedWork() int64 {
	if m != nil {
		return m.CompletedWork
	}
	return 0
}

func init() {
	proto.RegisterType((*FieldOperationMetadata)(nil), "google.firestore.admin.v1beta2.FieldOperationMetadata")
	proto.RegisterType((*FieldOperationMetadata_IndexConfigDelta)(nil), "google.firestore.admin.v1beta2.FieldOperationMetadata.IndexConfigDelta")
	proto.RegisterType((*Progress)(nil), "google.firestore.admin.v1beta2.Progress")
	proto.RegisterEnum("google.firestore.admin.v1beta2.OperationState", OperationState_name, OperationState_value)
	proto.RegisterEnum("google.firestore.admin.v1beta2.FieldOperationMetadata_IndexConfigDelta_ChangeType", FieldOperationMetadata_IndexConfigDelta_ChangeType_name, FieldOperationMetadata_IndexConfigDelta_ChangeType_value)
}

func init() {
	proto.RegisterFile("google/firestore/admin/v1beta2/operation.proto", fileDescriptor_f946ae2a57533608)
}

var fileDescriptor_f946ae2a57533608 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4f, 0x6f, 0x12, 0x4f,
	0x1c, 0xc6, 0x7f, 0x0b, 0x05, 0xda, 0x6f, 0x7f, 0x25, 0xdb, 0xd1, 0x28, 0xa1, 0xc6, 0x12, 0x62,
	0x13, 0xd2, 0xc3, 0x6c, 0x8a, 0xf1, 0x60, 0x9a, 0x98, 0x6c, 0x97, 0x05, 0x37, 0xa1, 0x40, 0x16,
	0xa8, 0xda, 0xcb, 0x66, 0x60, 0x87, 0x75, 0x53, 0x76, 0x67, 0xb3, 0x3b, 0xb5, 0xf6, 0xee, 0xd9,
	0x97, 0xe0, 0xc1, 0x9b, 0xbe, 0x4a, 0x33, 0xb3, 0x7f, 0x1a, 0x1b, 0x95, 0x98, 0x78, 0xe3, 0x19,
	0x9e, 0xe7, 0xf3, 0xfd, 0xc3, 0x0c, 0x80, 0x3d, 0xc6, 0xbc, 0x35, 0xd5, 0x56, 0x7e, 0x4c, 0x13,
	0xce, 0x62, 0xaa, 0x11, 0x37, 0xf0, 0x43, 0xed, 0xc3, 0xc9, 0x82, 0x72, 0xd2, 0xd5, 0x58, 0x44,
	0x63, 0xc2, 0x7d, 0x16, 0xe2, 0x28, 0x66, 0x9c, 0xa1, 0xa7, 0xa9, 0x1f, 0x17, 0x7e, 0x2c, 0xfd,
	0x38, 0xf3, 0x37, 0x9f, 0x64, 0x3c, 0x12, 0xf9, 0x1a, 0x09, 0x43, 0xc6, 0x65, 0x38, 0x49, 0xd3,
	0xcd, 0xe3, 0x0d, 0xd5, 0xfc, 0xd0, 0xa5, 0x1f, 0x33, 0xef, 0x61, 0xe6, 0x95, 0x6a, 0x71, 0xbd,
	0xd2, 0xb8, 0x1f, 0xd0, 0x84, 0x93, 0x20, 0x4a, 0x0d, 0xed, 0xcf, 0x55, 0x78, 0xd4, 0xf7, 0xe9,
	0xda, 0x1d, 0xe7, 0x3d, 0x9e, 0x53, 0x4e, 0x5c, 0xc2, 0x09, 0x7a, 0x09, 0x90, 0x70, 0x12, 0x73,
	0x47, 0x64, 0x1a, 0x4a, 0x4b, 0xe9, 0xec, 0x76, 0x9b, 0xd9, 0xa8, 0x38, 0x07, 0xe2, 0x59, 0x0e,
	0xb4, 0x77, 0xa4, 0x5b, 0x68, 0xf4, 0x02, 0xb6, 0x69, 0xe8, 0xa6, 0xc1, 0xd2, 0xc6, 0x60, 0x8d,
	0x86, 0xae, 0x8c, 0x3d, 0x84, 0xca, 0x4a, 0xf4, 0xd2, 0x28, 0xb7, 0x94, 0xce, 0x8e, 0x9d, 0x0a,
	0x74, 0x03, 0x0f, 0xe4, 0x48, 0xce, 0x92, 0x85, 0x2b, 0xdf, 0x73, 0x5c, 0xba, 0xe6, 0x24, 0x69,
	0x6c, 0xb5, 0xca, 0x9d, 0xdd, 0xee, 0x00, 0xff, 0x79, 0x97, 0xf8, 0xd7, 0xc3, 0x61, 0x4b, 0x10,
	0x0d, 0x09, 0xec, 0x09, 0x9e, 0xbd, 0xef, 0xdf, 0x3b, 0x49, 0x50, 0x0f, 0x2a, 0x09, 0x27, 0x9c,
	0x36, 0x2a, 0x2d, 0xa5, 0x53, 0xef, 0xe2, 0x4d, 0xa5, 0x8a, 0x2a, 0x53, 0x91, 0xb2, 0xd3, 0x30,
	0x9a, 0xc3, 0xbe, 0xcb, 0x96, 0xd7, 0x01, 0x0d, 0xb9, 0x13, 0xc5, 0xcc, 0x8b, 0x69, 0x92, 0x34,
	0xaa, 0x72, 0x29, 0x9d, 0x4d, 0xc4, 0x49, 0xe6, 0xb7, 0xd5, 0x1c, 0x91, 0x9f, 0xa0, 0x31, 0xd4,
	0x17, 0xb7, 0x9c, 0x26, 0x77, 0xcc, 0xda, 0x5f, 0x32, 0xf7, 0x64, 0x3e, 0x97, 0xcd, 0x4f, 0x25,
	0x50, 0xef, 0x6f, 0x05, 0x25, 0xb0, 0xbb, 0x7c, 0x4f, 0x42, 0x8f, 0x3a, 0xfc, 0x36, 0x4a, 0x2f,
	0x41, 0xbd, 0x6b, 0xff, 0xa3, 0x9d, 0x63, 0x43, 0xa2, 0x67, 0xb7, 0x11, 0xb5, 0x61, 0x59, 0x7c,
	0x46, 0xa7, 0x50, 0x91, 0x3f, 0x46, 0x76, 0x75, 0x8e, 0x36, 0x95, 0x93, 0x5c, 0x3b, 0xcd, 0xb4,
	0x5f, 0x01, 0xdc, 0x61, 0xd1, 0x01, 0x3c, 0x36, 0x5e, 0xeb, 0xa3, 0x81, 0xe9, 0xcc, 0xde, 0x4d,
	0x4c, 0x67, 0x3e, 0x9a, 0x4e, 0x4c, 0xc3, 0xea, 0x5b, 0x66, 0x4f, 0xfd, 0x0f, 0xd5, 0xa0, 0xac,
	0xf7, 0x7a, 0xaa, 0x82, 0x00, 0xaa, 0xb6, 0x79, 0x3e, 0xbe, 0x30, 0xd5, 0x52, 0xfb, 0x2d, 0x6c,
	0x17, 0x3b, 0x3e, 0x82, 0x3a, 0x4d, 0xb8, 0x1f, 0x10, 0x4e, 0x5d, 0xe7, 0x86, 0xc5, 0x57, 0x72,
	0x01, 0x65, 0x7b, 0xaf, 0x38, 0x7d, 0xc3, 0xe2, 0x2b, 0x61, 0x5b, 0xb2, 0x20, 0x5a, 0xd3, 0xc2,
	0x56, 0x4a, 0x6d, 0xc5, 0xa9, 0xb0, 0x1d, 0x7f, 0x51, 0xa0, 0xfe, 0xf3, 0x15, 0x41, 0x87, 0x70,
	0x30, 0x9e, 0x98, 0xb6, 0x3e, 0xb3, 0xc6, 0x23, 0x67, 0x3a, 0xd3, 0x67, 0xf7, 0x5b, 0x54, 0xe1,
	0x7f, 0x6b, 0x64, 0xcd, 0x2c, 0x7d, 0x68, 0x5d, 0x5a, 0xa3, 0x81, 0xaa, 0xa0, 0x3a, 0xc0, 0xc4,
	0x1e, 0x1b, 0xe6, 0x74, 0x2a, 0x74, 0x49, 0x68, 0x43, 0x1f, 0x19, 0xe6, 0x70, 0x28, 0x74, 0x59,
	0xe8, 0xbe, 0x35, 0xca, 0xfd, 0x5b, 0x42, 0x4f, 0xe7, 0x86, 0xf0, 0xf7, 0xe7, 0x43, 0xb5, 0x22,
	0x66, 0xed, 0xeb, 0xd6, 0xd0, 0xec, 0xa9, 0x55, 0xb4, 0x07, 0x3b, 0x59, 0xd6, 0xec, 0xa9, 0xb5,
	0xb3, 0x6f, 0x0a, 0xb4, 0x97, 0x2c, 0xd8, 0xb0, 0xee, 0xb3, 0xbb, 0x21, 0x26, 0xe2, 0x2d, 0x4f,
	0x94, 0x4b, 0x23, 0x4b, 0x78, 0x6c, 0x4d, 0x42, 0x0f, 0xb3, 0xd8, 0xd3, 0x3c, 0x1a, 0xca, 0x97,
	0xae, 0xa5, 0x5f, 0x91, 0xc8, 0x4f, 0x7e, 0xf7, 0x87, 0x75, 0x2a, 0xd5, 0xd7, 0xd2, 0xd6, 0xc0,
	0xe8, 0x4f, 0xbf, 0x97, 0x9e, 0x0d, 0x52, 0x98, 0xb1, 0x66, 0xd7, 0x2e, 0xee, 0x17, 0x4d, 0xe8,
	0xb2, 0x89, 0x8b, 0x93, 0x33, 0x91, 0x59, 0x54, 0x25, 0xfd, 0xf9, 0x8f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x2e, 0x32, 0x9e, 0x39, 0x7b, 0x05, 0x00, 0x00,
}
