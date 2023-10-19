package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"os"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("hypeauditor.com"),
	)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code", r.StatusCode)
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error", err.Error())
	})

	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"Nickname", "Full Name", "Rating", "Followers", "Country", "Engagement Authenticity", "Engagement Average"}
	err = w.Write(header)
	if err != nil {
		panic(err)
	}

	c.OnHTML(".row", func(e *colly.HTMLElement) {
		div := e.DOM
		nickname := div.Find(".contributor__name-content").Text()
		fullName := div.Find(".contributor__title").Text()
		rating := div.Find(".ml-2").Text()
		followers := div.Find(".row-cell.subscribers").Text()
		country := div.Find(".row-cell.audience").Text()
		engAuth := div.Find("row-cell.authentic").Text()
		engAvg := div.Find("row-cell.engagement").Text()
		data := []string{nickname, fullName, rating, followers, country, engAuth, engAvg}
		err = w.Write(data)
		if err != nil {
			panic(err)
		}
	})

	c.Visit("https://hypeauditor.com/top-instagram-all-russia/")
}
