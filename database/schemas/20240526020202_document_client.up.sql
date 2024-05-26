CREATE TYPE action AS ENUM ('borrado', 'actualizado', 'creado');

CREATE TABLE document_client(
    document_client_id BIGSERIAL PRIMARY KEY,
    client_id INT REFERENCES client(client_id) NOT NULL,
    document_id INT REFERENCES document(document_id) NOT NULL,
    action action NOT NULL,
    create_at TIMESTAMP NOT NULL
);