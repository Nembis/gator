package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nembis/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("error: agg command requires 1 argument <time_between_reqs>")
	}

	interval, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error: invalid time_between_reqs, example 1m, 1h")
	}

	scrapeFeeds(interval, s.db)

	return nil
}

func scrapeFeeds(interval time.Duration, db *database.Queries) {
	ticker := time.NewTicker(interval)
	ctx := context.Background()

	for {
		dbFeed, err := db.GetNextFeedToFetch(ctx)
		if err != nil {
			fmt.Println("error: failed to get the next feed to fetch: %v", err)
			break
		}

		err = db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
			LastFetchedAt: sql.NullTime{
				Time:  time.Now().UTC(),
				Valid: true,
			},
			ID: dbFeed.ID,
		})
		if err != nil {
			fmt.Printf("error: failed to mark %v feed as fetched: %v", dbFeed.Name, err)
			break
		}

		rssFeed, err := fetchFeed(ctx, dbFeed.Url)
		if err != nil {
			fmt.Printf("error: failed to fetch feed %v: %v", dbFeed.Name)
			<-ticker.C
			continue
		}

		timeNow := time.Now().UTC()
		for _, item := range rssFeed.Channel.Item {
			fmt.Println("ITEM: ", item.Title)
			id, err := uuid.NewV7()
			if err != nil {
				continue
			}
			pubTime, err := time.Parse(time.RFC1123Z, item.PubDate)
			if err != nil {
				pubTime = time.Now().UTC()
			}
			dbPost, err := db.CreatePost(ctx, database.CreatePostParams{
				ID:          id,
				CreatedAt:   timeNow,
				UpdatedAt:   timeNow,
				Title:       item.Title,
				Url:         dbFeed.Url,
				Description: item.Description,
				PublishedAt: pubTime,
				FeedID:      dbFeed.ID,
			})
			if err != nil {
				continue
			}
			fmt.Println("Added post: ", dbPost.Title)
		}

		<-ticker.C
	}
}
