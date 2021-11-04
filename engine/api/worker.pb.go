// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: engine/api/worker.proto

package api

import (
	api "github.com/vynaloze/mapreduce/api"
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

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_api_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_engine_api_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_engine_api_worker_proto_rawDescGZIP(), []int{0}
}

func (x *Key) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_api_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_engine_api_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_engine_api_worker_proto_rawDescGZIP(), []int{1}
}

func (x *Value) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *Key   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_api_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_engine_api_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_engine_api_worker_proto_rawDescGZIP(), []int{2}
}

func (x *Pair) GetKey() *Key {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Pair) GetValue() *Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type Split struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *api.DFSFile `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Offset int64        `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64        `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Split) Reset() {
	*x = Split{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_api_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Split) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Split) ProtoMessage() {}

func (x *Split) ProtoReflect() protoreflect.Message {
	mi := &file_engine_api_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Split.ProtoReflect.Descriptor instead.
func (*Split) Descriptor() ([]byte, []int) {
	return file_engine_api_worker_proto_rawDescGZIP(), []int{3}
}

func (x *Split) GetSource() *api.DFSFile {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Split) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Split) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr      string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Partition int64  `protobuf:"varint,2,opt,name=partition,proto3" json:"partition,omitempty"`
	TaskId    string `protobuf:"bytes,3,opt,name=taskId,proto3" json:"taskId,omitempty"`
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_api_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_engine_api_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_engine_api_worker_proto_rawDescGZIP(), []int{4}
}

func (x *Region) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Region) GetPartition() int64 {
	if x != nil {
		return x.Partition
	}
	return 0
}

func (x *Region) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type MapTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InputSplit *Split `protobuf:"bytes,2,opt,name=input_split,json=inputSplit,proto3" json:"input_split,omitempty"`
	Partitions int64  `protobuf:"varint,3,opt,name=partitions,proto3" json:"partitions,omitempty"` //TODO mapper???
}

func (x *MapTask) Reset() {
	*x = MapTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_api_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapTask) ProtoMessage() {}

func (x *MapTask) ProtoReflect() protoreflect.Message {
	mi := &file_engine_api_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapTask.ProtoReflect.Descriptor instead.
func (*MapTask) Descriptor() ([]byte, []int) {
	return file_engine_api_worker_proto_rawDescGZIP(), []int{5}
}

func (x *MapTask) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MapTask) GetInputSplit() *Split {
	if x != nil {
		return x.InputSplit
	}
	return nil
}

func (x *MapTask) GetPartitions() int64 {
	if x != nil {
		return x.Partitions
	}
	return 0
}

type MapTaskStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Task   *MapTask `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Region *Region  `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"` // stats?
}

func (x *MapTaskStatus) Reset() {
	*x = MapTaskStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_api_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapTaskStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapTaskStatus) ProtoMessage() {}

func (x *MapTaskStatus) ProtoReflect() protoreflect.Message {
	mi := &file_engine_api_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapTaskStatus.ProtoReflect.Descriptor instead.
func (*MapTaskStatus) Descriptor() ([]byte, []int) {
	return file_engine_api_worker_proto_rawDescGZIP(), []int{6}
}

func (x *MapTaskStatus) GetTask() *MapTask {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *MapTaskStatus) GetRegion() *Region {
	if x != nil {
		return x.Region
	}
	return nil
}

type ReduceTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutputSpec *api.OutputSpec `protobuf:"bytes,1,opt,name=output_spec,json=outputSpec,proto3" json:"output_spec,omitempty"`
	Partition  int64           `protobuf:"varint,2,opt,name=partition,proto3" json:"partition,omitempty"` //TODO reducer???
}

func (x *ReduceTask) Reset() {
	*x = ReduceTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_api_worker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceTask) ProtoMessage() {}

func (x *ReduceTask) ProtoReflect() protoreflect.Message {
	mi := &file_engine_api_worker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceTask.ProtoReflect.Descriptor instead.
func (*ReduceTask) Descriptor() ([]byte, []int) {
	return file_engine_api_worker_proto_rawDescGZIP(), []int{7}
}

func (x *ReduceTask) GetOutputSpec() *api.OutputSpec {
	if x != nil {
		return x.OutputSpec
	}
	return nil
}

func (x *ReduceTask) GetPartition() int64 {
	if x != nil {
		return x.Partition
	}
	return 0
}

type ReduceTaskStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *api.DFSFile `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"` // stats?
}

func (x *ReduceTaskStatus) Reset() {
	*x = ReduceTaskStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_api_worker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReduceTaskStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReduceTaskStatus) ProtoMessage() {}

func (x *ReduceTaskStatus) ProtoReflect() protoreflect.Message {
	mi := &file_engine_api_worker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReduceTaskStatus.ProtoReflect.Descriptor instead.
func (*ReduceTaskStatus) Descriptor() ([]byte, []int) {
	return file_engine_api_worker_proto_rawDescGZIP(), []int{8}
}

func (x *ReduceTaskStatus) GetResult() *api.DFSFile {
	if x != nil {
		return x.Result
	}
	return nil
}

type MissingRegions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regions []*Region `protobuf:"bytes,1,rep,name=regions,proto3" json:"regions,omitempty"`
}

func (x *MissingRegions) Reset() {
	*x = MissingRegions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_api_worker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MissingRegions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MissingRegions) ProtoMessage() {}

