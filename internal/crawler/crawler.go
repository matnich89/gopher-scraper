package crawler

import (
	"email-scraper/internal/email"
	"email-scraper/pkg/logging"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
)

type Crawler struct {
	httpClient   *http.Client
	urlsToVisit  []string
	scrapedURLs  map[string]bool
	count        int
	maxURLS      int
	emailScraper *email.Scraper
}

func NewCrawler(httpClient *http.Client, targetUrl string, maxURLS int, emailScraper *email.Scraper) *Crawler {
	return &Crawler{
		httpClient:   httpClient,
		urlsToVisit:  []string{targetUrl},
		scrapedURLs:  make(map[string]bool),
		count:        0,
		maxURLS:      maxURLS,
		emailScraper: emailScraper,
	}
}

func (c *Crawler) Crawl() {
	for len(c.urlsToVisit) > 0 && c.count < c.maxURLS {
		currentURL := c.urlsToVisit[0]
		c.urlsToVisit = c.urlsToVisit[1:]

		if c.scrapedURLs[currentURL] {
			continue
		}

		c.count++
		c.scrapedURLs[currentURL] = true

		parsedURL, err := url.Parse(currentURL)

		if err != nil {
			logging.LogWarning("Failed to parse URL %s: %v", currentURL, err)
			continue
		}
		baseURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)

		resp, err := c.httpClient.Get(currentURL)
		if err != nil {
			logging.LogWarning("HTTP request failed for %s: %v", currentURL, err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			logging.LogError("received none 200 response cannot scrape site")
			continue
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			logging.LogWarning("Failed to parse HTML from %s: %v", currentURL, err)
			continue
		}

		c.emailScraper.FindEmails(doc)

		c.scrapeLinks(baseURL, doc)

	}
}

func (c *Crawler) scrapeLinks(baseURL string, doc *goquery.Document) {
	doc.Find("a").Each(func(_ int, anchor *goquery.Selection) {
		link, exists := anchor.Attr("href")
		if !exists {
			return
		}

		// Handle relative URLs
		if strings.HasPrefix(link, "/") {
			link = baseURL + link
		} else if !strings.HasPrefix(link, "http") && !strings.HasPrefix(link, "//") {
			link = baseURL + "/" + strings.TrimPrefix(link, "./")
		}

		// Add new URLs to queue
		if !c.scrapedURLs[link] && strings.HasPrefix(link, "http") {
			c.urlsToVisit = append(c.urlsToVisit, link)
		}
	})
}
