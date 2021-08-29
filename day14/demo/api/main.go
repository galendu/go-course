package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	host "gitee.com/infraboard/go-course/day14/demo/api/pkg/host/http"
)

var (
	addr = ":8080"
)

func main() {
	router := httprouter.New()
	host.RegistAPI(router)

	log.Printf("listen on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
