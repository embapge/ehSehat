CREATE DATABASE teman_sehat_payments;

USE teman_sehat_payments;

CREATE TABLE IF NOT EXISTS payments (
    id VARCHAR(36) PRIMARY KEY,
    consultation_id VARCHAR(24) NOT NULL,
    consultation_date DATETIME,
    patient_id VARCHAR(36) NOT NULL,
    patient_name VARCHAR(100),
    doctor_id VARCHAR(36) NOT NULL,
    doctor_name VARCHAR(100),
    amount DECIMAL(10, 2) NOT NULL,
    method VARCHAR(50) NOT NULL,
    gateway VARCHAR(50) NOT NULL,
    status ENUM('pending', 'completed', 'failed') DEFAULT 'pending',
    created_by VARCHAR(36), -- Nanti akan di set not null
    created_name VARCHAR(100),
    created_email VARCHAR(100),
    created_role VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by VARCHAR(36), -- Nanti akan di set not null
    updated_name VARCHAR(100),
    updated_email VARCHAR(100),
    updated_role VARCHAR(50),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS payment_logs (
    id VARCHAR(36) PRIMARY KEY,
    payment_id VARCHAR(36) NOT NULL,
    response JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);