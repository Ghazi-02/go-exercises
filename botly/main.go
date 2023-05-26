package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Link struct {
	short string
	long  string
}

var linkMap = make(map[string]string)

func formHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var link Link; 
		link.long = r.FormValue("URL")
		link.short = r.FormValue("extension")
		linkMap["localhost:8080/"+link.short] = link.long	
		fmt.Println(linkMap)
		http.Handle("/"+link.short, http.RedirectHandler(link.long, 301))
		for k,v := range linkMap{
			fmt.Fprint(w,"<div>Your link: <div>",k,"<div> Directs to:</div> ",v)
			fmt.Fprint(w,"<p></p>")
		}	
		jsonData,err:= json.Marshal(linkMap)
		if err != nil {
			log.Fatal(err)
		}
		filePath := "linkMap.json"
		err = os.WriteFile(filePath,jsonData,0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	http.ServeFile(w, r, "index.html")
	


}
func main() {

	http.HandleFunc("/", formHandler)
	fmt.Println("Running on PORT: 8080 ...")
	http.ListenAndServe(":8080", nil)

}
