package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// SitemapIndex is this
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

// News is this
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	resp.Body.Close()
}
