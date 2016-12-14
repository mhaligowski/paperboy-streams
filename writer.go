package streams

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/datastore"
)

type streamItemsWriter func(context.Context, []StreamItem) (error)

func dsItemsWriter(ctx context.Context, items []StreamItem) (error) {
	keys := make([]*datastore.Key, len(items))
	for i, v := range items {
		keys[i] = datastore.NewKey(ctx, "StreamItem", v.StreamItemId, 0, nil)
	}

	_, err := datastore.PutMulti(ctx, keys, items)

	log.Debugf(ctx, "Written %d items", len(items))

	if err != nil {
		return err
	} else {
		return nil
	}
}
