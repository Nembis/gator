package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nembis/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, dbUser database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("error: addfeed requires 2 arguments <name> <url>")
	}

	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("error: command requires login")
	}
	ctx := context.Background()

	id, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("error: failed to generate new id for feed: %v", err)
	}

	timeNow := time.Now().UTC()

	dbFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        id,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    dbUser.ID,
	})
	if err != nil {
		return fmt.Errorf("error: failed to create new feed: %v", err)
	}

	id, err = uuid.NewV7()
	if err != nil {
		return fmt.Errorf("error: failed to generate new id for feed: %v", err)
	}

	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        id,
		CreatedAt: timeNow,
		UserID:    dbFeed.UserID,
		FeedID:    dbFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("error: failed to create feed follow: %v", err)
	}

	fmt.Println("Successfully created feed")
	fmt.Println("Feed:", dbFeed)

	return nil
}
