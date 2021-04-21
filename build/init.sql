CREATE TABLE fibonacci (
    id serial primary key,
    ordinal varchar unique not null,
    fibonacci varchar not null
);

INSERT INTO fibonacci (ordinal,fibonacci) VALUES (0,0), (1,1);

CREATE INDEX ordinal_idx on fibonacci(ordinal);