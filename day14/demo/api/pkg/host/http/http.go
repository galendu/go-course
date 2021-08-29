package http

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegistAPI(r *httprouter.Router) {
	r.GET("/users", queryUser)
	r.POST("/users", createUser)
	r.GET("/users/:id", describeUser)
	r.DELETE("/users/:id", deleteUser)
}

func queryUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "query user!\n")
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "create user!\n")
}

func describeUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "describe, %s!\n", ps.ByName("name"))
}

func deleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "delete, %s!\n", ps.ByName("name"))
}
