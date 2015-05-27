package main

import (
	"flag"
	"fmt"
	"os"

	"strudelline.net/gunsafe/deliver"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Usage: downloader <msg url>\n")
		os.Exit(1)
	}
	err := deliver.Deliver(flag.Arg(0))
	if err != nil {
		panic(err)
	}
}
