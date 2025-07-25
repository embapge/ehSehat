// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/consultation/consultation.proto

package consultation

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ConsultationRequest represents a request for a consultation
type Prescription struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // name of the medicine
	Dose          string                 `protobuf:"bytes,2,opt,name=dose,proto3" json:"dose,omitempty"` // dosage instructions
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Prescription) Reset() {
	*x = Prescription{}
	mi := &file_proto_consultation_consultation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Prescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prescription) ProtoMessage() {}

func (x *Prescription) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consultation_consultation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prescription.ProtoReflect.Descriptor instead.
func (*Prescription) Descriptor() ([]byte, []int) {
	return file_proto_consultation_consultation_proto_rawDescGZIP(), []int{0}
}

func (x *Prescription) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Prescription) GetDose() string {
	if x != nil {
		return x.Dose
	}
	return ""
}

type ConsultationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                // unique identifier for the consultation
	PatientId     string                 `protobuf:"bytes,2,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`       // from master-service
	PatientName   string                 `protobuf:"bytes,3,opt,name=patient_name,json=patientName,proto3" json:"patient_name,omitempty"` // snapshot, optional
	DoctorId      string                 `protobuf:"bytes,4,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id,omitempty"`          // from master-service
	DoctorName    string                 `protobuf:"bytes,5,opt,name=doctor_name,json=doctorName,proto3" json:"doctor_name,omitempty"`    // snapshot, optional
	RoomId        string                 `protobuf:"bytes,6,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`                // from master-service
	RoomName      string                 `protobuf:"bytes,7,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`          // snapshot, optional
	Symptoms      string                 `protobuf:"bytes,9,opt,name=symptoms,proto3" json:"symptoms,omitempty"`
	Prescription  []*Prescription        `protobuf:"bytes,8,rep,name=prescription,proto3" json:"prescription,omitempty"` // list of prescriptions
	Diagnosis     string                 `protobuf:"bytes,10,opt,name=diagnosis,proto3" json:"diagnosis,omitempty"`
	Date          string                 `protobuf:"bytes,11,opt,name=date,proto3" json:"date,omitempty"` // date of the consultation
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConsultationRequest) Reset() {
	*x = ConsultationRequest{}
	mi := &file_proto_consultation_consultation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConsultationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultationRequest) ProtoMessage() {}

func (x *ConsultationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consultation_consultation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultationRequest.ProtoReflect.Descriptor instead.
func (*ConsultationRequest) Descriptor() ([]byte, []int) {
	return file_proto_consultation_consultation_proto_rawDescGZIP(), []int{1}
}

func (x *ConsultationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ConsultationRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *ConsultationRequest) GetPatientName() string {
	if x != nil {
		return x.PatientName
	}
	return ""
}

func (x *ConsultationRequest) GetDoctorId() string {
	if x != nil {
		return x.DoctorId
	}
	return ""
}

func (x *ConsultationRequest) GetDoctorName() string {
	if x != nil {
		return x.DoctorName
	}
	return ""
}

func (x *ConsultationRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *ConsultationRequest) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *ConsultationRequest) GetSymptoms() string {
	if x != nil {
		return x.Symptoms
	}
	return ""
}

func (x *ConsultationRequest) GetPrescription() []*Prescription {
	if x != nil {
		return x.Prescription
	}
	return nil
}

func (x *ConsultationRequest) GetDiagnosis() string {
	if x != nil {
		return x.Diagnosis
	}
	return ""
}

