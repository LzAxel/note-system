CREATE TABLE ACCOUNT(
    ID SERIAL PRIMARY KEY NOT NULL,
    USERNAME VARCHAR(50) UNIQUE NOT NULL,
    PASSWORD_HASH VARCHAR(255) NOT NULL,
    HASH_SALT VARCHAR(32) NOT NULL,
    CREATED_AT timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE NOTE(
    ID SERIAL PRIMARY KEY NOT NULL,
    NAME VARCHAR(100) NOT NULL,
    TEXT VARCHAR(1000) NOT NULL,
    TAG VARCHAR(255),
    URL VARCHAR(100) UNIQUE NOT NULL,
    IS_PUBLIC BOOLEAN DEFAULT FALSE,
    CREATED_AT timestamptz NOT NULL DEFAULT now(),
    UPDATED_AT timestamptz NOT NULL DEFAULT now(),
    ACCOUNT_ID INT REFERENCES ACCOUNT (ID) ON DELETE CASCADE
);
