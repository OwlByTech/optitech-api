CREATE TABLE roles(
    role_id BIGSERIAL PRIMARY KEY, 
    role_name VARCHAR(50),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);