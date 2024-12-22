CREATE TABLE IF NOT EXISTS likes (
    id SERIAL PRIMARY KEY,
    echo_id INT NOT NULL,
    token_id TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (echo_id) REFERENCES echoes (id) ON DELETE CASCADE,
    CONSTRAINT unique_like UNIQUE (echo_id, token_id)
)