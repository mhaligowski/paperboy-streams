package streams

import (
	"net/http"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"fmt"
	"google.golang.org/appengine/log"
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

	return items, err
}

func handleGetStreamItems(w http.ResponseWriter, r *http.Request, g streamItemsGetter) {
	items, err := g(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, item := range items {
		fmt.Fprintf(w, "%v\n", item.Title)
	}
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
