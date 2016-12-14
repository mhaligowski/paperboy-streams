package streams

import (
	"net/http"
	"encoding/json"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type streamItemsGetter func(r *http.Request) ([]StreamItem, error)

type getStreamItemsHandler struct {
	getStreamItems streamItemsGetter
}

func (h getStreamItemsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	items, err := h.getStreamItems(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func dsGetter(r *http.Request) ([]StreamItem, error) {
	ctx := appengine.NewContext(r)

	query := datastore.NewQuery("StreamItem").Order("-OrderSequence")

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


