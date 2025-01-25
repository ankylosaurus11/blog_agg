package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/ankylosaurus11/blog_agg/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	nextFeed, err := s.db.NextFeedFetch(ctx)
	if err != nil {
		return err
	}
	feedID := nextFeed.ID
	now := time.Now()

	rssFeed, err := fetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return err
	}

	s.db.MarkFetched(ctx, database.MarkFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: feedID,
	})

	for i := range rssFeed.Channel.Item {
		postID := uuid.New()

		publishedAt, err := time.Parse(time.RFC1123Z, rssFeed.Channel.Item[i].PubDate)
		if err != nil {
			log.Printf("error parsing date for posts %s: %v", rssFeed.Channel.Item[i].Title, err)
			continue
		}

		_, err = s.db.CreatePost(ctx, database.CreatePostParams{
			ID:        postID,
			CreatedAt: now,
			UpdatedAt: now,
			Title:     rssFeed.Channel.Item[i].Title,
			Url:       rssFeed.Channel.Item[i].Link,
			Description: sql.NullString{
				String: rssFeed.Channel.Item[i].Description,
				Valid:  rssFeed.Channel.Item[i].Description != "",
			},
			PublishedAt: sql.NullTime{
				Time:  publishedAt,
				Valid: true,
			},
			FeedID: feedID,
		})
		if err != nil {
			log.Printf("error creating post: %v", err)
			continue
		}
	}

	return nil
}
