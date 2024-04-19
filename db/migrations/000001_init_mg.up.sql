CREATE TABLE IF NOT EXISTS urls (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    short_url VARCHAR NOT NULL UNIQUE,
    url VARCHAR NOT NULL UNIQUE
);