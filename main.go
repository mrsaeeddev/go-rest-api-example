package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Working")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles 🚀")
	json.NewEncoder(w).Encode((Articles))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Rest APIs")
	fmt.Println("Endpoint hit: 🚀")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", returnAllArticles)
	router.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe((":10000"), router))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		{
			Id:      "1",
			Title:   "IK steps down as PM of Pakistan",
			Desc:    "IK steps down as PM of Pakistan",
			Content: "IK steps down as PM of Pakistan",
		},
		{
			Id:      "2",
			Title:   "SS takes oath as PM of Pakistan",
			Desc:    "SS takes oath as PM of Pakistan",
			Content: "SS takes oath as PM of Pakistan",
		},
	}
	handleRequests()
}

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article
