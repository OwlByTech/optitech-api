CREATE TABLE directory_tree(
    directory_id BIGSERIAL PRIMARY KEY,
    parent_id INT REFERENCES directory_tree(directory_id),
    name VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);