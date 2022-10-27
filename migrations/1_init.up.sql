CREATE TABLE livros
(
    id         SERIAL PRIMARY KEY,
    title       VARCHAR NOT NULL,
    description VARCHAR,
    value FLOAT NOT NULL
);