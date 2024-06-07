CREATE TABLE institution (
  institution_id BIGSERIAL PRIMARY KEY,
  asesor_id INT REFERENCES asesor(asesor_id),
  institution_name VARCHAR(50) NOT NULL,
  logo VARCHAR(255),
  description VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);