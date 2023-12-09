package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/santosh1608/project-rss/dataConnector"
	"github.com/santosh1608/project-rss/models"
)

func startScraping(d time.Duration) {
	for ; ; <-time.Tick(d) {
		data, err := dataConnector.GetAllFeeds()
		if err != nil {
			fmt.Println("Error found")
		}
		wg := sync.WaitGroup{}
		for _, d := range data {
			wg.Add(1)
			fmt.Println(d.Url)
			go fetcher(&wg, d)
		}

		wg.Wait()
		// loop to each feed and get details of posts

		// Create post in db think of db design
	}
}

func fetcher(wg *sync.WaitGroup, feed *models.Feed) {
	defer wg.Done()
	resp, err := http.Get(feed.Url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}
	rssFeed := RSSFeed{}
	fmt.Println(string(data))
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		fmt.Println("error", err)
	}
	for _, item := range rssFeed.Channel.Item {
		log.Println("Found post", item.Title, "on feed", feed.Name)
		dataConnector.CreatePost(&models.Post{Title: item.Title, FeedId: feed.Id})
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
