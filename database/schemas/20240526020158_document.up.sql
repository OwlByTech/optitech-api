CREATE TYPE status AS ENUM ('aprobado', 'en revision', 'rechazado');

CREATE TABLE document(
    document_id BIGSERIAL PRIMARY KEY,
    directory_id INT REFERENCES directory_tree(directory_id) NOT NULL,
    format_id INT REFERENCES format(format_id),
    url VARCHAR(255) NOT NULL,
    status status NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);