package feedupdater

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)


func handleGetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func init() {
	r := mux.NewRouter();
	r.HandleFunc("/items", handleGetItems)

	http.Handle("/", r)
}
