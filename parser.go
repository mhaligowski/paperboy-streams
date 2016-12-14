package streams

import (
	"net/http"
	"encoding/json"
	"github.com/mhaligowski/paperboy-crawler"
)

func jsonParser(r *http.Request) (*crawler.StreamUpdate, error) {
	input := new(crawler.StreamUpdate)
	err := json.NewDecoder(r.Body).Decode(input)

	if err != nil {
		return nil, err
	}

	return input, nil
}

