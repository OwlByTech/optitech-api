CREATE TYPE extensions AS ENUM ('.pdf', '.doc', '.docx');

CREATE TABLE format(
    format_id SERIAL PRIMARY KEY,
    updated_format_id INT REFERENCES format(format_id),
    asesor_id INT REFERENCES asesor(asesor_id) NOT NULL,
    service_id INT REFERENCES services(service_id),
    format_name VARCHAR(50) NOT NULL,
    description VARCHAR(255) NOT NULL,
    extension extensions NOT NULL,
    version VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);