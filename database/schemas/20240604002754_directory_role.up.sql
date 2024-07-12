CREATE TYPE permissions AS ENUM (
  'r',
  'w',
  'x',
  'rw',
  'rx',
  'wx',
  'rwx'
);

CREATE TABLE directory_role(
    directory_id INT REFERENCES directory_tree(directory_id),
    user_id INT REFERENCES client(client_id),
    status permissions NOT NULL DEFAULT 'r',
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);