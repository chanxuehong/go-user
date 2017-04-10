package main

import (
	"flag"
	"log"
	"net/http"
	"runtime"

	"github.com/chanxuehong/go-user/frontend"
)

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	log.Println(http.ListenAndServe(":80", frontend.Engine))
}
