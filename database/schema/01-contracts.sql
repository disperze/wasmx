CREATE TABLE contracts
(
    address       TEXT                        NOT NULL UNIQUE PRIMARY KEY,
    code_id       BIGINT                      NOT NULL,
    creator       TEXT                        NOT NULL DEFAULT '',
    admin         TEXT                        NOT NULL DEFAULT '',
    label         TEXT                        NOT NULL DEFAULT '',
    creation_time TEXT                        NOT NULL DEFAULT '',
    height        BIGINT                      NOT NULL
);
