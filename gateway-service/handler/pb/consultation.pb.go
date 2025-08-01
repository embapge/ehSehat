// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: proto/consultation/consultation.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type UserSnapshot struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`       // unique identifier for the user
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`   // name of the user
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"` // email of the user
	Role          string                 `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`   // role of the user
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserSnapshot) Reset() {
	*x = UserSnapshot{}
	mi := &file_proto_consultation_consultation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSnapshot) ProtoMessage() {}

func (x *UserSnapshot) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserSnapshot.ProtoReflect.Descriptor instead.
func (*UserSnapshot) Descriptor() ([]byte, []int) {
	return file_proto_consultation_consultation_proto_rawDescGZIP(), []int{0}
}

func (x *UserSnapshot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserSnapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserSnapshot) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserSnapshot) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type DoctorSnapshot struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                         // unique identifier for the doctor
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                     // name of the doctor
	Specialization string                 `protobuf:"bytes,3,opt,name=specialization,proto3" json:"specialization,omitempty"` // specialization of the doctor
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DoctorSnapshot) Reset() {
	*x = DoctorSnapshot{}
	mi := &file_proto_consultation_consultation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DoctorSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoctorSnapshot) ProtoMessage() {}

func (x *DoctorSnapshot) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DoctorSnapshot.ProtoReflect.Descriptor instead.
func (*DoctorSnapshot) Descriptor() ([]byte, []int) {
	return file_proto_consultation_consultation_proto_rawDescGZIP(), []int{1}
}

func (x *DoctorSnapshot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DoctorSnapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DoctorSnapshot) GetSpecialization() string {
	if x != nil {
		return x.Specialization
	}
	return ""
}

type PatientSnapshot struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`     // unique identifier for the patient
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // name of the patient
	Age           string                 `protobuf:"bytes,3,opt,name=age,proto3" json:"age,omitempty"`   // age of the patient
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PatientSnapshot) Reset() {
	*x = PatientSnapshot{}
	mi := &file_proto_consultation_consultation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PatientSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientSnapshot) ProtoMessage() {}

func (x *PatientSnapshot) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PatientSnapshot.ProtoReflect.Descriptor instead.
func (*PatientSnapshot) Descriptor() ([]byte, []int) {
	return file_proto_consultation_consultation_proto_rawDescGZIP(), []int{2}
}

func (x *PatientSnapshot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PatientSnapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PatientSnapshot) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

type RoomSnapshot struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`     // unique identifier for the room
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // name of the room
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoomSnapshot) Reset() {
	*x = RoomSnapshot{}
	mi := &file_proto_consultation_consultation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomSnapshot) ProtoMessage() {}

func (x *RoomSnapshot) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RoomSnapshot.ProtoReflect.Descriptor instead.
func (*RoomSnapshot) Descriptor() ([]byte, []int) {
	return file_proto_consultation_consultation_proto_rawDescGZIP(), []int{3}
}

func (x *RoomSnapshot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RoomSnapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Prescription struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // name of the medicine
	Dose          string                 `protobuf:"bytes,2,opt,name=dose,proto3" json:"dose,omitempty"` // dosage instructions
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Prescription) Reset() {
	*x = Prescription{}
	mi := &file_proto_consultation_consultation_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Prescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prescription) ProtoMessage() {}

