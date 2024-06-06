CREATE TABLE membership(
    membership_id BIGSERIAL PRIMARY KEY,
    membership_type_id INT REFERENCES membership_type(membership_type_id),
    created_at TIMESTAMP NOT NULL,
    finish_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)