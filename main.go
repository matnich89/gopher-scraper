package main

import (
	"email-scraper/internal/crawler"
	"email-scraper/internal/email"
	"email-scraper/pkg/logging"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	targetURL := flag.String("url", "", "Target URL to scan (required)")
	maxURLs := flag.Int("max", 100, "Maximum number of URLs to scan")
	flag.Parse()

	if *targetURL == "" {
		logging.LogError("Target URL is required")
		fmt.Println("\nUsage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	emailScraper := email.NewScraper()
	httpClient := &http.Client{Timeout: 10 * time.Second}
	c := crawler.NewCrawler(httpClient, *targetURL, *maxURLs, emailScraper)
	c.Crawl()
}
