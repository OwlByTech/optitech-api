CREATE TABLE directory_role(
    directory_role_id BIGSERIAL PRIMARY KEY, 
    directory_id INT REFERENCES directory_tree(directory_id),
    role_id INT REFERENCES roles(role_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);