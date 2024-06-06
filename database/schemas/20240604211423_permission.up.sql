CREATE TABLE permission(
    permission_id BIGSERIAL PRIMARY KEY, 
    permission_code VARCHAR(50) NOT NULL,
    permission_description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);