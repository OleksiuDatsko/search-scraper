-- +goose Up
CREATE TABLE
    blacklist (
        id INTEGER PRIMARY KEY,
        domain VARCHAR(255) NOT NULL,
        url VARCHAR(255) NOT NULL,
        filter_type VARCHAR(5) NOT NULL -- url or domain
    );

CREATE TABLE
    whitelist (
        id INTEGER PRIMARY KEY,
        domain VARCHAR(255) NOT NULL,
        url VARCHAR(255) NOT NULL,
        filter_type VARCHAR(5) NOT NULL -- url or domain
    );

CREATE TABLE
    findedlist (
        id INTEGER PRIMARY KEY,
        domain VARCHAR(255) NOT NULL,
        url VARCHAR(255) NOT NULL,
        filter_type VARCHAR(5) NOT NULL -- url or domain
    );

-- +goose Down
DROP TABLE blacklist;

DROP TABLE whitelist;

DROP TABLE findedlist;