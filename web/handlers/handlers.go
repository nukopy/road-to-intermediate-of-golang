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
	// extract query parameter "page"
	qp := req.URL.Query()
	log.Printf("Query Params: %v\n", qp)
	pageValues, ok := qp["page"]

	// page パラメータがないとき、page=1 としてレスポンスを返す
	if !ok {
		page := 1
		n, err := io.WriteString(w, fmt.Sprintf("Articles List - Page %d\n", page))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("getArticleListHandler: %d bytes were written.\n", n)
		return
	}

	// page パラメータが 1 回以上指定されている場合、最初の値を指定する
	pageStr := pageValues[0]
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		// page パラメータが整数に変換できない場合、400 エラーレスポンスを返す
		http.Error(w, fmt.Sprintf("Invalid query parameter: \"%s\"\n", pageStr), http.StatusBadRequest)
		return
	}

	// 指定された page パラメータに応じてレスポンスを返す
	n, err := io.WriteString(w, fmt.Sprintf("Articles List - Page %d\n", page))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("getArticleListHandler: %d bytes were written.\n", n)
}

func GetArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleId, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		// ルーティングの実装上ここには来ないはず
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