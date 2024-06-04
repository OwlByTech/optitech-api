CREATE TABLE institution_services(
    institution_services_id BIGSERIAL PRIMARY KEY,
    institution_id INT REFERENCES institution(institution_id),
    services_id INT REFERENCES services(services_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);
