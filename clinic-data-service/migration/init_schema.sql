-- ============================
-- CREATE TABLE: patients
-- ============================

CREATE TABLE IF NOT EXISTS patients (
  id UUID PRIMARY KEY,
  user_id UUID,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE,
  birth_date DATE NOT NULL,
  gender VARCHAR(10) NOT NULL,
  phone_number VARCHAR(20),
  address TEXT,
  
  created_by UUID,
  created_name VARCHAR(100),
  created_email VARCHAR(100),
  created_role VARCHAR(20),
  updated_by UUID,
  updated_name VARCHAR(100),
  updated_email VARCHAR(100),
  updated_role VARCHAR(20),
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- CREATE TABLE: specializations
-- =========================

CREATE TABLE IF NOT EXISTS specializations (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,

    created_by UUID,
    created_name VARCHAR(100),
    created_email VARCHAR(100),
    created_role VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_by UUID,
    updated_name VARCHAR(100),
    updated_email VARCHAR(100),
    updated_role VARCHAR(50),
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- CREATE TABLE: doctors
-- =========================

CREATE TABLE IF NOT EXISTS doctors (
  id UUID PRIMARY KEY,
  user_id UUID,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL UNIQUE,
  specialization_id UUID NOT NULL REFERENCES specializations(id),
  age INT NOT NULL,
  consultation_fee NUMERIC(12, 2) NOT NULL,
  years_of_experience INT NOT NULL,
  license_number VARCHAR(50) NOT NULL,
  phone_number VARCHAR(20),
  is_active BOOLEAN DEFAULT FALSE,

  created_by UUID,
  created_name VARCHAR(100),
  created_email VARCHAR(100),
  created_role VARCHAR(20),
  updated_by UUID,
  updated_name VARCHAR(100),
  updated_email VARCHAR(100),
  updated_role VARCHAR(20),
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- =========================
-- CREATE TABLE: rooms
-- =========================

CREATE TABLE IF NOT EXISTS rooms (
  id UUID PRIMARY KEY,
  name VARCHAR(100) NOT NULL UNIQUE,
  is_active BOOLEAN DEFAULT TRUE,

  created_by UUID,
  created_name VARCHAR(100),
  created_email VARCHAR(100),
  created_role VARCHAR(20),
  updated_by UUID,
  updated_name VARCHAR(100),
  updated_email VARCHAR(100),
  updated_role VARCHAR(20),
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ======================================
-- TABLE: schedule_fixed
-- ======================================
CREATE TABLE IF NOT EXISTS schedule_fixed (
  id UUID PRIMARY KEY,
  doctor_id UUID NOT NULL REFERENCES doctors(id),
  room_id UUID NOT NULL REFERENCES rooms(id),
  day_of_week INT NOT NULL CHECK (day_of_week BETWEEN 1 AND 7),
  start_time TIME NOT NULL,
  end_time TIME NOT NULL,
  status VARCHAR(20) NOT NULL DEFAULT 'active',

  created_by UUID,
  created_name VARCHAR(100),
  created_email VARCHAR(100),
  created_role VARCHAR(20),
  updated_by UUID,
  updated_name VARCHAR(100),
  updated_email VARCHAR(100),
  updated_role VARCHAR(20),

  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ======================================
-- TABLE: schedule_override
-- ======================================
CREATE TABLE IF NOT EXISTS schedule_overrides (
  id UUID PRIMARY KEY,
  doctor_id UUID NOT NULL REFERENCES doctors(id),
  room_id UUID NOT NULL REFERENCES rooms(id),
  day_of_week INT NOT NULL CHECK (day_of_week BETWEEN 1 AND 7),
  start_time TIME NOT NULL,
  end_time TIME NOT NULL,
  status VARCHAR(20) NOT NULL DEFAULT 'active',

  created_by UUID,
  created_name VARCHAR(100),
  created_email VARCHAR(100),
  created_role VARCHAR(20),
  updated_by UUID,
  updated_name VARCHAR(100),
  updated_email VARCHAR(100),
  updated_role VARCHAR(20),

  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);