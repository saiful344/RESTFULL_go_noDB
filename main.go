package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)


type Article struct{
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var article []Article
func RetrieveArticle(w http.ResponseWriter, r *http.Request){
	fmt.Println("hit:retrieve")
	json.NewEncoder(w).Encode(article)
}
func homepage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Welcome to home")
	fmt.Println("ENDPOINT HIT : homepage")
}
func HomePageShow(){
	http.HandleFunc("/",homepage)
	http.HandleFunc("/get",RetrieveArticle)
	log.Fatal(http.ListenAndServe(":8000",nil))
}
func main(){

	article = []Article{
		Article{Title:"name",Desc:"kjskd",Content:"kjsdkskjfhjshf"},
		Article{Title:"name2",Desc:"kjs2kd",Content:"kjsdksk2jfhjshf"},
	}
	HomePageShow()
}