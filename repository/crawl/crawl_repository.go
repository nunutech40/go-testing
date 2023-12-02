package crawl_repository

import (
	"errors"
	crawler "go-testing/models/crawl"
)

type CrawlRepository struct {
	// Add fields if needed, e.g., database connection
}

func NewCrawlRepository() *CrawlRepository {
	return &CrawlRepository{}
}

// PerformCrawl calls CrawlData and returns its results or error.
func (repo *CrawlRepository) PerformCrawl(url string) (string, error) {
	if url == "" {
		return "", errors.New("URL cannot be empty")
	}

	crawl := crawler.Crawler{}
	return crawl.CrawlData(url) // Return results and error directly from CrawlData
}
