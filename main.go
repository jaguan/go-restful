package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article is ...
type Article struct {
	ID      string `json:ID"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content`
}

// Articles is ...
var Articles []Article

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)

	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")

	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleArticle")

	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}
