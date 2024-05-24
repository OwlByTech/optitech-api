CREATE TABLE asesor (
  asesor_id BIGSERIAL PRIMARY KEY,
  client_id INT REFERENCES client(client_id) NOT NULL,
  username VARCHAR(255) NOT NULL,
  photo VARCHAR(255) NOT NULL,
  about VARCHAR(255) NOT NULL,
  create_at TIMESTAMP NOT NULL,
  update_at TIMESTAMP
);
