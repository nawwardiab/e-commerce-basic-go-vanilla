package view

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"text/template"
)

// parseTemplate – parses the templates files and assing the returned template.Template type to global variable
var templates = parseTemplates("templates")

func parseTemplates(dir string) *template.Template{
	pattern := filepath.Join(dir, "*.tpl")
  t, err := template.ParseGlob(pattern)
  if err != nil {
    log.Fatalf("failed to parse templates: %v", err)
  }
	return t
}

// Render – execute the template with specific file name and its required data
func Render(w io.Writer, name string, data interface{}) error {
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		return fmt.Errorf("view: executing template %q: %w", name, err)
	}
	return nil
}