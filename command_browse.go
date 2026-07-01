package main

import (
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(s *state, cmd command) error {
	limit := 2
	if len(cmd.args) == 1 {
		if cmdLimit, err := strconv.ParseInt(cmd.args[0], 10, 64); err == nil {
			limit = int(cmdLimit)
		}
	}

	dbPosts, err := s.db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return fmt.Errorf("error: failed to get posts: %v", err)
	}

	for _, post := range dbPosts {
		fmt.Println("Title:", post.Title)
		fmt.Println("Published on:", post.PublishedAt)
		fmt.Println("URL on:", post.Url)
		fmt.Println()
	}

	return nil
}
