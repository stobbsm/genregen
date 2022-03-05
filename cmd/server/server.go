package main

import (
	"html/template"
	"net/http"
	"strings"

	words "github.com/stobbsm/genregen"
)

// Start an http server to generate a unique
// genre type.

func main() {
	http.HandleFunc("/", WordsHandler)
	http.ListenAndServe(":8090", nil)
}

func WordsHandler(w http.ResponseWriter, req *http.Request) {
	wordList, err := words.GetRandomN(2)
	if err != nil {
		panic(err)
	}
	tmpl := template.Must(template.ParseFiles("tmpl/layout.html"))
	tmpl.Execute(w, strings.Join(wordList, " "))
}
