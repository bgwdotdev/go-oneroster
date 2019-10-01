package handlers

import (
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("https://github.com/fffnite/go-oneroster"))
}
