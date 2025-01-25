package main

import (
	"context"
	"fmt"
	"time"

	gatorconfig "github.com/ankylosaurus11/blog_agg/internal/config"
	"github.com/ankylosaurus11/blog_agg/internal/database"
)

func browse(s *state, _ command, limit int32) error {
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

	posts, err := s.db.GetPostsForUser(ctx, database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("error fetching posts: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found!")
		return nil
	}

	for _, post := range posts {
		fmt.Printf("\nTitle: %s\n", post.Title)
		fmt.Printf("URL: %s\n", post.Url)
		if post.Description.Valid {
			fmt.Printf("Description: %s\n", post.Description.String)
		}
		if post.PublishedAt.Valid {
			fmt.Printf("Published: %s\n", post.PublishedAt.Time.Format(time.RFC822))
		}
		fmt.Println("___________________")
	}

	return nil
}
