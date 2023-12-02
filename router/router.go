package router

import (
	crawl_handler "go-testing/handlers/crawl"
	user_handler "go-testing/handlers/user"
	"net/http"
)

func SetupRoutes(userHandler *user_handler.UserHandler, crawlHandler *crawl_handler.CrawlHandler) {
	http.HandleFunc("/users", userHandler.GetUsers)
	http.HandleFunc("/crawl", crawlHandler.HandleCrawl)

	
}
