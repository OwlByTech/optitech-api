CREATE TABLE asesor (
  asesor_id BIGSERIAL PRIMARY KEY,
  client_id INT REFERENCES client(client_id) NOT NULL,
  username VARCHAR(50) NOT NULL,
  photo VARCHAR(255) NOT NULL,
  about VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,  
  deleted_at TIMESTAMP,
);
