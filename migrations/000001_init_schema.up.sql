-- Create enum types
CREATE TYPE user_role AS ENUM ('normal', 'doctor', 'admin');
CREATE TYPE user_status AS ENUM ('active', 'inactive', 'banned');

-- Create users table
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    surname VARCHAR(255),
    role user_role NOT NULL DEFAULT 'normal',
    status user_status NOT NULL DEFAULT 'active',
    last_login TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create shift_locations table
CREATE TABLE shift_locations (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create doctors table
CREATE TABLE doctors (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    specialization VARCHAR(255),
    title VARCHAR(255),
    shift_limit INTEGER,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create doctor_shift_locations table
CREATE TABLE doctor_shift_locations (
    id BIGSERIAL PRIMARY KEY,
    doctor_id BIGINT NOT NULL REFERENCES doctors(id),
    location_id BIGINT NOT NULL REFERENCES shift_locations(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create shifts table
CREATE TABLE shifts (
    id BIGSERIAL PRIMARY KEY,
    doctor_id BIGINT NOT NULL REFERENCES doctors(id),
    location_id BIGINT NOT NULL REFERENCES shift_locations(id),
    shift_date DATE NOT NULL,
    start_time VARCHAR(5) NOT NULL, -- Format: HH:MM
    end_time VARCHAR(5) NOT NULL,   -- Format: HH:MM
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create holidays table
CREATE TABLE holidays (
    id BIGSERIAL PRIMARY KEY,
    doctor_id BIGINT NOT NULL REFERENCES doctors(id),
    location_id BIGINT NOT NULL REFERENCES shift_locations(id),
    holiday_date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create shifts_status table
CREATE TABLE shifts_status (
    id BIGSERIAL PRIMARY KEY,
    year INTEGER NOT NULL,
    month INTEGER NOT NULL CHECK (month BETWEEN 1 AND 12),
    done BOOLEAN NOT NULL DEFAULT FALSE,
    location_id BIGINT NOT NULL REFERENCES shift_locations(id),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create shift_swap_requests table
CREATE TABLE shift_swap_requests (
    id BIGSERIAL PRIMARY KEY,
    location_id BIGINT NOT NULL REFERENCES shift_locations(id),
    requester_id BIGINT NOT NULL REFERENCES doctors(id),
    request_shift_date DATE NOT NULL,
    requested_shift_id BIGINT NOT NULL REFERENCES shifts(id),
    acceptor_id BIGINT NOT NULL REFERENCES doctors(id),
    offered_shift_date DATE NOT NULL,
    offered_shift_id BIGINT NOT NULL REFERENCES shifts(id),
    mutual_agreement BOOLEAN DEFAULT FALSE,
    requester_comment TEXT,
    acceptor_comment TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create notifications table
CREATE TABLE notifications (
    id BIGSERIAL PRIMARY KEY,
    message TEXT NOT NULL,
    doctor_id BIGINT NOT NULL REFERENCES doctors(id),
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Add indexes
CREATE INDEX idx_users_email ON users(email) WHERE deleted_at IS NULL;
CREATE INDEX idx_doctors_user_id ON doctors(user_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_shifts_doctor_location ON shifts(doctor_id, location_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_shifts_date ON shifts(shift_date) WHERE deleted_at IS NULL;
CREATE INDEX idx_holidays_doctor_date ON holidays(doctor_id, holiday_date) WHERE deleted_at IS NULL;
CREATE INDEX idx_shift_swap_requests_status ON shift_swap_requests(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_notifications_doctor ON notifications(doctor_id, is_read) WHERE deleted_at IS NULL;

-- Tokens table
CREATE TABLE tokens (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    access_token TEXT NOT NULL UNIQUE,
    refresh_token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    revoked_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create indexes for tokens
CREATE INDEX idx_tokens_user_id ON tokens(user_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_tokens_access_token ON tokens(access_token) WHERE deleted_at IS NULL;
CREATE INDEX idx_tokens_refresh_token ON tokens(refresh_token) WHERE deleted_at IS NULL;

-- Sessions table
CREATE TABLE sessions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    refresh_token TEXT NOT NULL UNIQUE,
    user_agent TEXT,
    client_ip VARCHAR(45),
    is_blocked BOOLEAN NOT NULL DEFAULT FALSE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create indexes for sessions
CREATE INDEX idx_sessions_user_id ON sessions(user_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_sessions_refresh_token ON sessions(refresh_token) WHERE deleted_at IS NULL;

-- Token blacklist table
CREATE TABLE token_blacklist (
    id BIGSERIAL PRIMARY KEY,
    token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Create index for blacklisted tokens
CREATE INDEX idx_token_blacklist_token ON token_blacklist(token) WHERE deleted_at IS NULL;

-- Function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create triggers for updated_at
CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_tokens_updated_at
    BEFORE UPDATE ON tokens
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_sessions_updated_at
    BEFORE UPDATE ON sessions
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_doctors_updated_at
    BEFORE UPDATE ON doctors
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_shift_locations_updated_at
    BEFORE UPDATE ON shift_locations
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_shifts_updated_at
    BEFORE UPDATE ON shifts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_holidays_updated_at
    BEFORE UPDATE ON holidays
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_shift_swap_requests_updated_at
    BEFORE UPDATE ON shift_swap_requests
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_notifications_updated_at
    BEFORE UPDATE ON notifications
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column(); 