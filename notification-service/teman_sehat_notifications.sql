CREATE DATABASE IF NOT EXISTS teman_sehat_notifications;

USE teman_sehat_notifications;

CREATE TABLE IF NOT EXISTS notifications (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    channel VARCHAR(30) NOT NULL,
    recipient VARCHAR(255) NOT NULL,
    template_name VARCHAR(100) NOT NULL,
    subject TEXT,
    body TEXT NOT NULL,
    source_service VARCHAR(50) NOT NULL,
    reference_id CHAR(36),
    context JSON,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    error_message TEXT,
    retry_count INT DEFAULT 0,
    sent_at DATETIME(3),
    created_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)
);