package streams

import (
	"net/http"
	"github.com/gorilla/mux"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"fmt"
	"google.golang.org/appengine/log"
)

type StreamItem struct {
	StreamItemId  string `datastore:"id"`;
	UserId        string;
	TargetId      string;
	Title         string;
	OrderSequence int64;
}

func handleGetItems(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	query := datastore.NewQuery("StreamItem").Order("-order_sequence")

	var items []StreamItem
	_, err := query.GetAll(ctx, &items)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Errorf(ctx, "Error when querying from datastore", err)
		return
	}

	for _, item := range items {
		fmt.Fprintf(w, "%v\n", item.Title)
	}

}

func init() {
	r := mux.NewRouter();
	r.HandleFunc("/items", handleGetItems)

	http.Handle("/", r)
}
