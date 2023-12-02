package crawl_handler

import (
	"encoding/json"
	crawl_service "go-testing/service/crawl"
	"net/http"
)

type CrawlHandler struct {
	service *crawl_service.CrawlService
}

func NewCrawlHandler(service *crawl_service.CrawlService) *CrawlHandler {
	return &CrawlHandler{service: service}
}

type URLPayload struct {
	URL string `json:"url"`
}

func (h *CrawlHandler) HandleCrawl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	var payload URLPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	if payload.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// Call CrawlURL and handle the response and error
	crawlResult, err := h.service.CrawlURL(payload.URL)
	if err != nil {
		http.Error(w, "Error performing crawl: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": crawlResult})
}
