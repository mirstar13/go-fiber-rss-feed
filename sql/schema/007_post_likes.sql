-- +goose Up
CREATE TABLE post_likes(
id UUID PRIMARY KEY,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL,
user_id UUID NOT NULL REFERENCES users
ON DELETE CASCADE,
post_id UUID NOT NULL REFERENCES posts
ON DELETE CASCADE,
CONSTRAINT fk_users
FOREIGN KEY(user_id)
REFERENCES users(id),
CONSTRAINT fk_posts
FOREIGN KEY(post_id)
REFERENCES posts(id)
);

-- +goose Down
DROP TABLE post_likes;