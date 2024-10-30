// Copyright © 2024 Jakub Kapusta <jakub-dev1@protonmail.com>
package dbupdate

const (
	prepDb = `PRAGMA encoding = 'UTF-8';
PRAGMA foreign_keys = TRUE;
CREATE TABLE IF NOT EXISTS
types (
    id INTEGER NOT NULL PRIMARY KEY,
    type TEXT NOT NULL
);
INSERT OR IGNORE INTO types (id, type)
VALUES
    (0, 'unknown'),
    (1, 'directory'),
    (2, 'regular_file'),
    (3, 'block_device_file'),
    (4, 'character_device_file'),
    (5, 'named_pipe'),
    (6, 'symbolic_link'),
    (7, 'socket');
CREATE TABLE IF NOT EXISTS
paths (
    path TEXT NOT NULL PRIMARY KEY,
    type_id INTEGER DEFAULT 0,
    FOREIGN KEY(type_id) REFERENCES types(id)
);`

	preparedInsert = `INSERT OR IGNORE INTO
    paths (path,type_id)
    VALUES (?,?);`
)
