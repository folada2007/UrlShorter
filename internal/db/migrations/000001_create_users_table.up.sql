CREATE TABLE shorter (
    id SERIAL PRIMARY KEY,
    long_url TEXT NOT NULL DEFAULT 'unknown',
    short_url TEXT NOT NULL DEFAULT 'unknown'
)

