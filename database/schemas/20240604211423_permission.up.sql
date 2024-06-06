CREATE TABLE permission(
    permission_id BIGSERIAL PRIMARY KEY, 
    permission_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);