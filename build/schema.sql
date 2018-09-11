CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE movies (
       id SERIAL PRIMARY KEY,
       title TEXT NOT NULL,
       description TEXT NOT NULL
);
