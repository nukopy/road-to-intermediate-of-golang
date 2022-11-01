package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HttpMethod string

func HealthcheckHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Healthcheck OK!\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("healthcheck: %d bytes were written.\n", n)
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Posting Article...\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PostArticleHandler: %d bytes were written.\n", n)
}

func GetArticlesHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Article List\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("getArticleListHandler: %d bytes were written.\n", n)
}

func GetArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid path parameter", http.StatusBadRequest)
		return
	}

	n, err := io.WriteString(w, fmt.Sprintf("Article No.%d\n", articleId))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("getArticle1Handler: %d bytes were written.\n", n)
}

func PostArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Posting Nice...\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PostArticleNiceHandler: %d bytes were written.\n", n)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Posting Comment...\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PostCommentHandler: %d bytes were written.\n", n)
}