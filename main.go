package main
import (
	"net/http"
	"github.com/gorilla/mux"

	"github.com/futurice/door2door-hackaton/api"
)


func main() {

	router := mux.NewRouter()
	router.
		Methods("GET").
		Path("/api/hello").
		Name("Hello").
		Handler(http.HandlerFunc(api.Greeter))

	http.ListenAndServe(":8081", api.CorsHeaders(router))

}
