-- +goose Up
CREATE TABLE posts(
id UUID PRIMARY KEY,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL,
title TEXT NOT NULL,
url TEXT UNIQUE NOT NULL,
description TEXT NOT NULL,
published_at TEXT NOT NULL,
feed_id UUID NOT NULL REFERENCES feeds
ON DELETE CASCADE,
CONSTRAINT fk_feeds
FOREIGN KEY(feed_id)
REFERENCES feeds(id)
);

-- +goose Down
DROP TABLE posts;