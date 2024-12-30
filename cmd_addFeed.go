package main

import (
	"context"
	"fmt"
	"time"

	gatorconfig "github.com/ankylosaurus11/blog_agg/internal/config"
	"github.com/ankylosaurus11/blog_agg/internal/database"
	"github.com/google/uuid"
)

func addFeed(s *state, _ command, name string, url string) error {
	jsonData, err := gatorconfig.Read()
	if err != nil {
		return err
	}

	currentUser := jsonData.CurrentUserName

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, currentUser)
	if err != nil {
		return err
	}

	feedID := uuid.New()
	now := time.Now()

	newFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        feedID,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeed(ctx, url)
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        feedID,
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Created feed: ", newFeed.Name, newFeed.Url)
	return nil
}
