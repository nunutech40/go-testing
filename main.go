package main

import (
	"go-testing/database"
	"go-testing/handlers"
	"go-testing/module/crawler"
	"go-testing/repository"
	"go-testing/router"
	"go-testing/service"
	"net/http"
)

func main() {
	database.InitDB()
	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserRepository(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	router.SetupRoutes(userHandler)

	crawl := crawler.Crawler{}
	crawl.CrawlData("http://example.com")

	http.ListenAndServe(":8080", nil)
}
