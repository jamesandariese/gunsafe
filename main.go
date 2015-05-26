package main

import (
	"fmt"
	"log"
	"net/http"
	"flag"
)

var apikey = flag.String("apikey", "", "API key for mailgun")
var spooldir = flag.String("spool", "/spool", "Spool directory")

// hello world, the web server
func HandleMailgunStoreForwardHook(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(500)
		fmt.Fprintf(w, "No thanks.\n")
	}

	if err := req.ParseForm(); err != nil {
		panic(err)
	}

	fd, err := ioutil.TempFile(*spooldir, "gunsafe")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	fmt.Fprintf(fd, "%s", req.FormValue("message-url"))

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
