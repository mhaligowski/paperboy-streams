package streams

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"encoding/json"
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
