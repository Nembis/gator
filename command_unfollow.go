package main

import (
	"context"
	"fmt"

	"github.com/nembis/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, dbUser database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("error: unfollow requires 1 arguemnt url")
	}
	ctx := context.Background()
	url := cmd.args[0]

	dbFeed, err := s.db.GetFeed(ctx, url)
	if err != nil {
		return fmt.Errorf("error: failed to get feed from database: %v", err)
	}

	err = s.db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		FeedID: dbFeed.ID,
		UserID: dbUser.ID,
	})
	if err != nil {
		return fmt.Errorf("error: failed to delete feed: %v", err)
	}

	fmt.Printf("Successfully unfollowed %s feed", dbFeed.Name)

	return nil
}
