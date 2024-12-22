CREATE TABLE IF NOT EXISTS echoes (
    id SERIAL PRIMARY KEY,
    echo TEXT NOT NULL,
    author TEXT DEFAULT 'Anonymous',
    token_id TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)