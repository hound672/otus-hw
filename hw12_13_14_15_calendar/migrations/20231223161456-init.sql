
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4()
);

CREATE TABLE events (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    start_at TIMESTAMP NOT NULL,
    end_at TIMESTAMP NOT NULL,
    description TEXT,
    user_uuid  UUID, -- TODO: add reference after implement users
    notify_before int
);

-- +migrate Down
DROP TABLE events;
DROP TABLE users;
