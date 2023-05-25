package main

import (
	"fmt"
	"net/http"
)

type Link struct {
	short string
	long  string
}



func formHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var link Link; 
		link.long = r.FormValue("URL")
		link.short = r.FormValue("extension")
		fmt.Println(link.long, link.short)	
		http.Handle("/"+link.short, http.RedirectHandler(link.long, 301))
	}
	http.ServeFile(w, r, "index.html")
	
	

}
func main() {

	http.HandleFunc("/", formHandler)
	fmt.Println("hello world")
	http.ListenAndServe(":8080", nil)

}
