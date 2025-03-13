-- Drop triggers
DROP TRIGGER IF EXISTS update_users_updated_at ON users;
DROP TRIGGER IF EXISTS update_tokens_updated_at ON tokens;
DROP TRIGGER IF EXISTS update_sessions_updated_at ON sessions;
DROP TRIGGER IF EXISTS update_doctors_updated_at ON doctors;
DROP TRIGGER IF EXISTS update_shift_locations_updated_at ON shift_locations;
DROP TRIGGER IF EXISTS update_shifts_updated_at ON shifts;
DROP TRIGGER IF EXISTS update_holidays_updated_at ON holidays;
DROP TRIGGER IF EXISTS update_shift_swap_requests_updated_at ON shift_swap_requests;
DROP TRIGGER IF EXISTS update_notifications_updated_at ON notifications;

-- Drop function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop tables in reverse order (to handle foreign key constraints)
DROP TABLE IF EXISTS token_blacklist;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS tokens;
DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS shift_swap_requests;
DROP TABLE IF EXISTS shifts_status;
DROP TABLE IF EXISTS holidays;
DROP TABLE IF EXISTS shifts;
DROP TABLE IF EXISTS doctor_shift_locations;
DROP TABLE IF EXISTS doctors;
DROP TABLE IF EXISTS shift_locations;
DROP TABLE IF EXISTS users;

-- Drop enum types
DROP TYPE IF EXISTS user_status;
DROP TYPE IF EXISTS user_role; 