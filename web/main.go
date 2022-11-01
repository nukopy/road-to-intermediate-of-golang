package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nukopy/road-to-intermediate-of-golang/config"
	"github.com/nukopy/road-to-intermediate-of-golang/handlers"
)

func main() {
	// load config
	cfg := config.New()

	// register handlers
	http.HandleFunc("/healthcheck", handlers.HealthcheckHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.GetArticlesHandler)
	http.HandleFunc("/article/1", handlers.GetArticle1Handler)
	http.HandleFunc("/article/nice", handlers.PostArticleNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	// start http server
	log.Printf("Server start at port %d\n", cfg.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.PORT), nil))
}
