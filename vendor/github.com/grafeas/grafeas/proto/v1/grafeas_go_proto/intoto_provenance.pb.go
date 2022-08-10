// Copyright 2021 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: proto/v1/intoto_provenance.proto

package grafeas_go_proto

import (
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

// Steps taken to build the artifact.
// For a TaskRun, typically each container corresponds to one step in the recipe.
type Recipe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URI indicating what type of recipe was performed. It determines the meaning of recipe.entryPoint, recipe.arguments, recipe.environment, and materials.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Index in materials containing the recipe steps that are not implied by recipe.type.
	// For example, if the recipe type were "make", then this would point to the source containing the Makefile, not the make program itself.
	// Set to -1 if the recipe doesn't come from a material, as zero is default unset value for int64.
	DefinedInMaterial int64 `protobuf:"varint,2,opt,name=defined_in_material,json=definedInMaterial,proto3" json:"defined_in_material,omitempty"`
	// String identifying the entry point into the build.
	// This is often a path to a configuration file and/or a target label within that file.
	// The syntax and meaning are defined by recipe.type.
	// For example, if the recipe type were "make", then this would reference the directory in which to run make as well as which target to use.
	EntryPoint string `protobuf:"bytes,3,opt,name=entry_point,json=entryPoint,proto3" json:"entry_point,omitempty"`
	// Collection of all external inputs that influenced the build on top of recipe.definedInMaterial and recipe.entryPoint.
	// For example, if the recipe type were "make", then this might be the flags passed to make aside from the target, which is captured in recipe.entryPoint.
	// Since the arguments field can greatly vary in structure, depending on the builder and recipe type, this is of form "Any".
	Arguments []*any1.Any `protobuf:"bytes,4,rep,name=arguments,proto3" json:"arguments,omitempty"`
	// Any other builder-controlled inputs necessary for correctly evaluating the recipe. Usually only needed for reproducing the build but not evaluated as part of policy.
	// Since the environment field can greatly vary in structure, depending on the builder and recipe type, this is of form "Any".
	Environment []*any1.Any `protobuf:"bytes,5,rep,name=environment,proto3" json:"environment,omitempty"`
}

func (x *Recipe) Reset() {
	*x = Recipe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_intoto_provenance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipe) ProtoMessage() {}

func (x *Recipe) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_intoto_provenance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipe.ProtoReflect.Descriptor instead.
func (*Recipe) Descriptor() ([]byte, []int) {
	return file_proto_v1_intoto_provenance_proto_rawDescGZIP(), []int{0}
}

func (x *Recipe) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Recipe) GetDefinedInMaterial() int64 {
	if x != nil {
		return x.DefinedInMaterial
	}
	return 0
}

func (x *Recipe) GetEntryPoint() string {
	if x != nil {
		return x.EntryPoint
	}
	return ""
}

func (x *Recipe) GetArguments() []*any1.Any {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *Recipe) GetEnvironment() []*any1.Any {
	if x != nil {
		return x.Environment
	}
	return nil
}

// Indicates that the builder claims certain fields in this message to be complete.
type Completeness struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If true, the builder claims that recipe.arguments is complete, meaning that all external inputs are properly captured in the recipe.
	Arguments bool `protobuf:"varint,1,opt,name=arguments,proto3" json:"arguments,omitempty"`
	// If true, the builder claims that recipe.environment is claimed to be complete.
	Environment bool `protobuf:"varint,2,opt,name=environment,proto3" json:"environment,omitempty"`
	// If true, the builder claims that materials are complete, usually through some controls to prevent network access. Sometimes called "hermetic".
	Materials bool `protobuf:"varint,3,opt,name=materials,proto3" json:"materials,omitempty"`
}

func (x *Completeness) Reset() {
	*x = Completeness{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_intoto_provenance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Completeness) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Completeness) ProtoMessage() {}

func (x *Completeness) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_intoto_provenance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Completeness.ProtoReflect.Descriptor instead.
func (*Completeness) Descriptor() ([]byte, []int) {
	return file_proto_v1_intoto_provenance_proto_rawDescGZIP(), []int{1}
}

func (x *Completeness) GetArguments() bool {
	if x != nil {
		return x.Arguments
	}
	return false
}

func (x *Completeness) GetEnvironment() bool {
	if x != nil {
		return x.Environment
	}
	return false
}

