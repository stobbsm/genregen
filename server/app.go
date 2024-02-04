package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/stobbsm/genregen/list"
)

type WordsResponse struct {
	Genre string `json:"genre"`
	Error error  `json:"error"`
}

func App() {
	//guiDir, _ := fs.Sub(frontend, "vue/dist")
	//gui := http.FileServer(http.FS(guiDir))
	gui := http.FileServer(http.Dir("vue/dist"))
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
