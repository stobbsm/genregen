package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/stobbsm/genregen/list"
)

// Start an http server to generate a unique
// genre type.
// © Matthew Stobbs, 2022

var (
	//go:embed vue/dist
	frontend embed.FS
)

type WordsResponse struct {
	Genre string `json:"genre"`
	Error error  `json:"error"`
}

func main() {

	guiDir, _ := fs.Sub(frontend, "vue/dist")
	gui := http.FileServer(http.FS(guiDir))
	http.Handle("/", gui)

	http.HandleFunc("/api", GetGenre)

	http.ListenAndServe(":8090", nil)
}

func GetGenre(w http.ResponseWriter, req *http.Request) {
	var res WordsResponse
	wordList, err := list.GetRandomN(2)
	if err != nil {
		res.Error = err
	} else {
		res.Genre = strings.Join(wordList, " ")
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	j, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	w.Write(j)
}
