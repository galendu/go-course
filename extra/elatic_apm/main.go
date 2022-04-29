package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmhttp"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", mux.Vars(req)["name"])
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", apmhttp.Wrap(mux))
}
