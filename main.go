package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"strudelline.net/gunsafe/deliver"
)

var apikey = flag.String("apikey", "", "API key for mailgun")

// hello world, the web server
func HandleMailgunStoreForwardHook(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(500)
		fmt.Fprintf(w, "No thanks.\n")
	}

	if err := req.ParseForm(); err != nil {
		panic(err)
	}

	err := deliver.Deliver(req.FormValue("message-url"))
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error processing: %v\n", err)
		log.Printf("ERROR: URL: %s %v", req.FormValue("message-url"), err)
	} else {
		w.WriteHeader(200)
		fmt.Fprintf(w, "Delivered")
	}

	/*
		request, err := http.NewRequest("GET", req.FormValue("message-url"), nil)
		if err != nil {
			panic(err)
		}
		request.SetBasicAuth("api", *apikey)

		request.Header.Add("Accept", "message/rfc2822")

		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			panic(err)
		}
	*/
}

func main() {
	http.HandleFunc("/gunsafe", HandleMailgunStoreForwardHook)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