func (x *Prescription) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consultation_consultation_proto_msgTypes[4]
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
	return file_proto_consultation_consultation_proto_rawDescGZIP(), []int{4}
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
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	QueueId       string                 `protobuf:"bytes,2,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	AppointmentId string                 `protobuf:"bytes,3,opt,name=appointment_id,json=appointmentId,proto3" json:"appointment_id,omitempty"`
	Patient       *PatientSnapshot       `protobuf:"bytes,4,opt,name=patient,proto3" json:"patient,omitempty"` // from master-service
	Doctor        *DoctorSnapshot        `protobuf:"bytes,5,opt,name=doctor,proto3" json:"doctor,omitempty"`   // from master-service
	Room          *RoomSnapshot          `protobuf:"bytes,6,opt,name=room,proto3" json:"room,omitempty"`       // from master-service
	Symptoms      string                 `protobuf:"bytes,7,opt,name=symptoms,proto3" json:"symptoms,omitempty"`
	Prescription  []*Prescription        `protobuf:"bytes,8,rep,name=prescription,proto3" json:"prescription,omitempty"` // list of prescriptions
	Diagnosis     string                 `protobuf:"bytes,9,opt,name=diagnosis,proto3" json:"diagnosis,omitempty"`
	Date          string                 `protobuf:"bytes,10,opt,name=date,proto3" json:"date,omitempty"`     // date of the consultation
	Amount        string                 `protobuf:"bytes,11,opt,name=amount,proto3" json:"amount,omitempty"` // amount for the consultation
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConsultationRequest) Reset() {
	*x = ConsultationRequest{}
	mi := &file_proto_consultation_consultation_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConsultationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultationRequest) ProtoMessage() {}

func (x *ConsultationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consultation_consultation_proto_msgTypes[5]
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
	return file_proto_consultation_consultation_proto_rawDescGZIP(), []int{5}
}

func (x *ConsultationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConsultationRequest) GetQueueId() string {
	if x != nil {
		return x.QueueId
	}
	return ""
}

func (x *ConsultationRequest) GetAppointmentId() string {
	if x != nil {
		return x.AppointmentId
	}
	return ""
}

func (x *ConsultationRequest) GetPatient() *PatientSnapshot {
	if x != nil {
		return x.Patient
	}
	return nil
}

func (x *ConsultationRequest) GetDoctor() *DoctorSnapshot {
	if x != nil {
		return x.Doctor
	}
	return nil
}

func (x *ConsultationRequest) GetRoom() *RoomSnapshot {
	if x != nil {
		return x.Room
	}
	return nil
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

func (x *ConsultationRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type ConsultationResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Id            string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // unique identifier for the consultation
	QueueId       *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	AppointmentId *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=appointment_id,json=appointmentId,proto3" json:"appointment_id,omitempty"`
	User          *UserSnapshot           `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`       // user who created the consultation
	Patient       *PatientSnapshot        `protobuf:"bytes,5,opt,name=patient,proto3" json:"patient,omitempty"` // from master-service
	Doctor        *DoctorSnapshot         `protobuf:"bytes,6,opt,name=doctor,proto3" json:"doctor,omitempty"`   // from master-service
	Room          *RoomSnapshot           `protobuf:"bytes,7,opt,name=room,proto3" json:"room,omitempty"`       // from master-service
	Symptoms      string                  `protobuf:"bytes,8,opt,name=symptoms,proto3" json:"symptoms,omitempty"`
	Prescription  []*Prescription         `protobuf:"bytes,9,rep,name=prescription,proto3" json:"prescription,omitempty"` // list of prescriptions
	Diagnosis     string                  `protobuf:"bytes,10,opt,name=diagnosis,proto3" json:"diagnosis,omitempty"`
	Date          string                  `protobuf:"bytes,11,opt,name=date,proto3" json:"date,omitempty"`
	Status        string                  `protobuf:"bytes,12,opt,name=status,proto3" json:"status,omitempty"`
	Amount        string                  `protobuf:"bytes,13,opt,name=amount,proto3" json:"amount,omitempty"`                                 // amount for the consultation
	TotalPayment  string                  `protobuf:"bytes,14,opt,name=total_payment,json=totalPayment,proto3" json:"total_payment,omitempty"` // total payment amount
	CreatedAt     *timestamppb.Timestamp  `protobuf:"bytes,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp  `protobuf:"bytes,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConsultationResponse) Reset() {
	*x = ConsultationResponse{}
	mi := &file_proto_consultation_consultation_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConsultationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultationResponse) ProtoMessage() {}

func (x *ConsultationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consultation_consultation_proto_msgTypes[6]
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
	return file_proto_consultation_consultation_proto_rawDescGZIP(), []int{6}
}

func (x *ConsultationResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConsultationResponse) GetQueueId() *wrapperspb.StringValue {
	if x != nil {
		return x.QueueId
	}
	return nil
}

func (x *ConsultationResponse) GetAppointmentId() *wrapperspb.StringValue {
	if x != nil {
		return x.AppointmentId
	}
	return nil
}

