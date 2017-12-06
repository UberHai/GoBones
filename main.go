//Name of our package
package main

//Important libraries
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

//Create a new router and return to instantiate and test the router outside the main func
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

func main() {
	//Construct a new router
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}