package streams

import (
	"net/http"
	"github.com/gorilla/schema"
	"github.com/mhaligowski/paperboy-crawler"
	"encoding/json"
)

type StreamUpdate struct {
	crawler.Feed
}

var formDecoder = schema.NewDecoder()

type inputParser func(*http.Request) (*StreamUpdate, error)

func formParser(r *http.Request) (*StreamUpdate, error) {
	input := new(StreamUpdate)
	err := formDecoder.Decode(input, r.Form)

	if err != nil {
		return nil, err
	}

	return input, nil
}

func jsonParser(r *http.Request) (*StreamUpdate, error) {
	input := new(StreamUpdate)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(input)

	if err != nil {
		return nil, err
	}

	return input, nil
}

