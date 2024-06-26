create table directory_institution(
    directory_institution_id INT PRIMARY KEY,
    institution_id INT REFERENCES institution(institution_id) NOT NULL,
    directory_id INT REFERENCES directory_tree(directory_id) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)