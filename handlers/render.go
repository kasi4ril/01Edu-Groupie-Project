package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// RenderTemplate parses and executes an HTML template.
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	path := filepath.Join("templates", tmpl)

	t, err := template.ParseFiles(path)
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}

	if err := t.Execute(w, data); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	return nil
}