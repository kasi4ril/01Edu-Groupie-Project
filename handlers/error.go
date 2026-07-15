package handlers

import (
	"net/http"
)

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)

	err := RenderTemplate(w, "404.html", nil)
	if err != nil {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
	}
}

func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)

	err := RenderTemplate(w, "500.html", nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
