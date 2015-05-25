package main

import (
	"fmt"
	"log"
	"net/http"
)

// hello world, the web server
func HandleMailgunStoreForwardHook(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(500)
		fmt.Fprintf(w, "No thanks.\n")
	}

	if err := req.ParseForm(); err != nil {
		panic(err)
	}
	log.Print(req.FormValue("message-url"))
}

func main() {
	http.HandleFunc("/hello", HandleMailgunStoreForwardHook)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
