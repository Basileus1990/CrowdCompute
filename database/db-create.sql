-- This file contains SQL queries for creating tables in the database
-- It is also schema for the database

CREATE TABLE users (
    username      TEXT   NOT NULL,
    email         TEXT   UNIQUE NOT NULL,
    password      TEXT   NOT NULL, -- hashed
    created_at    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (username)
);

CREATE TABLE task_info (
    title               TEXT           NOT NULL,
    author_username     TEXT           NOT NULL,
    description         TEXT           NOT NULL,
    code                TEXT           NOT NULL,
    created_at          TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- img TEXT
    PRIMARY KEY (title),

    FOREIGN KEY (author_username) REFERENCES users(username) ON DELETE CASCADE
);

CREATE TABLE task (
    id                  SERIAL         NOT NULL,
    task_info_title     TEXT           NOT NULL,
    data                TEXT           NOT NULL, -- JSON

    PRIMARY KEY (id),

    FOREIGN KEY (task_info_title) REFERENCES task_info(title) ON DELETE CASCADE
);

CREATE TABLE task_result (
    id                  SERIAL         NOT NULL,
    task_id             INTEGER        NOT NULL,
    executor_username   TEXT,
    results             TEXT, -- JSON
    created_at          TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (task_id) REFERENCES task(id) ON DELETE CASCADE,
    FOREIGN KEY (executor_username) REFERENCES users(username) ON DELETE SET NULL
);