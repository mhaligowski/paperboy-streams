package streams

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestEmpty(t *testing.T) {
	f := func(r *http.Request) ([]StreamItem, error) {
		return make([]StreamItem, 0), nil
	}

	w := httptest.NewRecorder()
	handler := getStreamItemsHandler{f}
	handler.ServeHTTP(w, nil)

	if w.Body.String() != "[]\n" {
		t.Errorf("Expected [], got %q", w.Body.String())
	}
}