func (x *ConsultationRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ConsultationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                      // unique identifier for the consultation
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                // unique identifier for the consultation
	PatientId     string                 `protobuf:"bytes,3,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`       // from master-service
	PatientName   string                 `protobuf:"bytes,4,opt,name=patient_name,json=patientName,proto3" json:"patient_name,omitempty"` // snapshot, optional
	DoctorId      string                 `protobuf:"bytes,5,opt,name=doctor_id,json=doctorId,proto3" json:"doctor_id,omitempty"`          // from master-service
	DoctorName    string                 `protobuf:"bytes,6,opt,name=doctor_name,json=doctorName,proto3" json:"doctor_name,omitempty"`    // snapshot, optional
	RoomId        string                 `protobuf:"bytes,7,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`                // from master-service
	RoomName      string                 `protobuf:"bytes,8,opt,name=room_name,json=roomName,proto3" json:"room_name,omitempty"`          // snapshot, optional
	Symptoms      string                 `protobuf:"bytes,9,opt,name=symptoms,proto3" json:"symptoms,omitempty"`
	Prescription  []*Prescription        `protobuf:"bytes,10,rep,name=prescription,proto3" json:"prescription,omitempty"` // list of prescriptions
	Diagnosis     string                 `protobuf:"bytes,11,opt,name=diagnosis,proto3" json:"diagnosis,omitempty"`
	Date          string                 `protobuf:"bytes,12,opt,name=date,proto3" json:"date,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConsultationResponse) Reset() {
	*x = ConsultationResponse{}
	mi := &file_proto_consultation_consultation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConsultationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultationResponse) ProtoMessage() {}

func (x *ConsultationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consultation_consultation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultationResponse.ProtoReflect.Descriptor instead.
func (*ConsultationResponse) Descriptor() ([]byte, []int) {
	return file_proto_consultation_consultation_proto_rawDescGZIP(), []int{2}
}

func (x *ConsultationResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConsultationResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ConsultationResponse) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *ConsultationResponse) GetPatientName() string {
	if x != nil {
		return x.PatientName
	}
	return ""
}

func (x *ConsultationResponse) GetDoctorId() string {
	if x != nil {
		return x.DoctorId
	}
	return ""
}

func (x *ConsultationResponse) GetDoctorName() string {
	if x != nil {
		return x.DoctorName
	}
	return ""
}

func (x *ConsultationResponse) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *ConsultationResponse) GetRoomName() string {
	if x != nil {
		return x.RoomName
	}
	return ""
}

func (x *ConsultationResponse) GetSymptoms() string {
	if x != nil {
		return x.Symptoms
	}
	return ""
}

func (x *ConsultationResponse) GetPrescription() []*Prescription {
	if x != nil {
		return x.Prescription
	}
	return nil
}

func (x *ConsultationResponse) GetDiagnosis() string {
	if x != nil {
		return x.Diagnosis
	}
	return ""
}

func (x *ConsultationResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ConsultationResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ConsultationResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ConsultationIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // unique identifier for the consultation
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConsultationIDRequest) Reset() {
	*x = ConsultationIDRequest{}
	mi := &file_proto_consultation_consultation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConsultationIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultationIDRequest) ProtoMessage() {}

func (x *ConsultationIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consultation_consultation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultationIDRequest.ProtoReflect.Descriptor instead.
func (*ConsultationIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_consultation_consultation_proto_rawDescGZIP(), []int{3}
}

func (x *ConsultationIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_consultation_consultation_proto protoreflect.FileDescriptor

const file_proto_consultation_consultation_proto_rawDesc = "" +
	"\n" +
	"%proto/consultation/consultation.proto\x12\fconsultation\x1a\x1fgoogle/protobuf/timestamp.proto\"6\n" +
	"\fPrescription\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n" +
	"\x04dose\x18\x02 \x01(\tR\x04dose\"\xf2\x02\n" +
	"\x13ConsultationRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1d\n" +
	"\n" +
	"patient_id\x18\x02 \x01(\tR\tpatientId\x12!\n" +
	"\fpatient_name\x18\x03 \x01(\tR\vpatientName\x12\x1b\n" +
	"\tdoctor_id\x18\x04 \x01(\tR\bdoctorId\x12\x1f\n" +
	"\vdoctor_name\x18\x05 \x01(\tR\n" +
	"doctorName\x12\x17\n" +
	"\aroom_id\x18\x06 \x01(\tR\x06roomId\x12\x1b\n" +
	"\troom_name\x18\a \x01(\tR\broomName\x12\x1a\n" +
	"\bsymptoms\x18\t \x01(\tR\bsymptoms\x12>\n" +
	"\fprescription\x18\b \x03(\v2\x1a.consultation.PrescriptionR\fprescription\x12\x1c\n" +
	"\tdiagnosis\x18\n" +
	" \x01(\tR\tdiagnosis\x12\x12\n" +
	"\x04date\x18\v \x01(\tR\x04date\"\xf9\x03\n" +
	"\x14ConsultationResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x1d\n" +
	"\n" +
	"patient_id\x18\x03 \x01(\tR\tpatientId\x12!\n" +
	"\fpatient_name\x18\x04 \x01(\tR\vpatientName\x12\x1b\n" +
	"\tdoctor_id\x18\x05 \x01(\tR\bdoctorId\x12\x1f\n" +
	"\vdoctor_name\x18\x06 \x01(\tR\n" +
	"doctorName\x12\x17\n" +
	"\aroom_id\x18\a \x01(\tR\x06roomId\x12\x1b\n" +
	"\troom_name\x18\b \x01(\tR\broomName\x12\x1a\n" +
	"\bsymptoms\x18\t \x01(\tR\bsymptoms\x12>\n" +
	"\fprescription\x18\n" +
	" \x03(\v2\x1a.consultation.PrescriptionR\fprescription\x12\x1c\n" +
	"\tdiagnosis\x18\v \x01(\tR\tdiagnosis\x12\x12\n" +
	"\x04date\x18\f \x01(\tR\x04date\x129\n" +
	"\n" +
	"created_at\x18\r \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x0e \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"'\n" +
	"\x15ConsultationIDRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id2\xb1\x02\n" +
	"\x13ConsultationService\x12[\n" +
	"\x12CreateConsultation\x12!.consultation.ConsultationRequest\x1a\".consultation.ConsultationResponse\x12^\n" +
	"\x13GetConsultationById\x12#.consultation.ConsultationIDRequest\x1a\".consultation.ConsultationResponse\x12]\n" +
	"\x12UpdateConsultation\x12#.consultation.ConsultationIDRequest\x1a\".consultation.ConsultationResponseB\x14Z\x12proto/consultationb\x06proto3"

var (
	file_proto_consultation_consultation_proto_rawDescOnce sync.Once
	file_proto_consultation_consultation_proto_rawDescData []byte
)

func file_proto_consultation_consultation_proto_rawDescGZIP() []byte {
	file_proto_consultation_consultation_proto_rawDescOnce.Do(func() {
		file_proto_consultation_consultation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_consultation_consultation_proto_rawDesc), len(file_proto_consultation_consultation_proto_rawDesc)))
	})
	return file_proto_consultation_consultation_proto_rawDescData
}

var file_proto_consultation_consultation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_consultation_consultation_proto_goTypes = []any{
	(*Prescription)(nil),          // 0: consultation.Prescription
	(*ConsultationRequest)(nil),   // 1: consultation.ConsultationRequest
	(*ConsultationResponse)(nil),  // 2: consultation.ConsultationResponse
	(*ConsultationIDRequest)(nil), // 3: consultation.ConsultationIDRequest
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_proto_consultation_consultation_proto_depIdxs = []int32{
	0, // 0: consultation.ConsultationRequest.prescription:type_name -> consultation.Prescription
	0, // 1: consultation.ConsultationResponse.prescription:type_name -> consultation.Prescription
	4, // 2: consultation.ConsultationResponse.created_at:type_name -> google.protobuf.Timestamp
	4, // 3: consultation.ConsultationResponse.updated_at:type_name -> google.protobuf.Timestamp
	1, // 4: consultation.ConsultationService.CreateConsultation:input_type -> consultation.ConsultationRequest
	3, // 5: consultation.ConsultationService.GetConsultationById:input_type -> consultation.ConsultationIDRequest
	3, // 6: consultation.ConsultationService.UpdateConsultation:input_type -> consultation.ConsultationIDRequest
	2, // 7: consultation.ConsultationService.CreateConsultation:output_type -> consultation.ConsultationResponse
	2, // 8: consultation.ConsultationService.GetConsultationById:output_type -> consultation.ConsultationResponse
	2, // 9: consultation.ConsultationService.UpdateConsultation:output_type -> consultation.ConsultationResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_consultation_consultation_proto_init() }
func file_proto_consultation_consultation_proto_init() {
	if File_proto_consultation_consultation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_consultation_consultation_proto_rawDesc), len(file_proto_consultation_consultation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_consultation_consultation_proto_goTypes,
		DependencyIndexes: file_proto_consultation_consultation_proto_depIdxs,
		MessageInfos:      file_proto_consultation_consultation_proto_msgTypes,
	}.Build()
	File_proto_consultation_consultation_proto = out.File
	file_proto_consultation_consultation_proto_goTypes = nil
	file_proto_consultation_consultation_proto_depIdxs = nil
}
