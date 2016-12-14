package streams

import (
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter();

	r.Handle("/items",
		getStreamItemsHandler{dsGetter})

	r.Handle("/jobs",
		startJobHandler{dsItemsWriter}).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")

	http.Handle("/", r)
}
