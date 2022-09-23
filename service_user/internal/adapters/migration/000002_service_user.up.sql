CREATE TABLE sessions (
    session_id      UUID            PRIMARY KEY     NOT NULL ,
    username        VARCHAR(128)                    NOT NULL ,
    refresh_token   VARCHAR                         NOT NULL ,
    user_agent      VARCHAR(128)                    NOT NULL ,
    client_ip       VARCHAR(128)                    NOT NULL ,
    is_blocked      boolean                         NOT NULL DEFAULT false,
    expires_at      TIMESTAMP WITH TIME ZONE        NOT NULL ,
    created_at      TIMESTAMP WITH TIME ZONE        NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE        NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
