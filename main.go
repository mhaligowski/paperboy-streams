package streams

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Run() {
	r := mux.NewRouter();

	r.Handle("/items",
		getStreamItemsHandler{dsGetter}).
		Methods(http.MethodGet)


	r.Handle("/jobs",
		startJobHandler{dsItemsWriter}).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")

	http.Handle("/", cors.Default().Handler(r))
}
