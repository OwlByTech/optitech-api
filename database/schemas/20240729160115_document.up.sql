CREATE TYPE status AS ENUM ('aprobado', 'en revision', 'rechazado');

CREATE TABLE document(
    document_id BIGSERIAL PRIMARY KEY,
    directory_id BIGINT REFERENCES directory_tree(directory_id) NOT NULL,
    format_id INT REFERENCES format(format_id),
    name VARCHAR(100) NOT NULL,
    file_rute VARCHAR(255) NOT NULL,
    status status ,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

