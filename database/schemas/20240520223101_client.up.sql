CREATE TABLE client (
  client_id BIGSERIAL PRIMARY KEY,
  given_name VARCHAR(255) NOT NULL,
  surname VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  pass VARCHAR(255) NOT NULL,
  asesor_id INTEGER REFERENCES asesor(asesor_id),
  institution_id INTEGER REFERENCES institution(institution_id),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);
