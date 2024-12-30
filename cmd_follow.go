package main

import (
	"context"
	"fmt"
	"time"

	gatorconfig "github.com/ankylosaurus11/blog_agg/internal/config"
	"github.com/ankylosaurus11/blog_agg/internal/database"
	"github.com/google/uuid"
)

func follow(s *state, _ command, url string) error {
	ctx := context.Background()
	jsonData, err := gatorconfig.Read()
	if err != nil {
		return err
	}

	currentUser := jsonData.CurrentUserName

	user, err := s.db.GetUser(ctx, currentUser)
	if err != nil {
		return err
	}
	feedID := uuid.New()
	now := time.Now()
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

	fmt.Println(feed.Name)
	fmt.Println(user.Name)
	return nil
}

func following(s *state, _ command) error {
	ctx := context.Background()
	jsonData, err := gatorconfig.Read()
	if err != nil {
		return err
	}

	currentUser := jsonData.CurrentUserName

	user, err := s.db.GetUser(ctx, currentUser)
	if err != nil {
		return err
	}
	feeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}
