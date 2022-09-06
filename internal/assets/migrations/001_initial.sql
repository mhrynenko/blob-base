-- +migrate Up

CREATE TABLE IF NOT EXISTS blobs (
    id bigserial primary key,
    owner text not null,
    attributes jsonb not null
);

-- +migrate Down

DROP TABLE IF EXISTS  blobs;
