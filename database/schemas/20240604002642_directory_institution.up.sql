create table directory_institution(
    directory_institution_id BIGSERIAL PRIMARY KEY,
    institution_id INT REFERENCES institution(institution_id),
    directory_id INT REFERENCES directory_tree(directory_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)