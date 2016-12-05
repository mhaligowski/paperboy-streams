package streams

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func init() {
	r := mux.NewRouter();
	r.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		HandleGetStreamItems(w, r, dsGetter)
	})

	r.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		HandleStartJob(w, r, jsonParser, dsItemsWriter)
	}).Methods(http.MethodPost).Headers("Content-Type", "application/json")

	r.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		HandleStartJob(w, r, formParser, dsItemsWriter)
	}).Methods(http.MethodPost).Headers("Content-Type", "application/x-www-form-encoded")


	http.Handle("/", handlers.CORS()(r))
}
