CREATE TABLE baskets
(
    id         text        NOT NULL,
    user_id    text,
    payment_id text,
    status     text,
    items      text,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);

CREATE TRIGGER created_at_baskets_trgr
    BEFORE UPDATE
    ON baskets
    FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
CREATE TRIGGER updated_at_baskets_trgr
    BEFORE UPDATE
    ON baskets
    FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();
