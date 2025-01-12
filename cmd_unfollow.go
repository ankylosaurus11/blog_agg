package main

import (
	"context"
	"fmt"

	gatorconfig "github.com/ankylosaurus11/blog_agg/internal/config"
	"github.com/ankylosaurus11/blog_agg/internal/database"
)

func unfollow(s *state, _ command, url string) error {
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
	err = s.db.Unfollow(ctx, database.UnfollowParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return err
	}

	fmt.Printf("User: %v, unfollowed feed at url: %v", currentUser, url)
	return nil
}
