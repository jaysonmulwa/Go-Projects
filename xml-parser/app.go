package main

import (

	"fmt"
	"net/http"
	"html/template"
	"io/ioutil"
	"encoding/xml"
	"sync"
)

var wg sync.WaitGroup 

type SitemapIndex struct {
	
	Locations []string `xml:"sitemap>loc"`
}

type News struct {

	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {

	Keyword string 
	Location string 

} 

type NewAggPage struct{
	Title string
	News map[string]NewsMap

}

func cleanup() {

	if r := recover(); r != nil {

		fmt.Println("recovered in cleanup: a 2!", r)
	}

}

func newsRoutine(c chan News, Location string){

	defer wg.Done()
	defer cleanup()
	var n News

	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()

	c <- n

}

func newsAggHandler(w http.ResponseWriter, r *http.Request){

	var s SitemapIndex
	
	
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	news_map := make(map[string]NewsMap)
	resp.Body.Close()
	queue := make(chan News, 10)


	for _, Location := range s.Locations {

		wg.Add(1)
		go newsRoutine(queue, Location)
			

	}

	wg.Wait()
	close(queue)

	for elem := range queue {

		for idx, _ := range elem.Keywords {

			news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
		}

	}

		

	p:= NewAggPage{Title:"News Aggregator", News:news_map}

	t, _ := template.ParseFiles("basictemplating.html")

	//t.Execute(w, p)

	//incase of error
	fmt.Println(t.Execute(w, p))



	

}


func index_handler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "<p>Jayson %s Mulwa</p>", "Mulwa")

}



func main() {

	http.HandleFunc("/",index_handler)
	http.HandleFunc("/agg/",newsAggHandler)
	http.ListenAndServe(":8000", nil)

}
