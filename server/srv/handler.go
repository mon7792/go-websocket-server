package srv

import (
	"fmt"
	"net/http"
	"time"
)

func NewHandler() http.Handler {
	nsw := http.NewServeMux()
	nsw.HandleFunc("/", homeHandler)
	return nsw
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "Hello World")
}
