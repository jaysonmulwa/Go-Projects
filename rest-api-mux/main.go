package main

import (

	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	
)

type Article struct {

	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`

}

type Articles []Article


func allArticles(w http.ResponseWriter, r *http.Request){

	articles := Articles {

		Article { Title: "My Life in Crime", Desc: "Biography", Content: "John Kiriamiti"},
		Article { Title: "The Secret", Desc: "New Age", Content: "JRhonda Bryne"},

	}

	fmt.Println("Endpoint MIT: All Articles Endpoint")

	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Test POST Endpoints")

}


func homePage(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "HomePage Endpoint MIT")
}

func handleRequest (){

	//mux helps us specify verbs

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)

	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", myRouter))

}

func main() {

	
	handleRequest()

}
