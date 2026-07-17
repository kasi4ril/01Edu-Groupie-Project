package handlers

import (
	"net/http"
)

// NotFound renders the 404 error page.
func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)

	if err := RenderTemplate(w, "404.html", nil); err != nil {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
	}
}

// InternalServerError renders the 500 error page.
func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)

	if err := RenderTemplate(w, "500.html", nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}