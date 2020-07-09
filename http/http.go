package example

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	route "github.com/milamice62/fakeapi/api/routes"
)

func HttpServer() {
	err := http.ListenAndServe("localhost:8888", handler())
	if err != nil {
		log.Fatalf("Error occurs: %v", err)
	}
}

func handler() http.Handler {
	r := mux.NewRouter()
	route.Genres(r)
	route.Movies(r)
	return r
}

func doubleHandler(res http.ResponseWriter, req *http.Request) {
	// retrieve query value
	queryValue := req.FormValue("v")
	if queryValue == "" {
		http.Error(res, "Missing Query Value", http.StatusBadRequest)
		return
	}
	// convert string to number
	num, err := strconv.Atoi(queryValue)
	if err != nil {
		http.Error(res, "not a number: "+queryValue, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(res, num*2)
}

func tripleHandler(res http.ResponseWriter, req *http.Request) {
	queryValue := req.FormValue("v")
	if queryValue == "" {
		http.Error(res, "Missing Query Value", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(queryValue)
	if err != nil {
		http.Error(res, "not a number: "+queryValue, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(res, num*3)
}
