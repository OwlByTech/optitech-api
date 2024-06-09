CREATE TABLE standards(
    standard_id BIGSERIAL PRIMARY KEY,
    service_id INT REFERENCES services(services_id) NOT NULL,
    standard VARCHAR(255) NOT NULL,
    complexity VARCHAR(255),
    modality VARCHAR(255) NOT NULL,
    article VARCHAR(30) NOT NULL,
    section VARCHAR(30) NOT NULL,
    paragraph VARCHAR(30),
    criteria VARCHAR(1000) NOT NULL,
    comply BOOLEAN,
    applys BOOLEAN,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)