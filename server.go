package streams

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func Start() {
	r := mux.NewRouter();
	r.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		HandleGetStreamItems(w, r, dsGetter)
	})

	http.Handle("/", handlers.CORS()(r))
}

func init() {
	Start()
}
