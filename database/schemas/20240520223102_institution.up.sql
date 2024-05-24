CREATE TABLE institution (
  institution_id BIGSERIAL PRIMARY KEY,
  asesor_id INT REFERENCES asesor(asesor_id),
  institution_name VARCHAR(255) NOT NULL,
  logo VARCHAR(255),
  description  VARCHAR(255) NOT NULL,
  services text[] NOT NULL,
  create_at TIMESTAMP NOT NULL,
  update_at TIMESTAMP
);
