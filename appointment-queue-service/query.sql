CREATE DATABASE IF NOT EXISTS teman_sehat_appointment_queues;

USE teman_sehat_appointment_queues;

CREATE TABLE IF NOT EXISTS appointments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    user_full_name VARCHAR(100),
    doctor_id INTEGER NOT NULL,
    doctor_name VARCHAR(100) NOT NULL,
    doctor_specialization VARCHAR(100) NOT NULL,
    appointment_at TIMESTAMP NOT NULL,
    is_paid BOOLEAN DEFAULT FALSE,
    status VARCHAR(20) NOT NULL, -- e.g., 'paid', 'unpaid', 'void'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS queues (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(36),
    user_name VARCHAR(100),
    user_role VARCHAR(20), -- admin/member
    user_email VARCHAR(20), -- admin/member
    patient_id VARCHAR(36),             -- nullable
    patient_name VARCHAR(100),      -- nullable
    doctor_id VARCHAR(36),
    doctor_name VARCHAR(100),
    doctor_specialization VARCHAR(100),
    appointment_id INTEGER,         -- nullable
    type VARCHAR(20) NOT NULL,      -- online / offline
    queue_number INTEGER NOT NULL,
    start_from TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL,    -- active / fail / done
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
