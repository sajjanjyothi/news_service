package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	address  = "localhost:50051"
	restPort = "0.0.0.0:8000"
)

type database struct {
	Title     string
	Timestamp time.Time
	Tags      []string
}

var newsDatabase = []database{
	database{"News title 1", time.Now(), []string{"sports"}},
	database{"News title 2", time.Now(), []string{"live"}},
	database{"News title 3", time.Now(), []string{"live", "sports"}},
	database{"News title 4", time.Now(), []string{"sports"}},
	database{"News title 5", time.Now(), []string{"live"}},
	database{"News title 6", time.Now(), []string{"live"}},
	database{"News title 7", time.Now(), []string{"sports"}},
	database{"News title 8", time.Now(), []string{"sports", "live"}},
	database{"News title 9", time.Now(), []string{"sports"}},
	database{"News title 10", time.Now(), []string{"live"}},
}

func _GetDataBase() []database {
	return newsDatabase
}

func main() {
	route := mux.NewRouter()

	//Set the news article routes
	setRoutes(route)

	log.Println("REST server listening on port ", restPort)
	log.Fatal(http.ListenAndServe(restPort, route))
}
