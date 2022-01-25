CREATE TABLE codes
(
    code_id       BIGINT                      NOT NULL UNIQUE PRIMARY KEY,
    creator       TEXT                        NOT NULL DEFAULT '',
    creation_time TEXT                        NOT NULL DEFAULT '',
    hash          TEXT                        NOT NULL DEFAULT '',
    size          BIGINT                      NOT NULL,
    version       TEXT                        NULL,
    height        BIGINT                      NOT NULL,
    ibc           BOOLEAN                     NULL,
    cw20          BOOLEAN                     NULL,
    verified      BOOLEAN                     NOT NULL DEFAULT FALSE
);

CREATE INDEX codes_creator_index ON codes (creator);
CREATE INDEX codes_ibc_index ON codes (ibc);
CREATE INDEX codes_cw20_index ON codes (cw20);

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
    height        BIGINT                      NOT NULL
);

ALTER TABLE contracts ADD CONSTRAINT fk_code FOREIGN KEY (code_id) REFERENCES codes (code_id);

CREATE INDEX contracts_code_id_index ON contracts (code_id);
CREATE INDEX contracts_creator_index ON contracts (creator);

CREATE TABLE tokens
(
    address      TEXT                        NOT NULL UNIQUE PRIMARY KEY,
    name         TEXT                        NOT NULL DEFAULT '',
    symbol       TEXT                        NOT NULL DEFAULT '',
    decimals     SMALLINT                    NOT NULL DEFAULT 0,
    supply       TEXT                        NOT NULL DEFAULT 0
);

ALTER TABLE tokens ADD CONSTRAINT fk_contract FOREIGN KEY (address) REFERENCES contracts (address);
