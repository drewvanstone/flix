CREATE TYPE STATUS AS ENUM ('open', 'done');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    task_id BIGINT NOT NULL
);

CREATE TABLE tasks (
       id SERIAL PRIMARY KEY,
       description TEXT NOT NULL,
       status STATUS NOT NULL DEFAULT 'open'
);
