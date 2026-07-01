package main

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd command) error {
	dbFeeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error: failed to get feeds: %v", err)
	}

	for _, record := range dbFeeds {
		fmt.Println("Feed Name: ", record.Feed.Name)
		fmt.Println("\tURL: ", record.Feed.Url)
		fmt.Println("\tUsername: ", record.User.Name)
	}

	return nil
}
