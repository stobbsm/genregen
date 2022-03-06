package main

import (
	"embed"
	"html/template"
	"net/http"
	"strings"

	"github.com/stobbsm/genregen/list"
)

// Start an http server to generate a unique
// genre type.
// © Matthew Stobbs, 2022

var (
	//go:embed templates
	templates embed.FS
)

type Metadata struct {
	Author    string
	Copyright string
	Email     string
	Title     string
}

func main() {
	http.HandleFunc("/", WordsHandler)

	http.ListenAndServe(":8090", nil)
}

func WordsHandler(w http.ResponseWriter, req *http.Request) {
	wordList, err := list.GetRandomN(2)
	if err != nil {
		panic(err)
	}
	Meta := Metadata{
		Author:    "Matthew Stobbs",
		Copyright: "Matthew Stobbs",
		Email:     "matthew@stobbs.ca",
		Title:     "Genre Generator",
	}
	Data := struct {
		Meta    Metadata
		Content string
	}{
		Meta:    Meta,
		Content: strings.Join(wordList, " "),
	}
	tmpl := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	tmpl.ExecuteTemplate(w, "layout", Data)
}
