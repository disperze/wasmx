CREATE TABLE codes
(
    code_id       BIGINT                      NOT NULL UNIQUE PRIMARY KEY,
    source        TEXT                        ,  
    builder       TEXT                        ,  
    creator       TEXT                        NOT NULL DEFAULT '',
    creation_time TEXT                        NOT NULL DEFAULT '',
    height        BIGINT                      NOT NULL
);

CREATE INDEX codes_creator_index ON codes (creator);
