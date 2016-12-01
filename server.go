package streams

import (
	"net/http"
	"github.com/gorilla/mux"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"

	"encoding/json"
)

type streamItemsGetter func(r *http.Request) ([]StreamItem, error)

func dsGetter(r *http.Request) ([]StreamItem, error) {
	ctx := appengine.NewContext(r)

	query := datastore.NewQuery("StreamItem").Order("-order_sequence")

	var items []StreamItem
	_, err := query.GetAll(ctx, &items)

	if err != nil {
		log.Errorf(ctx, "Error when querying from datastore", err)
	}

	if items == nil {
		return make([]StreamItem, 0), err
	} else {
		return items, err
	}
}

func handleGetStreamItems(w http.ResponseWriter, r *http.Request, g streamItemsGetter) {
	items, err := g(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func Start() {
	r := mux.NewRouter();
	r.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		handleGetStreamItems(w, r, dsGetter)
	})

	http.Handle("/", r)
}

func init() {
	Start()
}
