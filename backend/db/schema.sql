PRAGMA foreign_keys = true;

DROP TABLE IF EXISTS logbook;

CREATE TABLE IF NOT EXISTS logbook (
    id INTEGER NOT NULL UNIQUE,
    created_at DATETIME,
    deleted_at DATETIME,
    modified_at DATETIME,
    name TEXT NOT NULL UNIQUE
);
