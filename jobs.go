package streams

import (
	"net/http"
	"time"

	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func buildStreamItems(update *StreamUpdate, userId string) ([]StreamItem, []error) {
	result := make([]StreamItem, 0, len(update.Entries))
	errors := make([]error, 0, len(update.Entries))

	for _, entry := range update.Entries {
		keyValue, err := uuid.NewV4();
		if err != nil {
			errors = append(errors, err)
			continue
		}

		r := StreamItem{
			Title: entry.Title,
			FeedTitle: update.Title,
			OrderSequence: time.Now().UnixNano(),
			StreamItemId:keyValue.String(),
			TargetId:entry.Id,
			UserId:userId,
		}
		result = append(result, r)
	}

	return result, errors
}

func HandleStartJob(w http.ResponseWriter, r *http.Request, ip inputParser, siw streamItemsWriter) {
	ctx := appengine.NewContext(r)

	input, err := ip(r)

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

		err = siw(ctx, items)

		if err != nil {
			log.Errorf(ctx, "Detected problem when trying to write for user id %s: %v", userId, err)
		}
	}
}
