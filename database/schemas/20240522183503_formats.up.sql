CREATE TYPE extensions AS ENUM ('.pdf', '.doc', '.docx');

CREATE TABLE format(
    format_id BIGSERIAL PRIMARY KEY,
    updated_format_id INT REFERENCES format(format_id), 
    asesor_id INT REFERENCES asesor(asesor_id) NOT NULL,
    format_name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    items text[] NOT NULL,
    extension extensions NOT NULL,
    version VARCHAR(255) NOT NULL,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP
);