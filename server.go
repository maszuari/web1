package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	modeler "github.com/maszuari/web1/models"
)

type Handler struct {
	db modeler.Datasource
}

func (env *Handler) FindOrgByOrgname(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fname, err := env.db.FindOrgByName(vars["name"])
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Database error")
	}
	fmt.Fprintf(w, "Organization name: %v\n", fname)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["name"])
}

func main() {
	mydb, err := modeler.Connect()
	if err != nil {
		log.Fatal(err)
	}
	h := &Handler{db: mydb}
	defer mydb.Close()

	r := mux.NewRouter()
	//r.HandleFunc("/hello/{name}", HomeHandler).Methods("GET")
	r.HandleFunc("/org/{name}", h.FindOrgByOrgname).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
