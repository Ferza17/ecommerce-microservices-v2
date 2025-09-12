-- +goose Up


CREATE TABLE IF NOT EXISTS roles
(
    id         VARCHAR(255) PRIMARY KEY,
    role       VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP          NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP          NOT NULL DEFAULT NOW()
);

-- Create users table
CREATE TABLE IF NOT EXISTS users
(
    id           VARCHAR(255) PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    email        VARCHAR(255) NOT NULL,
    password     VARCHAR(255) NOT NULL,
    is_verified  BOOLEAN               DEFAULT FALSE,
    role_id      VARCHAR(255),
    created_at   TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP    NOT NULL DEFAULT NOW(),
    discarded_at TIMESTAMP,

    -- Foreign Key Constraints
    CONSTRAINT fk_users_role FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE SET NULL,

    -- Unique Constraints
    CONSTRAINT uk_users_email UNIQUE (email)
);

-- Create access_controls table
CREATE TABLE IF NOT EXISTS access_controls
(
    id               VARCHAR(255) PRIMARY KEY,
    full_method_name VARCHAR(255) NOT NULL,
    http_url         VARCHAR(255) NOT NULL,
    http_method      VARCHAR(10)  NOT NULL,
    event_type       VARCHAR(255),
    role_id          VARCHAR(255) NOT NULL,
    created_at       TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP    NOT NULL DEFAULT NOW(),

    -- Foreign Key Constraints
    CONSTRAINT fk_access_controls_role FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE
);


CREATE TABLE access_control_excluded
(
    id               VARCHAR(255) PRIMARY KEY,
    full_method_name VARCHAR(255) NOT NULL,
    http_url         VARCHAR(255) NOT NULL,
    http_method      VARCHAR(10)  NOT NULL,
    event_type       VARCHAR(255),
    created_at       TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP    NOT NULL DEFAULT NOW()
);


-- Create Indexes for better performance
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);
CREATE INDEX IF NOT EXISTS idx_users_role_id ON users (role_id);
CREATE INDEX IF NOT EXISTS idx_users_discarded_at ON users (discarded_at);
CREATE INDEX IF NOT EXISTS idx_access_control_excluded_discarded_at ON access_control_excluded (full_method_name, http_url, http_method);
CREATE UNIQUE INDEX IF NOT EXISTS idx_access_controls_service_method ON access_controls (full_method_name, role_id);