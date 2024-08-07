CREATE TYPE from_notification AS ENUM ('institution', 'asesor', 'super_user');
CREATE TYPE to_notification AS ENUM ('institution', 'asesor', 'super_user');
CREATE TYPE type_notification AS ENUM ('information', 'correction', 'error', 'aproved');

CREATE TABLE notification (
    notification_id BIGSERIAL PRIMARY KEY,
    "from" from_notification DEFAULT 'super_user' NOT NULL,
    "to" to_notification DEFAULT 'super_user' NOT NULL,
    from_id INT NOT NULL,
    to_id INT NOT NULL,
    message VARCHAR(255) NOT NULL,
    title VARCHAR(50) NOT NULL,
    visualized BOOLEAN DEFAULT FALSE,
    payload JSONB NOT NULL,
    type type_notification DEFAULT 'information',
    created_at TIMESTAMP NOT NULL
);