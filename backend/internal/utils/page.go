package utils

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func FetchPageSummary(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "Unable to fetch content."
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "Failed to parse page."
	}

	title := doc.Find("title").Text()
	desc, _ := doc.Find("meta[name=description]").Attr("content")

	return fmt.Sprintf("Title: %s\nDescription: %s", title, desc)
}
