# Clinic Data Service

`clinic-data-service` adalah salah satu dari enam layanan backend dalam sistem aplikasi klinik. Layanan ini menangani **data master klinik**, seperti:

- Data pasien
- Data dokter
- Spesialisasi medis
- Ruangan klinik
- Jadwal tetap dan pengganti dokter

Meski menggunakan pola modular service dengan komunikasi gRPC, seluruh layanan masih menggunakan **satu database terpusat (PostgreSQL)**. Struktur ini cocok disebut sebagai **modular monolith dengan pemisahan service secara logika dan transport**, namun **belum sepenuhnya microservice**.

---

## 🧱 Arsitektur & Teknologi

- **Bahasa**: Go (Golang)
- **Database**: PostgreSQL (shared database)
- **Transport**: gRPC
- **Struktur Folder**: Domain-Driven Design (DDD)
- **Tooling**: `protoc`, `protoc-gen-go`, `protoc-gen-go-grpc`
- **Testing**: direkomendasikan via Postman (gRPC)

---

## 📁 Struktur Folder

```
clinic-data-service/
├── .env                        # Variabel environment
├── cmd/main.go                # Entry point aplikasi
├── config/env.go              # Loader konfigurasi dari .env
├── internal/clinicdata/
│   ├── domain/                # Entity dan interface repository
│   ├── app/                   # Logika bisnis & validasi
│   ├── infra/                 # Implementasi PostgreSQL
│   └── delivery/grpc/         # Handler gRPC, error mapper, audit
├── migration/init_schema.sql  # Skema tabel PostgreSQL
├── proto/clinic_data.proto    # Definisi Protobuf
├── go.mod / go.sum
```

---

## ⚙️ Menjalankan Project Secara Lokal

### 1. Setup Database PostgreSQL

```bash
createdb teman_sehat_masters
psql -U postgres -d teman_sehat_masters -f migration/init_schema.sql
```

### 2. Konfigurasi `.env`

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=123456789
DB_NAME=teman_sehat_masters
GRPC_PORT=50052
```

### 3. Jalankan Server

```bash
go run cmd/main.go
```

Output:
```
gRPC server running on :50052
```

---

## 🔌 Pengujian dengan Postman (gRPC)

1. Buka Postman → tab `gRPC`
2. Target server: `localhost:50052`
3. Metadata wajib:
   - `ts-user-id`, `ts-user-name`, `ts-user-email`, `ts-user-role`
4. Pilih method gRPC (lihat daftar di bawah)
5. Kirim request JSON sesuai dengan field di `proto`

---

## 📌 Daftar Lengkap Endpoint gRPC

Semua method didefinisikan di file `proto/clinic_data.proto`.

### PATIENT
- `ClinicDataService/CreatePatient`
- `ClinicDataService/GetPatientByID`
- `ClinicDataService/GetAllPatients`
- `ClinicDataService/UpdatePatient`
- `ClinicDataService/DeletePatient`

### DOCTOR
- `ClinicDataService/CreateDoctor`
- `ClinicDataService/GetDoctorByID`
- `ClinicDataService/GetAllDoctors`
- `ClinicDataService/UpdateDoctor`
- `ClinicDataService/DeleteDoctor`

### SPECIALIZATION
- `ClinicDataService/CreateSpecialization`
- `ClinicDataService/GetSpecializationByID`
- `ClinicDataService/GetAllSpecializations`
- `ClinicDataService/UpdateSpecialization`
- `ClinicDataService/DeleteSpecialization`

### ROOM
- `ClinicDataService/CreateRoom`
- `ClinicDataService/GetRoomByID`
- `ClinicDataService/GetAllRooms`

### SCHEDULE FIXED
- `ClinicDataService/CreateScheduleFixed`
- `ClinicDataService/GetFixedSchedulesByDoctorID`
- `ClinicDataService/UpdateScheduleFixed`

### SCHEDULE OVERRIDE
- `ClinicDataService/CreateScheduleOverride`
- `ClinicDataService/GetOverrideByDoctorID`
- `ClinicDataService/UpdateScheduleOverride`
- `ClinicDataService/DeleteScheduleOverride`

---

## 🔐 Audit Metadata

Semua request gRPC menyertakan metadata header berikut:
- `ts-user-id`, `ts-user-name`, `ts-user-email`, `ts-user-role`

Header ini akan diekstrak oleh `utils.ExtractAudit(ctx)` dan otomatis mengisi kolom:
- `created_by`, `created_name`, `created_email`, `created_role`
- `updated_by`, `updated_name`, `updated_email`, `updated_role`

---

## ✅ Validasi & Error Handling

Validasi terjadi di layer `app/` sebelum data disimpan. Contoh validasi:

- **Wajib diisi**: name, email, birth_date, gender, license_number
- **Format email valid & unik**
- **Angka**: usia ≥ 0, biaya konsultasi ≥ 0
- **Relasi valid**: ID spesialisasi, dokter, dan ruangan harus valid

Error dikembalikan dengan gRPC status code via `mapErrorToStatus()`:

| Kode Error (`app/error.go`)        | gRPC Status       |
|----------------------------------- |-------------------|
| `ErrNotFound`                      | `NotFound`        |
| `ErrMissingFields` / `ErrMissingID`| `InvalidArgument` |
| `ErrEmailAlreadyExists`            | `InvalidArgument` |
| `ErrInternal`                      | `Internal`        |
| error lainnya                      | `Unknown`         |

---

## 🧾 Struktur Tabel & Relasi (PostgreSQL)

### `patients`
- `id`, `user_id`, `name`, `email`, `birth_date`, `gender`, `phone_number`, `address`
- **Relasi**: `user_id` berasal dari `auth-service`

### `doctors`
- `specialization_id` → `specializations.id`
- `user_id` berasal dari `auth-service`

### `specializations`
- Digunakan oleh tabel `doctors`

### `rooms`
- Digunakan oleh `schedule_fixed` dan `schedule_overrides`

### `schedule_fixed`
- Jadwal mingguan tetap
- **Relasi**: `doctor_id`, `room_id`
- **Validasi**: tidak boleh bentrok per hari dan jam

### `schedule_overrides`
- Jadwal pengganti: cuti, pindah ruangan, dll
- **Relasi**: `doctor_id`, `room_id`
- Bisa override jadwal tetap

---

