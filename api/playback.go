package handler

import (
	"fmt"
	"net/http"
)

func Playback(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
