-- Create Users Table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,           -- Unique user ID
    username VARCHAR(255) UNIQUE NOT NULL,  -- Username, must be unique
    email VARCHAR(255) UNIQUE NOT NULL,     -- Email, must be unique
    password_hash VARCHAR(255) NOT NULL,    -- Hashed password
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Creation timestamp
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Last update timestamp
    last_login TIMESTAMP,             -- Last login timestamp
    role VARCHAR(50) DEFAULT 'user',  -- User role (default is 'user')
    status VARCHAR(20) DEFAULT 'active'  -- Account status (default is 'active')
);

-- Create Sessions Table
CREATE TABLE IF NOT EXISTS sessions (
    session_id SERIAL PRIMARY KEY,    -- Unique session ID
    user_id INT REFERENCES users(id) ON DELETE CASCADE,  -- Foreign key referencing user ID
    token VARCHAR(255) UNIQUE NOT NULL,  -- Authentication token (JWT or similar)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Session creation timestamp
    expires_at TIMESTAMP               -- Token expiration time
);

-- Indexes for faster queries
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_sessions_user_id ON sessions(user_id);
