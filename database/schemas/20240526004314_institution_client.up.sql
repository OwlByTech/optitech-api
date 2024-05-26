CREATE TABLE institution_client (
    institution_client_id BIGSERIAL PRIMARY KEY,
    client_id INT REFERENCES client(client_id) NOT NULL,
    institution_id INT REFERENCES institution(institution_id) NOT NULL,
    vinculated_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP
);