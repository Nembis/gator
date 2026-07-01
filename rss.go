package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)

	if err != nil {
		return nil, fmt.Errorf("error: failed to create fetch request: %v", err)
	}
	request.Header.Set("User-Agent", "gator")

	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error: failed to fetch rss fed: %vla", err)
	}

	dataByte, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error: failed to read response body to bytes: %v", err)
	}

	rssFeed := &RSSFeed{}
	if err = xml.Unmarshal(dataByte, rssFeed); err != nil {
		return nil, fmt.Errorf("error: failed to unmarshal to xml: %v", err)
	}

	rssFeed.Channel.Title = html.EscapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.EscapeString(rssFeed.Channel.Description)

	for i := 0; i < len(rssFeed.Channel.Item); i++ {
		item := rssFeed.Channel.Item[i]
		item.Title = html.EscapeString(item.Title)
		item.Description = html.EscapeString(item.Description)
	}

	return rssFeed, nil
}
