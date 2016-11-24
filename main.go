package feedupdater

import (
	"net/http"
	"fmt"
)


func init() {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!")
	})

}
