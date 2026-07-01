package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nembis/gator/internal/database"
)

func handlerFollow(s *state, cmd command, dbUser database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("error: follow command requires 1 argument, url")
	}
	url := cmd.args[0]
	ctx := context.Background()

	dbFeed, err := s.db.GetFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("error: failed to get feed with provided url: %v", err)
	}

	id, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("error: failed to generate uuid: %v", err)
	}

	dbFeedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        id,
		CreatedAt: time.Now().UTC(),
		UserID:    dbUser.ID,
		FeedID:    dbFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("error: failed to create feed_follow: %v", err)
	}

	fmt.Printf("Successfully set %v to follow %v", dbFeedFollow.UserName, dbFeedFollow.FeedName)

	return nil
}
