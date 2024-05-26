CREATE TYPE status AS ENUM ('aprobado', 'en revision', 'rechazado');

CREATE TABLE document(
    document_id BIGSERIAL PRIMARY KEY,
    format_id INT REFERENCES format(format_id) NOT NULL,
    institution_id INT REFERENCES institution(institution_id) NOT NULL,
    client_id INT REFERENCES client(client_id) NOT NULL,
    file_rute VARCHAR(255) NOT NULL,
    status status, 
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP
);