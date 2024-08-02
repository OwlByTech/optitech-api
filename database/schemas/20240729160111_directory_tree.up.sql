CREATE TABLE directory_tree(
    directory_id BIGSERIAL PRIMARY KEY,
    parent_id BIGINT REFERENCES directory_tree(directory_id),
    institution_id INT REFERENCES institution(institution_id),
    name VARCHAR(255),
    asesor_id INT REFERENCES asesor(asesor_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);