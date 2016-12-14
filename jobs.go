package streams

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"github.com/mhaligowski/paperboy-crawler"
)

type startJobHandler struct {
	write streamItemsWriter
}

func (h startJobHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	input, err := jsonParser(r)

	log.Debugf(ctx, "Got %d requests", len(input.Entries))

	if err != nil {
		log.Errorf(ctx, "Could not parse the input from %v", r.Body)
		http.Error(w, "Could not parse the input", http.StatusBadRequest)
		return
	}

	// TODO what happens with doubles, probably should start a transaction on a row

	// TODO get subscriptions
	userIds := []string{"dummy_user_id"}

	// TODO parallelize this
	for _, userId := range userIds {
		// TODO unignore errors
		items, _ := buildStreamItems(input, userId)
		log.Debugf(ctx, "Built %d items for user %q", len(items), userId)

		err = h.write(ctx, items)

		if err != nil {
			log.Errorf(ctx, "Detected problem when trying to write for user id %s: %v", userId, err)
		}
	}
}

func buildStreamItems(update *crawler.StreamUpdate, userId string) ([]StreamItem, []error) {
	result := make([]StreamItem, 0, len(update.Entries))
	errors := make([]error, 0, len(update.Entries))

	for _, entry := range update.Entries {
		r := StreamItem{
			Title: entry.Title,
			FeedTitle: update.Title,
			OrderSequence: time.Now().UnixNano(),
			TargetId:entry.Id,
			UserId:userId,
		}
		r.StreamItemId = id(&r)
		result = append(result, r)
	}

	return result, errors
}

func id(s *StreamItem) string {
	writer := sha256.New()
	writer.Write([]byte(s.TargetId))
	writer.Write([]byte(s.UserId))

	return hex.EncodeToString(writer.Sum(nil))
}
