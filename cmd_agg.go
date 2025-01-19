package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ankylosaurus11/blog_agg/internal/database"
)

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	nextFeed, err := s.db.NextFeedFetch(ctx)
	if err != nil {
		return err
	}
	feedID := nextFeed.ID

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
		fmt.Println(rssFeed.Channel.Item[i].Title)
	}

	return nil
}
