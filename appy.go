package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
)

// TODO:
// - Limit to run once a day
// - Implement AllowedDomains
// - Add logic to check if job is new
// - Add logic to process multiple domains
// - Add data export

type jobPosting struct {
	companyName string
	jobTitle    string
	datePosted  time.Time
}

func getUrlList() ([]string, error) {
	file, err := os.Open("urls.txt")

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var urls []string
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	return urls, scanner.Err()
}

func main() {
	fmt.Println("Starting Appy Bot")
	urlList, err := getUrlList()
	if err != nil {
		fmt.Printf("readLines: %s\n>", err)
	}

	c := colly.NewCollector(
	// colly.AllowedDomains(urlList...),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML(".li", func(e *colly.HTMLElement) {})

	for _, url := range urlList {
		c.Visit(url)
	}

}
