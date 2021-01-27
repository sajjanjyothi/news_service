package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type newsRouter struct {
	URL     string
	handler func(w http.ResponseWriter, r *http.Request)
	Method  string
}

var router = []newsRouter{
	{"/api/v1/news_list/", newsArticleHandler, "POST"},
}

func setRoutes(route *mux.Router) {
	for _, newsRoute := range router {
		route.HandleFunc(newsRoute.URL, newsArticleHandler).Methods(newsRoute.Method)
	}
}
