package handlers
package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	path := filepath.Join("templates", tmpl)

	t, err := template.ParseFiles(path)
	if err != nil {
		return err
	}

	return t.Execute(w, data)
}