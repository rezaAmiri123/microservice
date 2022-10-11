CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE emails (
    email_id        UUID            PRIMARY KEY     NOT NULL DEFAULT uuid_generate_v4(),
    from_email      VARCHAR(128)                    NOT NULL ,
    to_email        TEXT []                         NOT NULL ,
    subject         VARCHAR(128)                    NOT NULL ,
    body            TEXT  ,
    created_at      TIMESTAMP WITH TIME ZONE        NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE        NOT NULL DEFAULT CURRENT_TIMESTAMP
);