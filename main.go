package main

import (
	"go-testing/database"
	crawl_handler "go-testing/handlers/crawl" // Assuming you have a crawl handler
	user_handler "go-testing/handlers/user"
	crawl_repository "go-testing/repository/crawl" // Assuming you have a crawl repository
	user_repository "go-testing/repository/user"
	"go-testing/router"
	crawl_service "go-testing/service/crawl" // Assuming you have a crawl service
	user_service "go-testing/service/user"
	"log"
	"net/http"
)

func main() {
	database.InitDB()
	userRepository := user_repository.NewUserRepository(database.DB)
	userService := user_service.NewUserRepository(userRepository)
	userHandler := user_handler.NewUserHandler(userService)

	// Crawl functionality setup
	crawlRepo := crawl_repository.NewCrawlRepository()       // Assuming NewCrawlRepository
	crawlServ := crawl_service.NewCrawlService(crawlRepo)    // Assuming NewCrawlService
	crawlHandler := crawl_handler.NewCrawlHandler(crawlServ) //

	router.SetupRoutes(userHandler, crawlHandler)

	http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
