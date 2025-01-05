package main

import (
	"context"

	gatorconfig "github.com/ankylosaurus11/blog_agg/internal/config"
	"github.com/ankylosaurus11/blog_agg/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		jsonData, err := gatorconfig.Read()
		if err != nil {
			return err
		}

		loggedInUser := jsonData.CurrentUserName

		ctx := context.Background()
		currentUser, err := s.db.GetUser(ctx, loggedInUser)
		if err != nil {
			return err
		}

		return handler(s, cmd, currentUser)
	}
}
