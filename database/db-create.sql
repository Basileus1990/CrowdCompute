CREATE TABLE users (
    id            SERIAL         PRIMARY KEY,
    email         VARCHAR(100)   UNIQUE NOT NULL,
    username      VARCHAR(100)   UNIQUE NOT NULL,
    password      VARCHAR(100)   NOT NULL,
    created_at    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE task_info (
    id            SERIAL         NOT NULL,
    title         VARCHAR(100)   NOT NULL UNIQUE,
    author_id     INTEGER        NOT NULL,
    description   TEXT           NOT NULL,
    code          TEXT           NOT NULL,
    created_at    TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- img TEXT
    PRIMARY KEY (id),

    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE task (
    id            SERIAL         NOT NULL,
    task_info_id  INTEGER        NOT NULL,
    data          TEXT           NOT NULL, -- JSON

    PRIMARY KEY (id),

    FOREIGN KEY (task_info_id) REFERENCES task_info(id) ON DELETE CASCADE
);

CREATE TABLE task_result (
    id            SERIAL         NOT NULL,
    task_id       INTEGER        NOT NULL,
    executor_id   INTEGER,
    results       TEXT,                    -- JSON
    created_at    TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (task_id) REFERENCES task(id) ON DELETE CASCADE,
    FOREIGN KEY (executor_id) REFERENCES users(id)
);