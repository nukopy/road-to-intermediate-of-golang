package handlers

import (
	"io"
	"log"
	"net/http"
)

func HealthcheckHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Healthcheck OK!\n")
	if err != nil {
		log.Println(err)
	}
	log.Printf("healthcheck: %d bytes were written.\n", n)
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Posting Article...\n")
	if err != nil {
		log.Println(err)
	}
	log.Printf("postArticleHandler: %d bytes were written.\n", n)
}

func GetArticlesHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Article List\n")
	if err != nil {
		log.Println(err)
	}
	log.Printf("getArticleListHandler: %d bytes were written.\n", n)
}

func GetArticle1Handler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Article 1\n")
	if err != nil {
		log.Println(err)
	}
	log.Printf("getArticle1Handler: %d bytes were written.\n", n)
}

func PostArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Posting Nice...\n")
	if err != nil {
		log.Println(err)
	}
	log.Printf("postArticleNiceHandler: %d bytes were written.\n", n)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Posting Comment...\n")
	if err != nil {
		log.Println(err)
	}
	log.Printf("postCommentHandler: %d bytes were written.\n", n)
}