func (x *ConsultationResponse) GetUser() *UserSnapshot {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ConsultationResponse) GetPatient() *PatientSnapshot {
	if x != nil {
		return x.Patient
	}
	return nil
}

func (x *ConsultationResponse) GetDoctor() *DoctorSnapshot {
	if x != nil {
		return x.Doctor
	}
	return nil
}

func (x *ConsultationResponse) GetRoom() *RoomSnapshot {
	if x != nil {
		return x.Room
	}
	return nil
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

func (x *ConsultationResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ConsultationResponse) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *ConsultationResponse) GetTotalPayment() string {
	if x != nil {
		return x.TotalPayment
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
	mi := &file_proto_consultation_consultation_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConsultationIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultationIDRequest) ProtoMessage() {}

func (x *ConsultationIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consultation_consultation_proto_msgTypes[7]
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
	return file_proto_consultation_consultation_proto_rawDescGZIP(), []int{7}
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
	"%proto/consultation/consultation.proto\x12\fconsultation\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\\\n" +
	"\fUserSnapshot\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x12\n" +
	"\x04role\x18\x04 \x01(\tR\x04role\"\\\n" +
	"\x0eDoctorSnapshot\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12&\n" +
	"\x0especialization\x18\x03 \x01(\tR\x0especialization\"G\n" +
	"\x0fPatientSnapshot\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x10\n" +
	"\x03age\x18\x03 \x01(\tR\x03age\"2\n" +
	"\fRoomSnapshot\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\"6\n" +
	"\fPrescription\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n" +
	"\x04dose\x18\x02 \x01(\tR\x04dose\"\xac\x03\n" +
	"\x13ConsultationRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x19\n" +
	"\bqueue_id\x18\x02 \x01(\tR\aqueueId\x12%\n" +
	"\x0eappointment_id\x18\x03 \x01(\tR\rappointmentId\x127\n" +
	"\apatient\x18\x04 \x01(\v2\x1d.consultation.PatientSnapshotR\apatient\x124\n" +
	"\x06doctor\x18\x05 \x01(\v2\x1c.consultation.DoctorSnapshotR\x06doctor\x12.\n" +
	"\x04room\x18\x06 \x01(\v2\x1a.consultation.RoomSnapshotR\x04room\x12\x1a\n" +
	"\bsymptoms\x18\a \x01(\tR\bsymptoms\x12>\n" +
	"\fprescription\x18\b \x03(\v2\x1a.consultation.PrescriptionR\fprescription\x12\x1c\n" +
	"\tdiagnosis\x18\t \x01(\tR\tdiagnosis\x12\x12\n" +
	"\x04date\x18\n" +
	" \x01(\tR\x04date\x12\x16\n" +
	"\x06amount\x18\v \x01(\tR\x06amount\"\xcc\x05\n" +
	"\x14ConsultationResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x127\n" +
	"\bqueue_id\x18\x02 \x01(\v2\x1c.google.protobuf.StringValueR\aqueueId\x12C\n" +
	"\x0eappointment_id\x18\x03 \x01(\v2\x1c.google.protobuf.StringValueR\rappointmentId\x12.\n" +
	"\x04user\x18\x04 \x01(\v2\x1a.consultation.UserSnapshotR\x04user\x127\n" +
	"\apatient\x18\x05 \x01(\v2\x1d.consultation.PatientSnapshotR\apatient\x124\n" +
	"\x06doctor\x18\x06 \x01(\v2\x1c.consultation.DoctorSnapshotR\x06doctor\x12.\n" +
	"\x04room\x18\a \x01(\v2\x1a.consultation.RoomSnapshotR\x04room\x12\x1a\n" +
	"\bsymptoms\x18\b \x01(\tR\bsymptoms\x12>\n" +
	"\fprescription\x18\t \x03(\v2\x1a.consultation.PrescriptionR\fprescription\x12\x1c\n" +
	"\tdiagnosis\x18\n" +
	" \x01(\tR\tdiagnosis\x12\x12\n" +
	"\x04date\x18\v \x01(\tR\x04date\x12\x16\n" +
	"\x06status\x18\f \x01(\tR\x06status\x12\x16\n" +
	"\x06amount\x18\r \x01(\tR\x06amount\x12#\n" +
	"\rtotal_payment\x18\x0e \x01(\tR\ftotalPayment\x129\n" +
	"\n" +
	"created_at\x18\x0f \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x10 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"'\n" +
	"\x15ConsultationIDRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id2\xb0\x02\n" +
	"\x13ConsultationService\x12[\n" +
	"\x12CreateConsultation\x12!.consultation.ConsultationRequest\x1a\".consultation.ConsultationResponse\x12_\n" +
	"\x14FindByIDConsultation\x12#.consultation.ConsultationIDRequest\x1a\".consultation.ConsultationResponse\x12[\n" +
	"\x12UpdateConsultation\x12!.consultation.ConsultationRequest\x1a\".consultation.ConsultationResponseB@Z>consultation-service/internal/consultation/delivery/grpc/pb;pbb\x06proto3"

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

var file_proto_consultation_consultation_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_consultation_consultation_proto_goTypes = []any{
	(*UserSnapshot)(nil),           // 0: consultation.UserSnapshot
	(*DoctorSnapshot)(nil),         // 1: consultation.DoctorSnapshot
	(*PatientSnapshot)(nil),        // 2: consultation.PatientSnapshot
	(*RoomSnapshot)(nil),           // 3: consultation.RoomSnapshot
	(*Prescription)(nil),           // 4: consultation.Prescription
	(*ConsultationRequest)(nil),    // 5: consultation.ConsultationRequest
	(*ConsultationResponse)(nil),   // 6: consultation.ConsultationResponse
	(*ConsultationIDRequest)(nil),  // 7: consultation.ConsultationIDRequest
	(*wrapperspb.StringValue)(nil), // 8: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 9: google.protobuf.Timestamp
}
var file_proto_consultation_consultation_proto_depIdxs = []int32{
	2,  // 0: consultation.ConsultationRequest.patient:type_name -> consultation.PatientSnapshot
	1,  // 1: consultation.ConsultationRequest.doctor:type_name -> consultation.DoctorSnapshot
	3,  // 2: consultation.ConsultationRequest.room:type_name -> consultation.RoomSnapshot
	4,  // 3: consultation.ConsultationRequest.prescription:type_name -> consultation.Prescription
	8,  // 4: consultation.ConsultationResponse.queue_id:type_name -> google.protobuf.StringValue
	8,  // 5: consultation.ConsultationResponse.appointment_id:type_name -> google.protobuf.StringValue
	0,  // 6: consultation.ConsultationResponse.user:type_name -> consultation.UserSnapshot
	2,  // 7: consultation.ConsultationResponse.patient:type_name -> consultation.PatientSnapshot
	1,  // 8: consultation.ConsultationResponse.doctor:type_name -> consultation.DoctorSnapshot
	3,  // 9: consultation.ConsultationResponse.room:type_name -> consultation.RoomSnapshot
	4,  // 10: consultation.ConsultationResponse.prescription:type_name -> consultation.Prescription
	9,  // 11: consultation.ConsultationResponse.created_at:type_name -> google.protobuf.Timestamp
	9,  // 12: consultation.ConsultationResponse.updated_at:type_name -> google.protobuf.Timestamp
	5,  // 13: consultation.ConsultationService.CreateConsultation:input_type -> consultation.ConsultationRequest
	7,  // 14: consultation.ConsultationService.FindByIDConsultation:input_type -> consultation.ConsultationIDRequest
	5,  // 15: consultation.ConsultationService.UpdateConsultation:input_type -> consultation.ConsultationRequest
	6,  // 16: consultation.ConsultationService.CreateConsultation:output_type -> consultation.ConsultationResponse
	6,  // 17: consultation.ConsultationService.FindByIDConsultation:output_type -> consultation.ConsultationResponse
	6,  // 18: consultation.ConsultationService.UpdateConsultation:output_type -> consultation.ConsultationResponse
	16, // [16:19] is the sub-list for method output_type
	13, // [13:16] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
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
			NumMessages:   8,
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
