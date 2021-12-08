// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: protos/domain/affiliation.proto

package domain

import (
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

type Affiliation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AffiliationId string `protobuf:"bytes,1,opt,name=affiliation_id,json=affiliationId,proto3" json:"affiliation_id,omitempty"`
	Organization  string `protobuf:"bytes,2,opt,name=organization,proto3" json:"organization,omitempty"`
	Role          string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	Location      string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Description   string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Affiliation) Reset() {
	*x = Affiliation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_domain_affiliation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Affiliation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Affiliation) ProtoMessage() {}

func (x *Affiliation) ProtoReflect() protoreflect.Message {
	mi := &file_protos_domain_affiliation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Affiliation.ProtoReflect.Descriptor instead.
func (*Affiliation) Descriptor() ([]byte, []int) {
	return file_protos_domain_affiliation_proto_rawDescGZIP(), []int{0}
}

func (x *Affiliation) GetAffiliationId() string {
	if x != nil {
		return x.AffiliationId
	}
	return ""
}

func (x *Affiliation) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *Affiliation) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Affiliation) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Affiliation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type JobseekerAffiliation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Affiliations []*Affiliation `protobuf:"bytes,2,rep,name=affiliations,proto3" json:"affiliations,omitempty"`
}

func (x *JobseekerAffiliation) Reset() {
	*x = JobseekerAffiliation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_domain_affiliation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobseekerAffiliation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobseekerAffiliation) ProtoMessage() {}

func (x *JobseekerAffiliation) ProtoReflect() protoreflect.Message {
	mi := &file_protos_domain_affiliation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobseekerAffiliation.ProtoReflect.Descriptor instead.
func (*JobseekerAffiliation) Descriptor() ([]byte, []int) {
	return file_protos_domain_affiliation_proto_rawDescGZIP(), []int{1}
}

func (x *JobseekerAffiliation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobseekerAffiliation) GetAffiliations() []*Affiliation {
	if x != nil {
		return x.Affiliations
	}
	return nil
}

var File_protos_domain_affiliation_proto protoreflect.FileDescriptor

var file_protos_domain_affiliation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x61, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x22, 0xaa, 0x01, 0x0a, 0x0b, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x61, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x66, 0x66, 0x69, 0x6c, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x66, 0x0a,
	0x14, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3e, 0x0a, 0x0c, 0x61, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x66, 0x66, 0x69,
	0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x6e, 0x70, 0x6e, 0x67, 0x2f, 0x6d, 0x71, 0x2d, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_domain_affiliation_proto_rawDescOnce sync.Once
	file_protos_domain_affiliation_proto_rawDescData = file_protos_domain_affiliation_proto_rawDesc
)

func file_protos_domain_affiliation_proto_rawDescGZIP() []byte {
	file_protos_domain_affiliation_proto_rawDescOnce.Do(func() {
		file_protos_domain_affiliation_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_domain_affiliation_proto_rawDescData)
	})
	return file_protos_domain_affiliation_proto_rawDescData
}

var file_protos_domain_affiliation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_domain_affiliation_proto_goTypes = []interface{}{
	(*Affiliation)(nil),          // 0: protos.domain.Affiliation
	(*JobseekerAffiliation)(nil), // 1: protos.domain.JobseekerAffiliation
}
var file_protos_domain_affiliation_proto_depIdxs = []int32{
	0, // 0: protos.domain.JobseekerAffiliation.affiliations:type_name -> protos.domain.Affiliation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_domain_affiliation_proto_init() }
func file_protos_domain_affiliation_proto_init() {
	if File_protos_domain_affiliation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_domain_affiliation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Affiliation); i {
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
		file_protos_domain_affiliation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobseekerAffiliation); i {
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
			RawDescriptor: file_protos_domain_affiliation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_domain_affiliation_proto_goTypes,
		DependencyIndexes: file_protos_domain_affiliation_proto_depIdxs,
		MessageInfos:      file_protos_domain_affiliation_proto_msgTypes,
	}.Build()
	File_protos_domain_affiliation_proto = out.File
	file_protos_domain_affiliation_proto_rawDesc = nil
	file_protos_domain_affiliation_proto_goTypes = nil
	file_protos_domain_affiliation_proto_depIdxs = nil
}
