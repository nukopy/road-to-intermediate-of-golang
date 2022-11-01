package handlers

import (
	"io"
	"log"
	"net/http"
)

type HttpMethod string

func HealthcheckHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		msg := "Invalid method"
		log.Printf("healthcheck: %s\n", msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}

	n, err := io.WriteString(w, "Healthcheck OK!\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("healthcheck: %d bytes were written.\n", n)
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		msg := "Invalid method"
		log.Printf("PostArticleHandler: %s\n", msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}

	n, err := io.WriteString(w, "Posting Article...\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PostArticleHandler: %d bytes were written.\n", n)
}

func GetArticlesHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		msg := "Invalid method"
		log.Printf("GetArticlesHandler: %s\n", msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}

	n, err := io.WriteString(w, "Article List\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("getArticleListHandler: %d bytes were written.\n", n)
}

func GetArticle1Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		msg := "Invalid method"
		log.Printf("GetArticle1Handler: %s\n", msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}

	n, err := io.WriteString(w, "Article 1\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("getArticle1Handler: %d bytes were written.\n", n)
}

func PostArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		msg := "Invalid method"
		log.Printf("PostArticleNiceHandler: %s\n", msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}

	n, err := io.WriteString(w, "Posting Nice...\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PostArticleNiceHandler: %d bytes were written.\n", n)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		msg := "Invalid method"
		log.Printf("PostCommentHandler: %s\n", msg)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}

	n, err := io.WriteString(w, "Posting Comment...\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PostCommentHandler: %d bytes were written.\n", n)
}