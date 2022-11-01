package main

import (
	"io"
	"log"
	"net/http"
)

/* interface: io.Writer
type Writer interface {
	Write(p []byte) (n int, err error)
}
*/

func main() {
	// register handlers
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", getArticleListHandler)
	http.HandleFunc("/article/1", getArticle1Handler)
	http.HandleFunc("/article/nice", postArticleNiceHandler)
	http.HandleFunc("/comment", postCommentHandler)

	// start http server
	log.Println("Server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "hello, world!\n")
	if err != nil {
		log.Println(err)
	}
	log.Printf("helloHandler: %d bytes were written.\n", n)
}

func postArticleHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Posting Article...\n")
	if err != nil {
		log.Println(err)
	}
	log.Printf("postArticleHandler: %d bytes were written.\n", n)
}

func getArticleListHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Article List\n")
	if err != nil {
		log.Println(err)
	}
	log.Printf("getArticleListHandler: %d bytes were written.\n", n)
}

func getArticle1Handler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Article 1\n")
	if err != nil {
		log.Println(err)
	}
	log.Printf("getArticle1Handler: %d bytes were written.\n", n)
}

func postArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Posting Nice...\n")
	if err != nil {
		log.Println(err)
	}
	log.Printf("postArticleNiceHandler: %d bytes were written.\n", n)
}

func postCommentHandler(w http.ResponseWriter, req *http.Request) {
	n, err := io.WriteString(w, "Posting Comment...\n")
	if err != nil {
		log.Println(err)
	}
	log.Printf("postCommentHandler: %d bytes were written.\n", n)
}