func (x *Completeness) GetMaterials() bool {
	if x != nil {
		return x.Materials
	}
	return false
}

// Other properties of the build.
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifies the particular build invocation, which can be useful for finding associated logs or other ad-hoc analysis.
	// The value SHOULD be globally unique, per in-toto Provenance spec.
	BuildInvocationId string `protobuf:"bytes,1,opt,name=build_invocation_id,json=buildInvocationId,proto3" json:"build_invocation_id,omitempty"`
	// The timestamp of when the build started.
	BuildStartedOn *timestamp.Timestamp `protobuf:"bytes,2,opt,name=build_started_on,json=buildStartedOn,proto3" json:"build_started_on,omitempty"`
	// The timestamp of when the build completed.
	BuildFinishedOn *timestamp.Timestamp `protobuf:"bytes,3,opt,name=build_finished_on,json=buildFinishedOn,proto3" json:"build_finished_on,omitempty"`
	// Indicates that the builder claims certain fields in this message to be complete.
	Completeness *Completeness `protobuf:"bytes,4,opt,name=completeness,proto3" json:"completeness,omitempty"`
	// If true, the builder claims that running the recipe on materials will produce bit-for-bit identical output.
	Reproducible bool `protobuf:"varint,5,opt,name=reproducible,proto3" json:"reproducible,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_intoto_provenance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_intoto_provenance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_proto_v1_intoto_provenance_proto_rawDescGZIP(), []int{2}
}

func (x *Metadata) GetBuildInvocationId() string {
	if x != nil {
		return x.BuildInvocationId
	}
	return ""
}

func (x *Metadata) GetBuildStartedOn() *timestamp.Timestamp {
	if x != nil {
		return x.BuildStartedOn
	}
	return nil
}

func (x *Metadata) GetBuildFinishedOn() *timestamp.Timestamp {
	if x != nil {
		return x.BuildFinishedOn
	}
	return nil
}

func (x *Metadata) GetCompleteness() *Completeness {
	if x != nil {
		return x.Completeness
	}
	return nil
}

func (x *Metadata) GetReproducible() bool {
	if x != nil {
		return x.Reproducible
	}
	return false
}

type BuilderConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BuilderConfig) Reset() {
	*x = BuilderConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_intoto_provenance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuilderConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuilderConfig) ProtoMessage() {}

func (x *BuilderConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_intoto_provenance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuilderConfig.ProtoReflect.Descriptor instead.
func (*BuilderConfig) Descriptor() ([]byte, []int) {
	return file_proto_v1_intoto_provenance_proto_rawDescGZIP(), []int{3}
}

func (x *BuilderConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type InTotoProvenance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuilderConfig *BuilderConfig `protobuf:"bytes,1,opt,name=builder_config,json=builderConfig,proto3" json:"builder_config,omitempty"` // required
	// Identifies the configuration used for the build.
	// When combined with materials, this SHOULD fully describe the build,
	// such that re-running this recipe results in bit-for-bit identical output
	// (if the build is reproducible).
	Recipe   *Recipe   `protobuf:"bytes,2,opt,name=recipe,proto3" json:"recipe,omitempty"` // required
	Metadata *Metadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The collection of artifacts that influenced the build including sources, dependencies, build tools, base images, and so on.
	// This is considered to be incomplete unless metadata.completeness.materials is true. Unset or null is equivalent to empty.
	Materials []string `protobuf:"bytes,4,rep,name=materials,proto3" json:"materials,omitempty"`
}

func (x *InTotoProvenance) Reset() {
	*x = InTotoProvenance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_intoto_provenance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InTotoProvenance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InTotoProvenance) ProtoMessage() {}

func (x *InTotoProvenance) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_intoto_provenance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InTotoProvenance.ProtoReflect.Descriptor instead.
func (*InTotoProvenance) Descriptor() ([]byte, []int) {
	return file_proto_v1_intoto_provenance_proto_rawDescGZIP(), []int{4}
}

func (x *InTotoProvenance) GetBuilderConfig() *BuilderConfig {
	if x != nil {
		return x.BuilderConfig
	}
	return nil
}

func (x *InTotoProvenance) GetRecipe() *Recipe {
	if x != nil {
		return x.Recipe
	}
	return nil
}

func (x *InTotoProvenance) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *InTotoProvenance) GetMaterials() []string {
	if x != nil {
		return x.Materials
	}
	return nil
}

var File_proto_v1_intoto_provenance_proto protoreflect.FileDescriptor

var file_proto_v1_intoto_provenance_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x74, 0x6f, 0x74,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x06, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x49,
	0x6e, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x09, 0x61, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x36,
	0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x6c, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x22, 0xaa, 0x02, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x2e, 0x0a, 0x13, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x44, 0x0a, 0x10, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x46, 0x0a, 0x11, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x4f, 0x6e, 0x12,
	0x3c, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52,
	0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x69, 0x62, 0x6c,
	0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xd0, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x54, 0x6f, 0x74, 0x6f, 0x50, 0x72, 0x6f,
	0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x72, 0x61, 0x66,
	0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x52, 0x06, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x73, 0x42, 0x64, 0x0a, 0x0d, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x66,
	0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x49, 0x6e, 0x54, 0x6f, 0x74, 0x6f, 0x50, 0x72,
	0x6f, 0x76, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66,
	0x65, 0x61, 0x73, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x5f, 0x67, 0x6f, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_intoto_provenance_proto_rawDescOnce sync.Once
	file_proto_v1_intoto_provenance_proto_rawDescData = file_proto_v1_intoto_provenance_proto_rawDesc
)

func file_proto_v1_intoto_provenance_proto_rawDescGZIP() []byte {
	file_proto_v1_intoto_provenance_proto_rawDescOnce.Do(func() {
		file_proto_v1_intoto_provenance_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_intoto_provenance_proto_rawDescData)
	})
	return file_proto_v1_intoto_provenance_proto_rawDescData
}

var file_proto_v1_intoto_provenance_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_v1_intoto_provenance_proto_goTypes = []interface{}{
	(*Recipe)(nil),              // 0: grafeas.v1.Recipe
	(*Completeness)(nil),        // 1: grafeas.v1.Completeness
	(*Metadata)(nil),            // 2: grafeas.v1.Metadata
	(*BuilderConfig)(nil),       // 3: grafeas.v1.BuilderConfig
	(*InTotoProvenance)(nil),    // 4: grafeas.v1.InTotoProvenance
	(*any1.Any)(nil),            // 5: google.protobuf.Any
	(*timestamp.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_proto_v1_intoto_provenance_proto_depIdxs = []int32{
	5, // 0: grafeas.v1.Recipe.arguments:type_name -> google.protobuf.Any
	5, // 1: grafeas.v1.Recipe.environment:type_name -> google.protobuf.Any
	6, // 2: grafeas.v1.Metadata.build_started_on:type_name -> google.protobuf.Timestamp
	6, // 3: grafeas.v1.Metadata.build_finished_on:type_name -> google.protobuf.Timestamp
	1, // 4: grafeas.v1.Metadata.completeness:type_name -> grafeas.v1.Completeness
	3, // 5: grafeas.v1.InTotoProvenance.builder_config:type_name -> grafeas.v1.BuilderConfig
	0, // 6: grafeas.v1.InTotoProvenance.recipe:type_name -> grafeas.v1.Recipe
	2, // 7: grafeas.v1.InTotoProvenance.metadata:type_name -> grafeas.v1.Metadata
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_proto_v1_intoto_provenance_proto_init() }
func file_proto_v1_intoto_provenance_proto_init() {
	if File_proto_v1_intoto_provenance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_intoto_provenance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recipe); i {
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
		file_proto_v1_intoto_provenance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Completeness); i {
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
		file_proto_v1_intoto_provenance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_proto_v1_intoto_provenance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuilderConfig); i {
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
		file_proto_v1_intoto_provenance_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InTotoProvenance); i {
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
			RawDescriptor: file_proto_v1_intoto_provenance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_v1_intoto_provenance_proto_goTypes,
		DependencyIndexes: file_proto_v1_intoto_provenance_proto_depIdxs,
		MessageInfos:      file_proto_v1_intoto_provenance_proto_msgTypes,
	}.Build()
	File_proto_v1_intoto_provenance_proto = out.File
	file_proto_v1_intoto_provenance_proto_rawDesc = nil
	file_proto_v1_intoto_provenance_proto_goTypes = nil
	file_proto_v1_intoto_provenance_proto_depIdxs = nil
}
