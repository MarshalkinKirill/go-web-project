CREATE TABLE IF NOT EXISTS users (
    device_id TEXT PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS feeds (
    feed_id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    device_id TEXT NOT NULL,
    FOREIGN KEY(device_id) REFERENCES users(device_id)
);
