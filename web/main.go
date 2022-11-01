package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nukopy/road-to-intermediate-of-golang/config"
	"github.com/nukopy/road-to-intermediate-of-golang/handlers"
)

func main() {
	// load config
	log.Println("Loading config...")
	cfg := config.New()

	// register handlers
	log.Println("Registering HTTP request handlers...")
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", handlers.HealthcheckHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.GetArticlesHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.GetArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	// start http server
	log.Printf("Server start at port %d\n", cfg.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.PORT), r))
}
