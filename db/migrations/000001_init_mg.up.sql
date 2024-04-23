CREATE TABLE IF NOT EXISTS urls (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    short_url VARCHAR(8) NOT NULL UNIQUE,
    url VARCHAR(2048) NOT NULL UNIQUE
);