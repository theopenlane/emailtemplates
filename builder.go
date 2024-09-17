package emailtemplates

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rs/zerolog/log"
)

const (
	// Email templates must be provided in this directory and are loaded at compile time
	templatesDir = "templates"

	// Partials are included when rendering templates for composability and reuse - includes footer, header, etc.
	partialsDir = "partials"
)

var (
	//go:embed templates/*.html templates/*.txt templates/partials/*html
	files     embed.FS
	templates map[string]*template.Template
)

// Load templates when the package is imported
func init() {
	templates = make(map[string]*template.Template)

	templateFiles, err := fs.ReadDir(files, templatesDir)
	if err != nil {
		log.Panic().Err(err).Msg("could not read template files")
	}

	// Each template needs to be parsed independently to ensure that define directives
	// are not overwritten if they have the same name; e.g. to use the base template
	for _, file := range templateFiles {
		if file.IsDir() {
			continue
		}

		// Each template will be accessible by its base name in the global map
		patterns := make([]string, 0, 2) //nolint:mnd
		patterns = append(patterns, filepath.Join(templatesDir, file.Name()))

		if filepath.Ext(file.Name()) == ".html" {
			patterns = append(patterns, filepath.Join(templatesDir, partialsDir, "*.html"))
		}

		templates[file.Name()] = template.Must(template.ParseFS(files, patterns...))
	}
}

// Render returns the text and html executed templates for the specified name and data
func Render(name string, data interface{}) (text, html string, err error) {
	if text, err = render(name+".txt", data); err != nil {
		return "", "", err
	}

	if html, err = render(name+".html", data); err != nil {
		return "", "", err
	}

	return text, html, nil
}

func render(name string, data interface{}) (_ string, err error) {
	t, ok := templates[name]
	if !ok {
		return "", fmt.Errorf("%w: %q not found in templates", ErrMissingTemplate, name)
	}

	buf := &strings.Builder{}
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}