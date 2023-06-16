CREATE TABLE users (
    id            SERIAL         PRIMARY KEY,
    email         VARCHAR(100)   UNIQUE NOT NULL,
    username      VARCHAR(100)   UNIQUE NOT NULL,
    password      VARCHAR(100)   NOT NULL,
    created_at    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
    id            SERIAL         PRIMARY KEY,
    title         VARCHAR(100)   NOT NULL,
    author_id     INTEGER        NOT NULL,
    executor_id   INTEGER,
    description   TEXT           NOT NULL,
    available     BOOLEAN        NOT NULL DEFAULT TRUE,
    created_at    TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    code          TEXT           NOT NULL,
    results       TEXT,
    data          TEXT,
    -- img TEXT
    FOREIGN KEY (author_id) REFERENCES users(id),
    FOREIGN KEY (executor_id) REFERENCES users(id)
);
