CREATE TYPE status_client AS ENUM ('activo', 'inactivo');
CREATE TABLE client (
  client_id BIGSERIAL PRIMARY KEY,
  given_name VARCHAR(50) NOT NULL,
  surname VARCHAR(50) NOT NULL,
  email VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  status status_client DEFAULT 'inactivo',
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

