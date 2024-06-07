CREATE TABLE client_role(
    client_role_id BIGSERIAL PRIMARY KEY, 
    client_id INT REFERENCES client(client_id) NOT NULL,
    role_id INT REFERENCES roles(role_id) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);