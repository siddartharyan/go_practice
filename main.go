package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allAtricles(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintf(w, "POST end point hit")
		return
	}
	articles := Articles{
		Article{Title: "siddarth", Desc: "associate", Content: "r-pay"},
		Article{Title: "sai", Desc: "assistant", Content: "vas"},
	}
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page Endpoint hit")
}

func handleRequest() {
	http.HandleFunc("/articles", allAtricles)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
