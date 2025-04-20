-- Create users table
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       role VARCHAR(20) NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create departments table
CREATE TABLE departments (
                             id SERIAL PRIMARY KEY,
                             name VARCHAR(255) UNIQUE NOT NULL,
                             description TEXT,
                             floor INTEGER NOT NULL,
                             created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create patients table
CREATE TABLE patients (
                          id SERIAL PRIMARY KEY,
                          user_id INTEGER UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                          first_name VARCHAR(255) NOT NULL,
                          last_name VARCHAR(255) NOT NULL,
                          date_of_birth VARCHAR(255) NOT NULL,
                          gender VARCHAR(20) NOT NULL,
                          address TEXT,
                          phone_number VARCHAR(20),
                          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create doctors table
CREATE TABLE doctors (
                         id SERIAL PRIMARY KEY,
                         user_id INTEGER UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                         first_name VARCHAR(255) NOT NULL,
                         last_name VARCHAR(255) NOT NULL,
                         specialization VARCHAR(255) NOT NULL,
                         room_number VARCHAR(20) NOT NULL,
                         phone_number VARCHAR(20),
                         department_id INTEGER REFERENCES departments(id) ON DELETE SET NULL,
                         created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create appointments table
CREATE TABLE appointments (
                              id SERIAL PRIMARY KEY,
                              patient_id INTEGER NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
                              doctor_id INTEGER NOT NULL REFERENCES doctors(id) ON DELETE CASCADE,
                              department_id INTEGER NOT NULL REFERENCES departments(id) ON DELETE CASCADE,
                              appointment_date TIMESTAMP NOT NULL,
                              status VARCHAR(20) NOT NULL,
                              diagnosis TEXT,
                              prescription TEXT,
                              notes TEXT,
                              created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Insert initial admin user
INSERT INTO users (username, password, role, email)
VALUES ('admin', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'admin', 'admin@hospital.com');
