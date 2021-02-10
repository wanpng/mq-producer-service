// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: protos/domain/candidate.proto

package domain

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

type Candidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName       string            `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName        string            `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	PhotoUrl        string            `protobuf:"bytes,4,opt,name=photo_url,json=photoUrl,proto3" json:"photo_url,omitempty"`
	IsLandOwner     bool              `protobuf:"varint,5,opt,name=is_land_owner,json=isLandOwner,proto3" json:"is_land_owner,omitempty"`
	CareerLevelId   int32             `protobuf:"varint,6,opt,name=career_level_id,json=careerLevelId,proto3" json:"career_level_id,omitempty"`
	CareerLevel     string            `protobuf:"bytes,7,opt,name=career_level,json=careerLevel,proto3" json:"career_level,omitempty"`
	Summary         string            `protobuf:"bytes,8,opt,name=summary,proto3" json:"summary,omitempty"`
	ProvinceId      int32             `protobuf:"varint,9,opt,name=province_id,json=provinceId,proto3" json:"province_id,omitempty"`
	Province        string            `protobuf:"bytes,10,opt,name=province,proto3" json:"province,omitempty"`
	EmploymentTypes []*EmploymentType `protobuf:"bytes,11,rep,name=employment_types,json=employmentTypes,proto3" json:"employment_types,omitempty"`
	AvailabilityId  int32             `protobuf:"varint,12,opt,name=availability_id,json=availabilityId,proto3" json:"availability_id,omitempty"`
	Availability    string            `protobuf:"bytes,13,opt,name=availability,proto3" json:"availability,omitempty"`
	Skills          []*Skill          `protobuf:"bytes,14,rep,name=skills,proto3" json:"skills,omitempty"`
	Educations      []*Education      `protobuf:"bytes,15,rep,name=educations,proto3" json:"educations,omitempty"`
	WorkExperiences []*WorkExperience `protobuf:"bytes,16,rep,name=work_experiences,json=workExperiences,proto3" json:"work_experiences,omitempty"`
	Trainings       []*Training       `protobuf:"bytes,17,rep,name=trainings,proto3" json:"trainings,omitempty"`
}

func (x *Candidate) Reset() {
	*x = Candidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_domain_candidate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candidate) ProtoMessage() {}

func (x *Candidate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_domain_candidate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candidate.ProtoReflect.Descriptor instead.
func (*Candidate) Descriptor() ([]byte, []int) {
	return file_protos_domain_candidate_proto_rawDescGZIP(), []int{0}
}

func (x *Candidate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Candidate) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Candidate) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Candidate) GetPhotoUrl() string {
	if x != nil {
		return x.PhotoUrl
	}
	return ""
}

func (x *Candidate) GetIsLandOwner() bool {
	if x != nil {
		return x.IsLandOwner
	}
	return false
}

func (x *Candidate) GetCareerLevelId() int32 {
	if x != nil {
		return x.CareerLevelId
	}
	return 0
}

func (x *Candidate) GetCareerLevel() string {
	if x != nil {
		return x.CareerLevel
	}
	return ""
}

func (x *Candidate) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Candidate) GetProvinceId() int32 {
	if x != nil {
		return x.ProvinceId
	}
	return 0
}

func (x *Candidate) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Candidate) GetEmploymentTypes() []*EmploymentType {
	if x != nil {
		return x.EmploymentTypes
	}
	return nil
}

func (x *Candidate) GetAvailabilityId() int32 {
	if x != nil {
		return x.AvailabilityId
	}
	return 0
}

func (x *Candidate) GetAvailability() string {
	if x != nil {
		return x.Availability
	}
	return ""
}

func (x *Candidate) GetSkills() []*Skill {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *Candidate) GetEducations() []*Education {
	if x != nil {
		return x.Educations
	}
	return nil
}

func (x *Candidate) GetWorkExperiences() []*WorkExperience {
	if x != nil {
		return x.WorkExperiences
	}
	return nil
}

func (x *Candidate) GetTrainings() []*Training {
	if x != nil {
		return x.Trainings
	}
	return nil
}

var File_protos_domain_candidate_proto protoreflect.FileDescriptor

var file_protos_domain_candidate_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x23,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x64,
	0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xba, 0x05, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x6c, 0x61,
	0x6e, 0x64, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x4c, 0x61, 0x6e, 0x64, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x63,
	0x61, 0x72, 0x65, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x65, 0x65,
	0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x48, 0x0a,
	0x10, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x0e,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x12, 0x38, 0x0a, 0x0a, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x48, 0x0a, 0x10,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x52, 0x09, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x6e, 0x70,
	0x6e, 0x67, 0x2f, 0x6d, 0x71, 0x2d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_domain_candidate_proto_rawDescOnce sync.Once
	file_protos_domain_candidate_proto_rawDescData = file_protos_domain_candidate_proto_rawDesc
)

func file_protos_domain_candidate_proto_rawDescGZIP() []byte {
	file_protos_domain_candidate_proto_rawDescOnce.Do(func() {
		file_protos_domain_candidate_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_domain_candidate_proto_rawDescData)
	})
	return file_protos_domain_candidate_proto_rawDescData
}

var file_protos_domain_candidate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_domain_candidate_proto_goTypes = []interface{}{
	(*Candidate)(nil),      // 0: protos.domain.Candidate
	(*EmploymentType)(nil), // 1: protos.domain.EmploymentType
	(*Skill)(nil),          // 2: protos.domain.Skill
	(*Education)(nil),      // 3: protos.domain.Education
	(*WorkExperience)(nil), // 4: protos.domain.WorkExperience
	(*Training)(nil),       // 5: protos.domain.Training
}
var file_protos_domain_candidate_proto_depIdxs = []int32{
	1, // 0: protos.domain.Candidate.employment_types:type_name -> protos.domain.EmploymentType
	2, // 1: protos.domain.Candidate.skills:type_name -> protos.domain.Skill
	3, // 2: protos.domain.Candidate.educations:type_name -> protos.domain.Education
	4, // 3: protos.domain.Candidate.work_experiences:type_name -> protos.domain.WorkExperience
	5, // 4: protos.domain.Candidate.trainings:type_name -> protos.domain.Training
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_domain_candidate_proto_init() }
func file_protos_domain_candidate_proto_init() {
	if File_protos_domain_candidate_proto != nil {
		return
	}
	file_protos_domain_employment_type_proto_init()
	file_protos_domain_skill_proto_init()
	file_protos_domain_education_proto_init()
	file_protos_domain_work_experience_proto_init()
	file_protos_domain_training_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_domain_candidate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candidate); i {
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
			RawDescriptor: file_protos_domain_candidate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_domain_candidate_proto_goTypes,
		DependencyIndexes: file_protos_domain_candidate_proto_depIdxs,
		MessageInfos:      file_protos_domain_candidate_proto_msgTypes,
	}.Build()
	File_protos_domain_candidate_proto = out.File
	file_protos_domain_candidate_proto_rawDesc = nil
	file_protos_domain_candidate_proto_goTypes = nil
	file_protos_domain_candidate_proto_depIdxs = nil
}
