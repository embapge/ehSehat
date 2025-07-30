CREATE DATABASE teman_sehat_appointment_queues;

USE teman_sehat_appointment_queues;

CREATE TABLE appointments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    user_full_name VARCHAR(100) NOT NULL,
    doctor_id INTEGER NOT NULL,
    doctor_name VARCHAR(100) NOT NULL,
    doctor_specialization VARCHAR(100) NOT NULL,
    appointment_at TIMESTAMP NOT NULL,
    is_paid BOOLEAN DEFAULT FALSE,
    status VARCHAR(20) NOT NULL, -- e.g., 'paid', 'unpaid', 'void'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE queues (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    user_name VARCHAR(100) NOT NULL,
    user_role VARCHAR(20) NOT NULL, -- admin/member
    patient_id INTEGER,             -- nullable
    patient_name VARCHAR(100),      -- nullable
    doctor_id INTEGER NOT NULL,
    doctor_name VARCHAR(100) NOT NULL,
    doctor_specialization VARCHAR(100) NOT NULL,
    appointment_id INTEGER,         -- nullable
    type VARCHAR(20) NOT NULL,      -- online / offline
    queue_number INTEGER NOT NULL,
    start_from TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL,    -- active / fail / done
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
