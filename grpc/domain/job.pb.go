// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: protos/domain/job.proto

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

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EmployerId          string    `protobuf:"bytes,2,opt,name=employer_id,json=employerId,proto3" json:"employer_id,omitempty"`
	Employer            string    `protobuf:"bytes,3,opt,name=employer,proto3" json:"employer,omitempty"`
	JobTitle            string    `protobuf:"bytes,4,opt,name=job_title,json=jobTitle,proto3" json:"job_title,omitempty"`
	ProvinceId          int32     `protobuf:"varint,5,opt,name=province_id,json=provinceId,proto3" json:"province_id,omitempty"`
	Province            string    `protobuf:"bytes,6,opt,name=province,proto3" json:"province,omitempty"`
	CareerLevelId       int32     `protobuf:"varint,7,opt,name=career_level_id,json=careerLevelId,proto3" json:"career_level_id,omitempty"`
	CareerLevel         string    `protobuf:"bytes,8,opt,name=career_level,json=careerLevel,proto3" json:"career_level,omitempty"`
	EmploymentTypeId    int32     `protobuf:"varint,9,opt,name=employment_type_id,json=employmentTypeId,proto3" json:"employment_type_id,omitempty"`
	EmploymentType      string    `protobuf:"bytes,10,opt,name=employment_type,json=employmentType,proto3" json:"employment_type,omitempty"`
	PublishStatusDate   int64     `protobuf:"zigzag64,11,opt,name=publish_status_date,json=publishStatusDate,proto3" json:"publish_status_date,omitempty"`
	JobSummary          string    `protobuf:"bytes,12,opt,name=job_summary,json=jobSummary,proto3" json:"job_summary,omitempty"`
	JobDescription      string    `protobuf:"bytes,13,opt,name=job_description,json=jobDescription,proto3" json:"job_description,omitempty"`
	JobClassificationId int32     `protobuf:"varint,14,opt,name=job_classification_id,json=jobClassificationId,proto3" json:"job_classification_id,omitempty"`
	JobClassification   string    `protobuf:"bytes,15,opt,name=job_classification,json=jobClassification,proto3" json:"job_classification,omitempty"`
	DateCreated         int64     `protobuf:"zigzag64,16,opt,name=date_created,json=dateCreated,proto3" json:"date_created,omitempty"`
	PayPeriodId         int32     `protobuf:"varint,17,opt,name=pay_period_id,json=payPeriodId,proto3" json:"pay_period_id,omitempty"`
	PayPeriod           string    `protobuf:"bytes,18,opt,name=pay_period,json=payPeriod,proto3" json:"pay_period,omitempty"`
	MinimumRate         int32     `protobuf:"varint,19,opt,name=minimum_rate,json=minimumRate,proto3" json:"minimum_rate,omitempty"`
	EmployerCompany     *Employer `protobuf:"bytes,20,opt,name=employer_company,json=employerCompany,proto3" json:"employer_company,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_domain_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_protos_domain_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_protos_domain_job_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Job) GetEmployerId() string {
	if x != nil {
		return x.EmployerId
	}
	return ""
}

func (x *Job) GetEmployer() string {
	if x != nil {
		return x.Employer
	}
	return ""
}

func (x *Job) GetJobTitle() string {
	if x != nil {
		return x.JobTitle
	}
	return ""
}

func (x *Job) GetProvinceId() int32 {
	if x != nil {
		return x.ProvinceId
	}
	return 0
}

func (x *Job) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Job) GetCareerLevelId() int32 {
	if x != nil {
		return x.CareerLevelId
	}
	return 0
}

func (x *Job) GetCareerLevel() string {
	if x != nil {
		return x.CareerLevel
	}
	return ""
}

func (x *Job) GetEmploymentTypeId() int32 {
	if x != nil {
		return x.EmploymentTypeId
	}
	return 0
}

func (x *Job) GetEmploymentType() string {
	if x != nil {
		return x.EmploymentType
	}
	return ""
}

func (x *Job) GetPublishStatusDate() int64 {
	if x != nil {
		return x.PublishStatusDate
	}
	return 0
}

func (x *Job) GetJobSummary() string {
	if x != nil {
		return x.JobSummary
	}
	return ""
}

func (x *Job) GetJobDescription() string {
	if x != nil {
		return x.JobDescription
	}
	return ""
}

func (x *Job) GetJobClassificationId() int32 {
	if x != nil {
		return x.JobClassificationId
	}
	return 0
}

func (x *Job) GetJobClassification() string {
	if x != nil {
		return x.JobClassification
	}
	return ""
}

func (x *Job) GetDateCreated() int64 {
	if x != nil {
		return x.DateCreated
	}
	return 0
}

func (x *Job) GetPayPeriodId() int32 {
	if x != nil {
		return x.PayPeriodId
	}
	return 0
}

func (x *Job) GetPayPeriod() string {
	if x != nil {
		return x.PayPeriod
	}
	return ""
}

func (x *Job) GetMinimumRate() int32 {
	if x != nil {
		return x.MinimumRate
	}
	return 0
}

func (x *Job) GetEmployerCompany() *Employer {
	if x != nil {
		return x.EmployerCompany
	}
	return nil
}

type Employer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployerId       string `protobuf:"bytes,1,opt,name=employer_id,json=employerId,proto3" json:"employer_id,omitempty"`
	EmployerName     string `protobuf:"bytes,2,opt,name=employer_name,json=employerName,proto3" json:"employer_name,omitempty"`
	EmployerPhotoUrl string `protobuf:"bytes,3,opt,name=employer_photo_url,json=employerPhotoUrl,proto3" json:"employer_photo_url,omitempty"`
}

func (x *Employer) Reset() {
	*x = Employer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_domain_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employer) ProtoMessage() {}

func (x *Employer) ProtoReflect() protoreflect.Message {
	mi := &file_protos_domain_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employer.ProtoReflect.Descriptor instead.
func (*Employer) Descriptor() ([]byte, []int) {
	return file_protos_domain_job_proto_rawDescGZIP(), []int{1}
}

func (x *Employer) GetEmployerId() string {
	if x != nil {
		return x.EmployerId
	}
	return ""
}

func (x *Employer) GetEmployerName() string {
	if x != nil {
		return x.EmployerName
	}
	return ""
}

func (x *Employer) GetEmployerPhotoUrl() string {
	if x != nil {
		return x.EmployerPhotoUrl
	}
	return ""
}

type DeleteJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteJob) Reset() {
	*x = DeleteJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_domain_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJob) ProtoMessage() {}

func (x *DeleteJob) ProtoReflect() protoreflect.Message {
	mi := &file_protos_domain_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJob.ProtoReflect.Descriptor instead.
func (*DeleteJob) Descriptor() ([]byte, []int) {
	return file_protos_domain_job_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteJob) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type JobCompanyProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EmployerCompany *Employer `protobuf:"bytes,2,opt,name=employer_company,json=employerCompany,proto3" json:"employer_company,omitempty"`
}

func (x *JobCompanyProfile) Reset() {
	*x = JobCompanyProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_domain_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobCompanyProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobCompanyProfile) ProtoMessage() {}

func (x *JobCompanyProfile) ProtoReflect() protoreflect.Message {
	mi := &file_protos_domain_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobCompanyProfile.ProtoReflect.Descriptor instead.
func (*JobCompanyProfile) Descriptor() ([]byte, []int) {
	return file_protos_domain_job_proto_rawDescGZIP(), []int{3}
}

func (x *JobCompanyProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobCompanyProfile) GetEmployerCompany() *Employer {
	if x != nil {
		return x.EmployerCompany
	}
	return nil
}

var File_protos_domain_job_proto protoreflect.FileDescriptor

var file_protos_domain_job_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xf8, 0x05, 0x0a, 0x03, 0x4a, 0x6f, 0x62,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6a, 0x6f, 0x62, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x61, 0x72, 0x65, 0x65,
	0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x12, 0x52, 0x11, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6a, 0x6f, 0x62,
	0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6a, 0x6f, 0x62, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x6a, 0x6f,
	0x62, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x6a, 0x6f, 0x62, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x15, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x13, 0x6a, 0x6f, 0x62, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x6a, 0x6f, 0x62, 0x5f, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x6a, 0x6f, 0x62, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0b, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x61, 0x79,
	0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x70, 0x61, 0x79, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x42, 0x0a, 0x10, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x72, 0x52, 0x0f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x22, 0x7e, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x72, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x55, 0x72, 0x6c, 0x22, 0x1b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x67, 0x0a, 0x11, 0x4a, 0x6f, 0x62, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x10, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x52, 0x0f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x6e, 0x70, 0x6e, 0x67, 0x2f, 0x6d,
	0x71, 0x2d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_domain_job_proto_rawDescOnce sync.Once
	file_protos_domain_job_proto_rawDescData = file_protos_domain_job_proto_rawDesc
)

func file_protos_domain_job_proto_rawDescGZIP() []byte {
	file_protos_domain_job_proto_rawDescOnce.Do(func() {
		file_protos_domain_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_domain_job_proto_rawDescData)
	})
	return file_protos_domain_job_proto_rawDescData
}

var file_protos_domain_job_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_domain_job_proto_goTypes = []interface{}{
	(*Job)(nil),               // 0: protos.domain.Job
	(*Employer)(nil),          // 1: protos.domain.Employer
	(*DeleteJob)(nil),         // 2: protos.domain.DeleteJob
	(*JobCompanyProfile)(nil), // 3: protos.domain.JobCompanyProfile
}
var file_protos_domain_job_proto_depIdxs = []int32{
	1, // 0: protos.domain.Job.employer_company:type_name -> protos.domain.Employer
	1, // 1: protos.domain.JobCompanyProfile.employer_company:type_name -> protos.domain.Employer
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_domain_job_proto_init() }
func file_protos_domain_job_proto_init() {
	if File_protos_domain_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_domain_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_protos_domain_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employer); i {
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
		file_protos_domain_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJob); i {
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
		file_protos_domain_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobCompanyProfile); i {
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
			RawDescriptor: file_protos_domain_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_domain_job_proto_goTypes,
		DependencyIndexes: file_protos_domain_job_proto_depIdxs,
		MessageInfos:      file_protos_domain_job_proto_msgTypes,
	}.Build()
	File_protos_domain_job_proto = out.File
	file_protos_domain_job_proto_rawDesc = nil
	file_protos_domain_job_proto_goTypes = nil
	file_protos_domain_job_proto_depIdxs = nil
}
