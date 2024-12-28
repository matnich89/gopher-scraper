package email

import (
	"email-scraper/pkg/logging"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

type Scraper struct {
	emails     map[string]bool
	emailRegex *regexp.Regexp
}

func NewScraper() *Scraper {
	return &Scraper{emails: make(map[string]bool),
		emailRegex: regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)}
}

func (s *Scraper) FindEmails(doc *goquery.Document) {
	pageText := doc.Text()
	foundEmails := s.emailRegex.FindAllString(pageText, -1)
	newEmails := 0

	for _, email := range foundEmails {
		if !s.emails[email] {
			s.emails[email] = true
			newEmails++
			logging.LogSuccess("New email found: %s", email)
		}
	}
}

func (s *Scraper) Results() {
	var output string
	for email := range s.emails {
		output += email + "\n"
	}

	fmt.Println("\nFound Emails:")
	fmt.Print(output)
}
