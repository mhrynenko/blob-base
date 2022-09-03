-- +migrate Up

CREATE TABLE IF NOT EXISTS blobs (
    id bigserial primary key,
    attributes jsonb not null
);

-- +migrate Down

DROP TABLE IF EXISTS  blobs;
