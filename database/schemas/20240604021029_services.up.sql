CREATE TABLE services(
    service_id SERIAL PRIMARY KEY,
    service_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

