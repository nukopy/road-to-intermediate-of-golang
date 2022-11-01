package main

import (
	"log"
	"net/http"

	"github.com/nukopy/road-to-intermediate-of-golang/handlers"
)

func main() {
	// register handlers
	http.HandleFunc("/healthcheck", handlers.HealthcheckHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.GetArticlesHandler)
	http.HandleFunc("/article/1", handlers.GetArticle1Handler)
	http.HandleFunc("/article/nice", handlers.PostArticleNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	// start http server
	log.Println("Server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
