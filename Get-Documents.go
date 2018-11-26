package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Document struct {
	Id   string
	Name string
	Size int
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/documents", getDocuments).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", router))
}

func getDocuments(w http.ResponseWriter, r *http.Request) {
	var docs []Document
	id := 0

	// Open current directory
	f, err := os.Open("./Docs")
	if err != nil {
		panic(err)
	}

	// Get file
	files, err := f.Readdir(0)
	if err != nil {
		panic(err)
	}

	// Add info file
	for _, v := range files {
		id++
		docs = append(docs, Document{Id: "Doc-" + strconv.Itoa(id), Name: v.Name(), Size: int(v.Size())})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(docs)
}
