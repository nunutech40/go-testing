package crawler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Crawler struct{}

// CrawlData now returns the title and the concatenated text of all article paragraphs found, and an error if any occurs.
func (c *Crawler) CrawlData(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	// Find the article title, assuming it's within an <h1> tag
	var title string
	doc.Find("h1").Each(func(i int, s *goquery.Selection) {
		if title == "" { // Only take the first <h1> tag as the title
			title = s.Text()
		}
	})

	// Find and concatenate all article paragraph texts
	var articleText strings.Builder
	doc.Find("article p").Each(func(i int, s *goquery.Selection) {
		paragraph := s.Text()
		articleText.WriteString(paragraph + "\n")
	})

	// Combine the title and the article text
	fullArticle := title + "\n\n" + articleText.String()
	return fullArticle, nil
}
