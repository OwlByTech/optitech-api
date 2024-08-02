CREATE TABLE role_permission(
    role_permission_id BIGSERIAL PRIMARY KEY, 
    role_id INT REFERENCES roles(role_id) NOT NULL,
    permission_id INT REFERENCES permission(permission_id) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);