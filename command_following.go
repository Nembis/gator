package main

import (
	"context"
	"fmt"

	"github.com/nembis/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, dbuser database.User) error {
	dbFeedFollow, err := s.db.GetFeedFollowsForUser(context.Background(), dbuser.Name)
	if err != nil {
		return fmt.Errorf("error: faileld to get feed follows: %v", err)
	}

	if len(dbFeedFollow) == 0 {
		fmt.Println("Currently not following any feeds")
		return nil
	}

	fmt.Printf("%v is following:\n", dbFeedFollow[0].UserName)
	for _, feedFollow := range dbFeedFollow {
		fmt.Printf("\t-%v\n", feedFollow.FeedName)
	}

	return nil
}
