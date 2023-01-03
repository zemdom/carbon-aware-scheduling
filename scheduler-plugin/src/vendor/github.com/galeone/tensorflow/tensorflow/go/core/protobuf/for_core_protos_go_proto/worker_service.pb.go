// Copyright 2016 The TensorFlow Authors. All Rights Reserved.
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
//==============================================================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: tensorflow/core/protobuf/worker_service.proto

package for_core_protos_go_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_tensorflow_core_protobuf_worker_service_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_worker_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf0, 0x09, 0x0a, 0x0d, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x12, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x44, 0x65, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x22, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x65,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x52, 0x75, 0x6e, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x12, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52,
	0x75, 0x6e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x75, 0x6e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c,
	0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x1f, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75,
	0x70, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e,
	0x75, 0x70, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4b, 0x0a, 0x0a, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x41, 0x6c, 0x6c, 0x12, 0x1d, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e,
	0x75, 0x70, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x75,
	0x70, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a,
	0x52, 0x65, 0x63, 0x76, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x1d, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x63, 0x76, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x63, 0x76, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x07, 0x4c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x4c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x42, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x76, 0x42, 0x75, 0x66, 0x12, 0x1a,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x63, 0x76,
	0x42, 0x75, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x63, 0x76, 0x42, 0x75, 0x66, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x65, 0x70, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x22, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x65,
	0x70, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x65, 0x70, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x10, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x23, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x87, 0x01, 0x0a, 0x1a, 0x6f,
	0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01,
	0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6c,
	0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_tensorflow_core_protobuf_worker_service_proto_goTypes = []interface{}{
	(*GetStatusRequest)(nil),            // 0: tensorflow.GetStatusRequest
	(*CreateWorkerSessionRequest)(nil),  // 1: tensorflow.CreateWorkerSessionRequest
	(*DeleteWorkerSessionRequest)(nil),  // 2: tensorflow.DeleteWorkerSessionRequest
	(*RegisterGraphRequest)(nil),        // 3: tensorflow.RegisterGraphRequest
	(*DeregisterGraphRequest)(nil),      // 4: tensorflow.DeregisterGraphRequest
	(*RunGraphRequest)(nil),             // 5: tensorflow.RunGraphRequest
	(*CleanupGraphRequest)(nil),         // 6: tensorflow.CleanupGraphRequest
	(*CleanupAllRequest)(nil),           // 7: tensorflow.CleanupAllRequest
	(*RecvTensorRequest)(nil),           // 8: tensorflow.RecvTensorRequest
	(*LoggingRequest)(nil),              // 9: tensorflow.LoggingRequest
	(*TracingRequest)(nil),              // 10: tensorflow.TracingRequest
	(*RecvBufRequest)(nil),              // 11: tensorflow.RecvBufRequest
	(*GetStepSequenceRequest)(nil),      // 12: tensorflow.GetStepSequenceRequest
	(*CompleteGroupRequest)(nil),        // 13: tensorflow.CompleteGroupRequest
	(*CompleteInstanceRequest)(nil),     // 14: tensorflow.CompleteInstanceRequest
	(*GetStatusResponse)(nil),           // 15: tensorflow.GetStatusResponse
	(*CreateWorkerSessionResponse)(nil), // 16: tensorflow.CreateWorkerSessionResponse
	(*DeleteWorkerSessionResponse)(nil), // 17: tensorflow.DeleteWorkerSessionResponse
	(*RegisterGraphResponse)(nil),       // 18: tensorflow.RegisterGraphResponse
	(*DeregisterGraphResponse)(nil),     // 19: tensorflow.DeregisterGraphResponse
	(*RunGraphResponse)(nil),            // 20: tensorflow.RunGraphResponse
	(*CleanupGraphResponse)(nil),        // 21: tensorflow.CleanupGraphResponse
	(*CleanupAllResponse)(nil),          // 22: tensorflow.CleanupAllResponse
	(*RecvTensorResponse)(nil),          // 23: tensorflow.RecvTensorResponse
	(*LoggingResponse)(nil),             // 24: tensorflow.LoggingResponse
	(*TracingResponse)(nil),             // 25: tensorflow.TracingResponse
	(*RecvBufResponse)(nil),             // 26: tensorflow.RecvBufResponse
	(*GetStepSequenceResponse)(nil),     // 27: tensorflow.GetStepSequenceResponse
	(*CompleteGroupResponse)(nil),       // 28: tensorflow.CompleteGroupResponse
	(*CompleteInstanceResponse)(nil),    // 29: tensorflow.CompleteInstanceResponse
}
var file_tensorflow_core_protobuf_worker_service_proto_depIdxs = []int32{
	0,  // 0: tensorflow.grpc.WorkerService.GetStatus:input_type -> tensorflow.GetStatusRequest
	1,  // 1: tensorflow.grpc.WorkerService.CreateWorkerSession:input_type -> tensorflow.CreateWorkerSessionRequest
	2,  // 2: tensorflow.grpc.WorkerService.DeleteWorkerSession:input_type -> tensorflow.DeleteWorkerSessionRequest
	3,  // 3: tensorflow.grpc.WorkerService.RegisterGraph:input_type -> tensorflow.RegisterGraphRequest
	4,  // 4: tensorflow.grpc.WorkerService.DeregisterGraph:input_type -> tensorflow.DeregisterGraphRequest
	5,  // 5: tensorflow.grpc.WorkerService.RunGraph:input_type -> tensorflow.RunGraphRequest
	6,  // 6: tensorflow.grpc.WorkerService.CleanupGraph:input_type -> tensorflow.CleanupGraphRequest
	7,  // 7: tensorflow.grpc.WorkerService.CleanupAll:input_type -> tensorflow.CleanupAllRequest
	8,  // 8: tensorflow.grpc.WorkerService.RecvTensor:input_type -> tensorflow.RecvTensorRequest
	9,  // 9: tensorflow.grpc.WorkerService.Logging:input_type -> tensorflow.LoggingRequest
	10, // 10: tensorflow.grpc.WorkerService.Tracing:input_type -> tensorflow.TracingRequest
	11, // 11: tensorflow.grpc.WorkerService.RecvBuf:input_type -> tensorflow.RecvBufRequest
	12, // 12: tensorflow.grpc.WorkerService.GetStepSequence:input_type -> tensorflow.GetStepSequenceRequest
	13, // 13: tensorflow.grpc.WorkerService.CompleteGroup:input_type -> tensorflow.CompleteGroupRequest
	14, // 14: tensorflow.grpc.WorkerService.CompleteInstance:input_type -> tensorflow.CompleteInstanceRequest
	15, // 15: tensorflow.grpc.WorkerService.GetStatus:output_type -> tensorflow.GetStatusResponse
	16, // 16: tensorflow.grpc.WorkerService.CreateWorkerSession:output_type -> tensorflow.CreateWorkerSessionResponse
	17, // 17: tensorflow.grpc.WorkerService.DeleteWorkerSession:output_type -> tensorflow.DeleteWorkerSessionResponse
	18, // 18: tensorflow.grpc.WorkerService.RegisterGraph:output_type -> tensorflow.RegisterGraphResponse
	19, // 19: tensorflow.grpc.WorkerService.DeregisterGraph:output_type -> tensorflow.DeregisterGraphResponse
	20, // 20: tensorflow.grpc.WorkerService.RunGraph:output_type -> tensorflow.RunGraphResponse
	21, // 21: tensorflow.grpc.WorkerService.CleanupGraph:output_type -> tensorflow.CleanupGraphResponse
	22, // 22: tensorflow.grpc.WorkerService.CleanupAll:output_type -> tensorflow.CleanupAllResponse
	23, // 23: tensorflow.grpc.WorkerService.RecvTensor:output_type -> tensorflow.RecvTensorResponse
	24, // 24: tensorflow.grpc.WorkerService.Logging:output_type -> tensorflow.LoggingResponse
	25, // 25: tensorflow.grpc.WorkerService.Tracing:output_type -> tensorflow.TracingResponse
	26, // 26: tensorflow.grpc.WorkerService.RecvBuf:output_type -> tensorflow.RecvBufResponse
	27, // 27: tensorflow.grpc.WorkerService.GetStepSequence:output_type -> tensorflow.GetStepSequenceResponse
	28, // 28: tensorflow.grpc.WorkerService.CompleteGroup:output_type -> tensorflow.CompleteGroupResponse
	29, // 29: tensorflow.grpc.WorkerService.CompleteInstance:output_type -> tensorflow.CompleteInstanceResponse
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_worker_service_proto_init() }
func file_tensorflow_core_protobuf_worker_service_proto_init() {
	if File_tensorflow_core_protobuf_worker_service_proto != nil {
		return
	}
	file_tensorflow_core_protobuf_worker_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_protobuf_worker_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tensorflow_core_protobuf_worker_service_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_worker_service_proto_depIdxs,
	}.Build()
	File_tensorflow_core_protobuf_worker_service_proto = out.File
	file_tensorflow_core_protobuf_worker_service_proto_rawDesc = nil
	file_tensorflow_core_protobuf_worker_service_proto_goTypes = nil
	file_tensorflow_core_protobuf_worker_service_proto_depIdxs = nil
}
