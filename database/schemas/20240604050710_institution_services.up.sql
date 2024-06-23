CREATE TABLE institution_services(
    institution_id INT REFERENCES institution(institution_id) NOT NULL,
    service_id INT REFERENCES services(service_id) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (institution_id, service_id)
);
