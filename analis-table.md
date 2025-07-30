📘 ERD Title: Appointment & Queue Management System
✅ Attributes
Entity: Users
Attributes:

- ID (PK): UUID
- Name: VARCHAR(100)
- Email: VARCHAR(100) UNIQUE
- Role: ENUM('admin', 'doctor', 'patient')
- Password: TEXT
- CreatedAt: TIMESTAMP
- UpdatedAt: TIMESTAMP
  
Entity: Patients
Attributes:

- ID (PK): UUID
- UserID (FK): UUID
- Name: VARCHAR(100)
- Email: VARCHAR(100) UNIQUE
- BirthDate: DATE
- Gender: VARCHAR(10)
- PhoneNumber: VARCHAR(20)
- Address: TEXT
- CreatedAt: TIMESTAMP
- UpdatedAt: TIMESTAMP

Entity: Doctors
Attributes:

- ID (PK): UUID
- UserID (FK): UUID
- Name: VARCHAR(100)
- Email: VARCHAR(100) UNIQUE
- SpecializationID (FK): UUID
- Age: INT
- ConsultationFee: NUMERIC(12,2)
- YearsOfExperience: INT
- LicenseNumber: VARCHAR(50)
- PhoneNumber: VARCHAR(20)
- IsActive: BOOLEAN
- CreatedAt: TIMESTAMP
- UpdatedAt: TIMESTAMP

Entity: Specializations
Attributes:

- ID (PK): UUID
- Name: VARCHAR(100) UNIQUE
- CreatedAt: TIMESTAMP
- UpdatedAt: TIMESTAMP

Entity: Appointments
Attributes:

- ID (PK): SERIAL
- UserID (FK): INT
- DoctorID (FK): INT
- UserFullName: VARCHAR(100)
- DoctorName: VARCHAR(100)
- DoctorSpecialization: VARCHAR(100)
- AppointmentAt: TIMESTAMP
- IsPaid: BOOLEAN
- Status: VARCHAR(20) — ‘paid’, ‘unpaid’, ‘void’
- CreatedAt: TIMESTAMP
- UpdatedAt: TIMESTAMP (nullable)

Entity: Queues
Attributes:

- ID (PK): SERIAL
- UserID (FK): INT
- UserName: VARCHAR(100)
- UserRole: VARCHAR(20)
- PatientID (FK): UUID (nullable)
- PatientName: VARCHAR(100) (nullable)
- DoctorID (FK): INT
- DoctorName: VARCHAR(100)
- DoctorSpecialization: VARCHAR(100)
- AppointmentID (FK): INT (nullable)
- Type: VARCHAR(20) — ‘online’, ‘offline’
- QueueNumber: INTEGER
- StartFrom: TIMESTAMP
- Status: VARCHAR(20) — ‘active’, ‘fail’, ‘done’
- CreatedAt: TIMESTAMP
- UpdatedAt: TIMESTAMP

Entity: Payments
Attributes:

- ID (PK): UUID
- ConsultationID (FK): VARCHAR(24)
- ConsultationDate: DATETIME
- PatientID (FK): UUID
- DoctorID (FK): UUID
- Amount: DECIMAL(10,2)
- Method: VARCHAR(50)
- Gateway: VARCHAR(50)
- Status: ENUM('pending', 'completed', 'failed')
- CreatedAt: TIMESTAMP
- UpdatedAt: TIMESTAMP

Entity: PaymentLogs
Attributes:

- ID (PK): UUID
- PaymentID (FK): UUID
- Response: JSON
- CreatedAt: TIMESTAMP
- UpdatedAt: TIMESTAMP

Entity: Notifications
Attributes:

- ID (PK): UUID
- Channel: VARCHAR(30)
- Recipient: VARCHAR(255)
- TemplateName: VARCHAR(100)
- Subject: TEXT
- Body: TEXT
- SourceService: VARCHAR(50)
- ReferenceID: UUID (soft FK)
- Context: JSON
- Status: VARCHAR(20)
- ErrorMessage: TEXT
- RetryCount: INT
- SentAt: DATETIME
- CreatedAt: TIMESTAMP
- UpdatedAt: TIMESTAMP

🔄 Relationships
Users and Patients
- One-to-One
- Setiap patient berasal dari user (pasien memiliki user ID)

Users and Doctors
- One-to-One
- Setiap dokter adalah user juga

Users and Appointments
- One-to-Many
- Satu user bisa buat banyak janji temu

Users and Queues
- One-to-Many
- Satu user bisa mengantri berkali-kali

Patients and Queues
- One-to-Many (nullable)
- Admin bisa membuat antrean untuk pasien tertentu

Doctors and Appointments
- One-to-Many
- Dokter bisa punya banyak janji temu

Doctors and Queues
- One-to-Many
- Dokter punya banyak antrean

Specializations and Doctors
- One-to-Many
- Spesialisasi bisa dimiliki banyak dokter

Appointments and Queues
- One-to-One (nullable)
- Jika antrean dibuat dari appointment, maka antrean mengacu ke appointment ID

Appointments and Payments
- One-to-One (soft)
- ConsultationID pada payments mengacu ke appointment (jika dipakai)

Payments and PaymentLogs
- One-to-Many
- Satu pembayaran bisa punya banyak log

Notifications
- Mengacu ke entitas manapun melalui ReferenceID (soft FK)

✅ Integrity Constraints
- Users: Email harus unik, role harus salah satu dari admin, doctor, patient
- Patients: Nama, email, dan birthdate tidak boleh kosong
- Doctors: Nama, email unik, specialization ID valid, fee tidak boleh negatif
- Appointments: Tidak boleh tanggal di masa lalu, harus punya dokter dan user
- Queues: QueueNumber harus unik per dokter per hari (optional constraint)
- Payments: Amount harus positif, konsultasi ID valid
- PaymentLogs: Wajib punya payment ID valid
- Notifications: Channel dan recipient wajib terisi

🧠 Normalization
✅ 1NF
Semua tabel memenuhi 1NF: data atomik, tidak ada kolom multivalue.

✅ 2NF
Semua atribut non-primary sepenuhnya tergantung pada primary key di masing-masing tabel.

✅ 3NF
Tidak ada atribut non-primary yang bergantung transitif pada atribut non-kunci.