func (x *MissingRegions) ProtoReflect() protoreflect.Message {
	mi := &file_engine_api_worker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MissingRegions.ProtoReflect.Descriptor instead.
func (*MissingRegions) Descriptor() ([]byte, []int) {
	return file_engine_api_worker_proto_rawDescGZIP(), []int{9}
}

func (x *MissingRegions) GetRegions() []*Region {
	if x != nil {
		return x.Regions
	}
	return nil
}

var File_engine_api_worker_proto protoreflect.FileDescriptor

var file_engine_api_worker_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x61, 0x70, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17,
	0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x1d, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3c, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x16,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4b, 0x65,
	0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x57, 0x0a, 0x05, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x12, 0x20, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x44, 0x46, 0x53, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x52, 0x0a,
	0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73,
	0x6b, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x22, 0x62, 0x0a, 0x07, 0x4d, 0x61, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0b,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x06, 0x2e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x53, 0x70, 0x6c, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4e, 0x0a, 0x0d, 0x4d, 0x61, 0x70, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x61, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04,
	0x74, 0x61, 0x73, 0x6b, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x58, 0x0a, 0x0a, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x2c, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x34, 0x0a, 0x10, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x44, 0x46, 0x53, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x33, 0x0a, 0x0e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x4b, 0x0a, 0x09, 0x4d, 0x61,
	0x70, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x03, 0x4d, 0x61, 0x70, 0x12, 0x08,
	0x2e, 0x4d, 0x61, 0x70, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x0e, 0x2e, 0x4d, 0x61, 0x70, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x19, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x07, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x1a, 0x05, 0x2e, 0x50,
	0x61, 0x69, 0x72, 0x22, 0x00, 0x30, 0x01, 0x32, 0x64, 0x0a, 0x0c, 0x52, 0x65, 0x64, 0x75, 0x63,
	0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x06, 0x52, 0x65, 0x64, 0x75, 0x63,
	0x65, 0x12, 0x0b, 0x2e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x11,
	0x2e, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x26, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x07, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x28, 0x01, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x79, 0x6e, 0x61,
	0x6c, 0x6f, 0x7a, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2f, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_engine_api_worker_proto_rawDescOnce sync.Once
	file_engine_api_worker_proto_rawDescData = file_engine_api_worker_proto_rawDesc
)

func file_engine_api_worker_proto_rawDescGZIP() []byte {
	file_engine_api_worker_proto_rawDescOnce.Do(func() {
		file_engine_api_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_engine_api_worker_proto_rawDescData)
	})
	return file_engine_api_worker_proto_rawDescData
}

var file_engine_api_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_engine_api_worker_proto_goTypes = []interface{}{
	(*Key)(nil),              // 0: Key
	(*Value)(nil),            // 1: Value
	(*Pair)(nil),             // 2: Pair
	(*Split)(nil),            // 3: Split
	(*Region)(nil),           // 4: Region
	(*MapTask)(nil),          // 5: MapTask
	(*MapTaskStatus)(nil),    // 6: MapTaskStatus
	(*ReduceTask)(nil),       // 7: ReduceTask
	(*ReduceTaskStatus)(nil), // 8: ReduceTaskStatus
	(*MissingRegions)(nil),   // 9: MissingRegions
	(*api.DFSFile)(nil),      // 10: DFSFile
	(*api.OutputSpec)(nil),   // 11: OutputSpec
}
var file_engine_api_worker_proto_depIdxs = []int32{
	0,  // 0: Pair.key:type_name -> Key
	1,  // 1: Pair.value:type_name -> Value
	10, // 2: Split.source:type_name -> DFSFile
	3,  // 3: MapTask.input_split:type_name -> Split
	5,  // 4: MapTaskStatus.task:type_name -> MapTask
	4,  // 5: MapTaskStatus.region:type_name -> Region
	11, // 6: ReduceTask.output_spec:type_name -> OutputSpec
	10, // 7: ReduceTaskStatus.result:type_name -> DFSFile
	4,  // 8: MissingRegions.regions:type_name -> Region
	5,  // 9: MapWorker.Map:input_type -> MapTask
	4,  // 10: MapWorker.Get:input_type -> Region
	7,  // 11: ReduceWorker.Reduce:input_type -> ReduceTask
	4,  // 12: ReduceWorker.Notify:input_type -> Region
	6,  // 13: MapWorker.Map:output_type -> MapTaskStatus
	2,  // 14: MapWorker.Get:output_type -> Pair
	8,  // 15: ReduceWorker.Reduce:output_type -> ReduceTaskStatus
	9,  // 16: ReduceWorker.Notify:output_type -> MissingRegions
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_engine_api_worker_proto_init() }
func file_engine_api_worker_proto_init() {
	if File_engine_api_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_engine_api_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_engine_api_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_engine_api_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
		file_engine_api_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Split); i {
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
		file_engine_api_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
		file_engine_api_worker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapTask); i {
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
		file_engine_api_worker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapTaskStatus); i {
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
		file_engine_api_worker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReduceTask); i {
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
		file_engine_api_worker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReduceTaskStatus); i {
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
		file_engine_api_worker_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MissingRegions); i {
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
			RawDescriptor: file_engine_api_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_engine_api_worker_proto_goTypes,
		DependencyIndexes: file_engine_api_worker_proto_depIdxs,
		MessageInfos:      file_engine_api_worker_proto_msgTypes,
	}.Build()
	File_engine_api_worker_proto = out.File
	file_engine_api_worker_proto_rawDesc = nil
	file_engine_api_worker_proto_goTypes = nil
	file_engine_api_worker_proto_depIdxs = nil
}
