package models

import (
	"database/sql"
	"time"

	"github.com/MrAinslay/fiber-rss-feed/packages/config"
	"github.com/google/uuid"
)

type PostLike struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uuid.UUID `json:"user_id"`
	PostId    uuid.UUID `json:"post_id"`
}

type FeedFollow struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
}

type Post struct {
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	PublishedAt string    `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
}

type Feed struct {
	Id            uuid.UUID    `json:"id"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	UserId        uuid.UUID    `json:"user_id"`
	Name          string       `json:"name"`
	Url           string       `json:"url"`
	LastFetchedAt sql.NullTime `json:"last_fetched_at"`
}

func DatabasePostLikeToPostLike(postLike config.PostLike) PostLike {
	return PostLike{
		Id:        postLike.ID,
		CreatedAt: postLike.CreatedAt,
		UpdatedAt: postLike.UpdatedAt,
		UserId:    postLike.UserID,
		PostId:    postLike.PostID,
	}
}

func DatabasePostLikesToPostLikes(postLikes []config.PostLike) []PostLike {
	result := make([]PostLike, len(postLikes))
	for index, postLike := range postLikes {
		result[index] = DatabasePostLikeToPostLike(postLike)
	}
	return result
}

func DatabaseFeedFollowToFeedFollow(feedFollow config.FeedFollow) FeedFollow {
	return FeedFollow{
		Id:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserId:    feedFollow.UserID,
		FeedId:    feedFollow.FeedID,
	}
}

func DatabaseFeedFollowsToFeedFollows(feedFollows []config.FeedFollow) []FeedFollow {
	result := make([]FeedFollow, len(feedFollows))
	for index, feedFollow := range feedFollows {
		result[index] = DatabaseFeedFollowToFeedFollow(feedFollow)
	}
	return result
}

func DatabaseFeedToFeed(feed config.Feed) Feed {
	return Feed{
		Id:            feed.ID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		UserId:        feed.UserID,
		Name:          feed.Name,
		Url:           feed.Url,
		LastFetchedAt: feed.LastFetchedAt,
	}
}

func DatabaseFeedsToFeeds(feeds []config.Feed) []Feed {
	result := make([]Feed, len(feeds))
	for index, feed := range feeds {
		result[index] = DatabaseFeedToFeed(feed)
	}
	return result
}

func DatabasePostToPost(post config.Post) Post {
	return Post{
		Id:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: post.Description,
		PublishedAt: post.PublishedAt,
		FeedID:      post.FeedID,
	}
}

func DatabasePostsToPosts(posts []config.Post) []Post {
	result := make([]Post, len(posts))
	for index, post := range posts {
		result[index] = DatabasePostToPost(post)
	}
	return result
}
