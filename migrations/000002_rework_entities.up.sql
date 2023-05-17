DROP TABLE IF EXISTS feeds;
DROP TABLE IF EXISTS users;


CREATE TABLE IF NOT EXISTS users (
    id       SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    CONSTRAINT username_users_uniq_constr UNIQUE(username)
);

CREATE TABLE IF NOT EXISTS cache (
    id          SERIAL PRIMARY KEY,
    url         TEXT NOT NULL,
    last_update TIMESTAMP NOT NULL ,
    CONSTRAINT url_cache_uniq_constr UNIQUE(url)
);

CREATE TABLE IF NOT EXISTS feeds (
    id       SERIAL PRIMARY KEY,
    title    TEXT,
    user_id  INTEGER REFERENCES users(id) ON DELETE CASCADE,
    cache_id INTEGER REFERENCES cache(id) ON DELETE CASCADE,
    CONSTRAINT user_cache_uniq_constr UNIQUE(user_id, cache_id)
);

CREATE TABLE IF NOT EXISTS news_items (
    id          SERIAL PRIMARY KEY,
    title       TEXT NOT NULL,
    description TEXT,
    url         TEXT NOT NULL,
    image_url   TEXT,
    date        TIMESTAMP WITH TIME ZONE,
    hash        TEXT NOT NULL,
    CONSTRAINT news_hash_uniq_constr UNIQUE(hash)
);

CREATE TABLE IF NOT EXISTS cache_item (
    cache_id INTEGER REFERENCES cache(id) ON DELETE CASCADE,
    item_id  INTEGER REFERENCES news_items(id) ON DELETE CASCADE,
    PRIMARY KEY(cache_id, item_id)
);
