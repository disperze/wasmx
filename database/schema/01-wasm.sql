CREATE TABLE codes
(
    code_id       BIGINT                      NOT NULL UNIQUE PRIMARY KEY,
    creator       TEXT                        NOT NULL DEFAULT '',
    creation_time TEXT                        NOT NULL DEFAULT '',
    version       TEXT                        NULL,
    height        BIGINT                      NOT NULL
);

CREATE INDEX codes_creator_index ON codes (creator);

CREATE TABLE contracts
(
    address       TEXT                        NOT NULL UNIQUE PRIMARY KEY,
    code_id       BIGINT                      NOT NULL,
    creator       TEXT                        NOT NULL DEFAULT '',
    admin         TEXT                        NOT NULL DEFAULT '',
    label         TEXT                        NOT NULL DEFAULT '',
    creation_time TEXT                        NOT NULL DEFAULT '',
    tx            BIGINT                      NOT NULL DEFAULT 0,
    gas           BIGINT                      NOT NULL DEFAULT 0,
    fees          BIGINT                      NOT NULL DEFAULT 0,
    height        BIGINT                      NOT NULL,
    ibc           BOOLEAN                     NOT NULL
);

CREATE INDEX contracts_code_id_index ON contracts (code_id);
CREATE INDEX contracts_creator_index ON contracts (creator);

CREATE TABLE tokens
(
    address      TEXT                        NOT NULL UNIQUE PRIMARY KEY,
    name         TEXT                        NOT NULL DEFAULT '',
    symbol       TEXT                        NOT NULL DEFAULT '',
    decimals     TINYINT                     NOT NULL DEFAULT 0,
    supply       BIGINT                      NOT NULL DEFAULT 0
);
