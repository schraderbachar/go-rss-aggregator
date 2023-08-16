// get am rss link and add it to the database

package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RssFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Langauge    string    `xml:"language"`
		Item        []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func URLToFeed(url string) (RssFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second, //if it takes longer than 10 seconds to load, don't get it
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RssFeed{}, err
	}
	defer resp.Body.Close()

	//get all data from res body
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RssFeed{}, err
	}
	//we want to read the dat slice of Bytes into the Rss feed strcut
	rssFeed := RssFeed{}

	err = xml.Unmarshal(dat, &rssFeed) //pass in data and where we want to unmarshal the data
	if err != nil {
		return RssFeed{}, err
	}
	return rssFeed, nil //if no errors, return the new rss feed
}
