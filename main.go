package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

func main() {
	url := "https://news.ycombinator.com/"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to fetch URL:", err)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Failed to parse HTML:", err)
	}

	file, err := os.Create("hackernews.csv")
	if err != nil {
		log.Fatal("Failed to create file:", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Title", "URL"})
	color.Cyan("Latest Hacker New Articles:")
	doc.Find(".titleline > a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		link, _ := s.Attr("href")
		writer.Write([]string{title, link})
		fmt.Printf("%d. %s\n", i+1, title)
	})
	color.Green("✅ Data saved to hackernews.csv")
}
