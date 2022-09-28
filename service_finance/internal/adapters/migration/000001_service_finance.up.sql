CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE accounts (
    account_id  UUID  PRIMARY KEY           NOT NULL    DEFAULT uuid_generate_v4(),
    owner_id    UUID                        NOT NULL ,
    balance     BIGINT                      NOT NULL ,
    currency    VARCHAR                     NOT NULL ,
    created_at TIMESTAMPTZ                  NOT NULL    DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE     NOT NULL    DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE entries (
    entrie_id   UUID        PRIMARY KEY     NOT NULL    DEFAULT uuid_generate_v4(),
    account_id  UUID                        NOT NULL ,
    amount      BIGINT                      NOT NULL ,
    created_at  TIMESTAMPTZ                 NOT NULL    DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transfers (
    transfer_id     UUID  PRIMARY KEY           NOT NULL DEFAULT uuid_generate_v4(),
    from_account_id UUID                        NOT NULL ,
    to_account_id   UUID                        NOT NULL ,
    amount          BIGINT                      NOT NULL ,
    created_at      TIMESTAMPTZ                 NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP WITH TIME ZONE    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE entries ADD FOREIGN KEY (account_id) REFERENCES accounts(account_id);
ALTER TABLE transfers ADD FOREIGN KEY (from_account_id) REFERENCES accounts(account_id);
ALTER TABLE transfers ADD FOREIGN KEY (to_account_id) REFERENCES accounts(account_id);

CREATE INDEX IF NOT EXISTS accounts_owner_idx ON accounts (owner_id);
CREATE INDEX IF NOT EXISTS entries_account_id_idx ON entries(account_id);
CREATE INDEX IF NOT EXISTS transfers_from_account_id_idx ON transfers(from_account_id);
CREATE INDEX IF NOT EXISTS transfers_to_account_id_idx ON transfers(to_account_id);
CREATE INDEX IF NOT EXISTS transfers_from_account_id_to_account_id_idx ON transfers(from_account_id, to_account_id);
