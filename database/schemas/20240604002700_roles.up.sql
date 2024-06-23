CREATE TABLE roles(
    role_id SERIAL PRIMARY KEY, 
    role_name VARCHAR(50) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);