// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: protos/service/candidate_service.proto

package service

import (
	domain "github.com/wanpng/mq-producer-service/grpc/domain"
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

var File_protos_service_candidate_service_proto protoreflect.FileDescriptor

var file_protos_service_candidate_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6a, 0x6f,
	0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b,
	0x65, 0x72, 0x6a, 0x6f, 0x62, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8a, 0x06, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x14, 0x73, 0x61,
	0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x57, 0x0a, 0x19, 0x73, 0x61, 0x76,
	0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x5b, 0x0a, 0x1b, 0x73, 0x61, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65,
	0x6b, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x50, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x4a, 0x0a, 0x13, 0x73, 0x61, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x4d, 0x0a, 0x14, 0x73,
	0x61, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x51, 0x0a, 0x16, 0x73, 0x61,
	0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x45, 0x64, 0x75, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x45, 0x64,
	0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5b, 0x0a,
	0x1b, 0x73, 0x61, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x57, 0x6f,
	0x72, 0x6b, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4a, 0x6f, 0x62,
	0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x4f, 0x0a, 0x15, 0x73, 0x61,
	0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x54, 0x72, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x55, 0x0a, 0x18, 0x73,
	0x61, 0x76, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x41, 0x66, 0x66, 0x69,
	0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x73, 0x65, 0x65, 0x6b, 0x65,
	0x72, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x61, 0x6e, 0x70, 0x6e, 0x67, 0x2f, 0x6d, 0x71, 0x2d, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protos_service_candidate_service_proto_goTypes = []interface{}{
	(*domain.JobseekerProfile)(nil),        // 0: protos.domain.JobseekerProfile
	(*domain.JobseekerProfilePhoto)(nil),   // 1: protos.domain.JobseekerProfilePhoto
	(*domain.JobseekerJobPreferences)(nil), // 2: protos.domain.JobseekerJobPreferences
	(*domain.JobseekerSkill)(nil),          // 3: protos.domain.JobseekerSkill
	(*domain.JobseekerSummary)(nil),        // 4: protos.domain.JobseekerSummary
	(*domain.JobseekerEducation)(nil),      // 5: protos.domain.JobseekerEducation
	(*domain.JobseekerWorkExperience)(nil), // 6: protos.domain.JobseekerWorkExperience
	(*domain.JobseekerTraining)(nil),       // 7: protos.domain.JobseekerTraining
	(*domain.JobseekerAffiliation)(nil),    // 8: protos.domain.JobseekerAffiliation
	(*domain.Error)(nil),                   // 9: protos.domain.Error
}
var file_protos_service_candidate_service_proto_depIdxs = []int32{
	0, // 0: protos.service.CandidateService.saveJobseekerProfile:input_type -> protos.domain.JobseekerProfile
	1, // 1: protos.service.CandidateService.saveJobseekerProfilePhoto:input_type -> protos.domain.JobseekerProfilePhoto
	2, // 2: protos.service.CandidateService.saveJobseekerJobPreferences:input_type -> protos.domain.JobseekerJobPreferences
	3, // 3: protos.service.CandidateService.saveJobseekerSkills:input_type -> protos.domain.JobseekerSkill
	4, // 4: protos.service.CandidateService.saveJobseekerSummary:input_type -> protos.domain.JobseekerSummary
	5, // 5: protos.service.CandidateService.saveJobseekerEducation:input_type -> protos.domain.JobseekerEducation
	6, // 6: protos.service.CandidateService.saveJobseekerWorkExperience:input_type -> protos.domain.JobseekerWorkExperience
	7, // 7: protos.service.CandidateService.saveJobseekerTraining:input_type -> protos.domain.JobseekerTraining
	8, // 8: protos.service.CandidateService.saveJobseekerAffiliation:input_type -> protos.domain.JobseekerAffiliation
	9, // 9: protos.service.CandidateService.saveJobseekerProfile:output_type -> protos.domain.Error
	9, // 10: protos.service.CandidateService.saveJobseekerProfilePhoto:output_type -> protos.domain.Error
	9, // 11: protos.service.CandidateService.saveJobseekerJobPreferences:output_type -> protos.domain.Error
	9, // 12: protos.service.CandidateService.saveJobseekerSkills:output_type -> protos.domain.Error
	9, // 13: protos.service.CandidateService.saveJobseekerSummary:output_type -> protos.domain.Error
	9, // 14: protos.service.CandidateService.saveJobseekerEducation:output_type -> protos.domain.Error
	9, // 15: protos.service.CandidateService.saveJobseekerWorkExperience:output_type -> protos.domain.Error
	9, // 16: protos.service.CandidateService.saveJobseekerTraining:output_type -> protos.domain.Error
	9, // 17: protos.service.CandidateService.saveJobseekerAffiliation:output_type -> protos.domain.Error
	9, // [9:18] is the sub-list for method output_type
	0, // [0:9] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_service_candidate_service_proto_init() }
func file_protos_service_candidate_service_proto_init() {
	if File_protos_service_candidate_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_service_candidate_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_service_candidate_service_proto_goTypes,
		DependencyIndexes: file_protos_service_candidate_service_proto_depIdxs,
	}.Build()
	File_protos_service_candidate_service_proto = out.File
	file_protos_service_candidate_service_proto_rawDesc = nil
	file_protos_service_candidate_service_proto_goTypes = nil
	file_protos_service_candidate_service_proto_depIdxs = nil
}
