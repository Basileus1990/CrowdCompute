-- This file contains SQL queries for creating tables in the database
-- It is also schema for the database

CREATE TABLE users (
    username      VARCHAR(100)   NOT NULL,
    email         VARCHAR(100)   UNIQUE NOT NULL,
    password      VARCHAR(100)   NOT NULL,
    auth_token    VARCHAR(200),
    created_at    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (username)
);

CREATE TABLE task_info (
    title               VARCHAR(100)   NOT NULL,
    author_username     VARCHAR(100)   NOT NULL,
    description         TEXT           NOT NULL,
    code                TEXT           NOT NULL,
    created_at          TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- img TEXT
    PRIMARY KEY (title),

    FOREIGN KEY (author_username) REFERENCES users(username) ON DELETE CASCADE
);

CREATE TABLE task (
    id                  SERIAL         NOT NULL,
    task_info_title     VARCHAR(100)   NOT NULL,
    data                TEXT           NOT NULL, -- JSON

    PRIMARY KEY (id),

    FOREIGN KEY (task_info_title) REFERENCES task_info(title) ON DELETE CASCADE
);

CREATE TABLE task_result (
    id                  SERIAL         NOT NULL,
    task_id             INTEGER        NOT NULL,
    executor_username   VARCHAR(100),
    results             TEXT, -- JSON
    created_at          TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (task_id) REFERENCES task(id) ON DELETE CASCADE,
    FOREIGN KEY (executor_username) REFERENCES users(username) ON DELETE SET NULL
);