CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    user_id     UUID            PRIMARY KEY     NOT NULL DEFAULT uuid_generate_v4(),
    username    VARCHAR(128)    UNIQUE          NOT NULL CHECK ( username <> '' ),
    password    VARCHAR(128)                    NOT NULL CHECK ( password <> '' ),
    email       VARCHAR(128)    UNIQUE          NOT NULL ,
    bio         TEXT  ,
    image       VARCHAR  ,
    created_at  TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- CREATE INDEX IF NOT EXISTS users_username_idx ON users (username);
-- CREATE INDEX IF NOT EXISTS users_email_idx ON users (email);