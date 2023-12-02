package crawl_service

import (
	crawl_repository "go-testing/repository/crawl"
)

type CrawlService struct {
	repo *crawl_repository.CrawlRepository
}

// NewCrawlService creates a new instance of CrawlService
func NewCrawlService(repo *crawl_repository.CrawlRepository) *CrawlService {
	return &CrawlService{repo: repo}
}

// CrawlURL initiates a crawl operation for the given URL and returns the results
func (s *CrawlService) CrawlURL(url string) (string, error) {
	return s.repo.PerformCrawl(url)
}
