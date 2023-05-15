CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TABLE users (
                       id TEXT PRIMARY KEY,
                       user_name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE todos (
                       id SERIAL PRIMARY KEY,
                       user_id TEXT REFERENCES users(id),
                       title VARCHAR(255) NOT NULL,
                       description TEXT,
                       completed BOOLEAN NOT NULL DEFAULT false,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX todos_user_id_idx ON todos(user_id);
