CREATE TABLE asesor (
  asesor_id INT REFERENCES client(client_id) NOT NULL,
  about VARCHAR(255) NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  PRIMARY KEY(asesor_id)
);

