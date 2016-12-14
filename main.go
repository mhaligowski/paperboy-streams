package streams

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func init() {
	r := mux.NewRouter();

	r.Handle("/items",
		handlers.CORS()(getStreamItemsHandler{dsGetter}))

	r.Handle("/jobs",
		startJobHandler{dsItemsWriter}).
		Methods(http.MethodPost).
		Headers("Content-Type", "application/json")

	http.Handle("/", r)
}
