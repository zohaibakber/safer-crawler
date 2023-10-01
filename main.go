package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func crawlForMCMX(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mcmx := vars["mcmx"]
	fmt.Fprintf(w, "MC is %s",mcmx)
}

func main()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{mcmx}", crawlForMCMX)
	http.ListenAndServe(":8080", router)

}
