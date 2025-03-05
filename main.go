package main

import (
	"fmt"
	"log"
	"net/http"

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

	doc,err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Failed to parse HTML:", err)
	}

	color.Cyan("Latest Hacker New Articles:");
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
         title:= s.Text()
		 fmt.Printf("%d. %s\n", i+1, title)
	})
}
