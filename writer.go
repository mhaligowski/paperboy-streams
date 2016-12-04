package streams

import (
	"google.golang.org/appengine/datastore"
	"golang.org/x/net/context"
)

type streamItemsWriter func(context.Context, []StreamItem) (error)

func dsItemsWriter(ctx context.Context, items []StreamItem) (error) {
	keys := make([]*datastore.Key, len(items))
	for i, v := range items {
		keys[i] = datastore.NewKey(ctx, "StreamItem", v.StreamItemId, 0, nil)
	}

	_, err := datastore.PutMulti(ctx, keys, items)

	if err != nil {
		return err
	} else {
		return nil
	}
}


