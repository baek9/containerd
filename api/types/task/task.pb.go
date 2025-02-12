//
//Copyright The containerd Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.4
// source: github.com/containerd/containerd/api/types/task/task.proto

package task

import (
	_ "github.com/gogo/protobuf/gogoproto"
	any1 "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Status int32

const (
	Status_UNKNOWN Status = 0
	Status_CREATED Status = 1
	Status_RUNNING Status = 2
	Status_STOPPED Status = 3
	Status_PAUSED  Status = 4
	Status_PAUSING Status = 5
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UNKNOWN",
		1: "CREATED",
		2: "RUNNING",
		3: "STOPPED",
		4: "PAUSED",
		5: "PAUSING",
	}
	Status_value = map[string]int32{
		"UNKNOWN": 0,
		"CREATED": 1,
		"RUNNING": 2,
		"STOPPED": 3,
		"PAUSED":  4,
		"PAUSING": 5,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_containerd_containerd_api_types_task_task_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_github_com_containerd_containerd_api_types_task_task_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_types_task_task_proto_rawDescGZIP(), []int{0}
}

type Process struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string               `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ID          string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Pid         uint32               `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Status      Status               `protobuf:"varint,4,opt,name=status,proto3,enum=containerd.v1.types.Status" json:"status,omitempty"`
	Stdin       string               `protobuf:"bytes,5,opt,name=stdin,proto3" json:"stdin,omitempty"`
	Stdout      string               `protobuf:"bytes,6,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr      string               `protobuf:"bytes,7,opt,name=stderr,proto3" json:"stderr,omitempty"`
	Terminal    bool                 `protobuf:"varint,8,opt,name=terminal,proto3" json:"terminal,omitempty"`
	ExitStatus  uint32               `protobuf:"varint,9,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	ExitedAt    *timestamp.Timestamp `protobuf:"bytes,10,opt,name=exited_at,json=exitedAt,proto3" json:"exited_at,omitempty"`
}

func (x *Process) Reset() {
	*x = Process{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_types_task_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Process) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process) ProtoMessage() {}

func (x *Process) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_types_task_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process.ProtoReflect.Descriptor instead.
func (*Process) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_types_task_task_proto_rawDescGZIP(), []int{0}
}

func (x *Process) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *Process) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Process) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Process) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNKNOWN
}

func (x *Process) GetStdin() string {
	if x != nil {
		return x.Stdin
	}
	return ""
}

func (x *Process) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *Process) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *Process) GetTerminal() bool {
	if x != nil {
		return x.Terminal
	}
	return false
}

func (x *Process) GetExitStatus() uint32 {
	if x != nil {
		return x.ExitStatus
	}
	return 0
}

func (x *Process) GetExitedAt() *timestamp.Timestamp {
	if x != nil {
		return x.ExitedAt
	}
	return nil
}

type ProcessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PID is the process ID.
	Pid uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	// Info contains additional process information.
	//
	// Info varies by platform.
	Info *any1.Any `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ProcessInfo) Reset() {
	*x = ProcessInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_types_task_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessInfo) ProtoMessage() {}

func (x *ProcessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_types_task_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessInfo.ProtoReflect.Descriptor instead.
func (*ProcessInfo) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_types_task_task_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessInfo) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ProcessInfo) GetInfo() *any1.Any {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_github_com_containerd_containerd_api_types_task_task_proto protoreflect.FileDescriptor

var file_github_com_containerd_containerd_api_types_task_task_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x70, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x64,
	0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x64, 0x69, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x78, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x65, 0x78, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x09,
	0x65, 0x78, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65, 0x78, 0x69,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x49, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x2a, 0x55, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x41,
	0x55, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_github_com_containerd_containerd_api_types_task_task_proto_rawDescOnce sync.Once
	file_github_com_containerd_containerd_api_types_task_task_proto_rawDescData = file_github_com_containerd_containerd_api_types_task_task_proto_rawDesc
)

func file_github_com_containerd_containerd_api_types_task_task_proto_rawDescGZIP() []byte {
	file_github_com_containerd_containerd_api_types_task_task_proto_rawDescOnce.Do(func() {
		file_github_com_containerd_containerd_api_types_task_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_containerd_containerd_api_types_task_task_proto_rawDescData)
	})
	return file_github_com_containerd_containerd_api_types_task_task_proto_rawDescData
}

var file_github_com_containerd_containerd_api_types_task_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_containerd_containerd_api_types_task_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_containerd_containerd_api_types_task_task_proto_goTypes = []interface{}{
	(Status)(0),                 // 0: containerd.v1.types.Status
	(*Process)(nil),             // 1: containerd.v1.types.Process
	(*ProcessInfo)(nil),         // 2: containerd.v1.types.ProcessInfo
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*any1.Any)(nil),            // 4: google.protobuf.Any
}
var file_github_com_containerd_containerd_api_types_task_task_proto_depIdxs = []int32{
	0, // 0: containerd.v1.types.Process.status:type_name -> containerd.v1.types.Status
	3, // 1: containerd.v1.types.Process.exited_at:type_name -> google.protobuf.Timestamp
	4, // 2: containerd.v1.types.ProcessInfo.info:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_containerd_containerd_api_types_task_task_proto_init() }
func file_github_com_containerd_containerd_api_types_task_task_proto_init() {
	if File_github_com_containerd_containerd_api_types_task_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_containerd_containerd_api_types_task_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Process); i {
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
		file_github_com_containerd_containerd_api_types_task_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessInfo); i {
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
			RawDescriptor: file_github_com_containerd_containerd_api_types_task_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_containerd_containerd_api_types_task_task_proto_goTypes,
		DependencyIndexes: file_github_com_containerd_containerd_api_types_task_task_proto_depIdxs,
		EnumInfos:         file_github_com_containerd_containerd_api_types_task_task_proto_enumTypes,
		MessageInfos:      file_github_com_containerd_containerd_api_types_task_task_proto_msgTypes,
	}.Build()
	File_github_com_containerd_containerd_api_types_task_task_proto = out.File
	file_github_com_containerd_containerd_api_types_task_task_proto_rawDesc = nil
	file_github_com_containerd_containerd_api_types_task_task_proto_goTypes = nil
	file_github_com_containerd_containerd_api_types_task_task_proto_depIdxs = nil
}
