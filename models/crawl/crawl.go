package crawler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Crawler struct{}

// CrawlData now returns the concatenated text of all links found and an error if any occurs.
func (c *Crawler) CrawlData(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err // Return the error instead of logging it
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err // Return the error instead of logging it
	}

	var linksText strings.Builder
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		linkText := s.Text()
		linksText.WriteString(linkText + "\n") // Concatenate the link text with a newline
	})

	return linksText.String(), nil // Return the result and a nil error
}
