-- +goose Down
DROP TABLE IF EXISTS stores_cache;
DROP TABLE IF EXISTS products_cache;

DROP TABLE IF EXISTS sagas;
DROP TABLE IF EXISTS outbox;
DROP TABLE IF EXISTS inbox;
DROP TABLE IF EXISTS snapshots;
DROP TABLE IF EXISTS events;
