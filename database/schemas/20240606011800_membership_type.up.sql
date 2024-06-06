CREATE TABLE membership_type(
    membership_type_id BIGSERIAL PRIMARY KEY,
    membership_name VARCHAR(255) NOT NULL,
    users INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);