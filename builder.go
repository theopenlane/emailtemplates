package emailtemplates

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rs/zerolog/log"
	"github.com/stoewer/go-strcase"
)

const (
	// Email templates must be provided in this directory and are loaded at compile time
	defaultTemplatesDir = "templates"

	// Partials are included when rendering templates for composability and reuse - includes footer, header, etc.
	defaultPartialsDir = "partials"
)

var (
	//go:embed templates/*.html templates/*.txt templates/partials/*html templates/partials/*txt
	files     embed.FS
	templates map[string]*template.Template
	partials  []string

	// Shared function map
	fm = template.FuncMap{
		"ToUpper": strcase.UpperCamelCase,
	}
)

// Load templates when the package is imported
func init() {
	templates = make(map[string]*template.Template)

	templateFiles, err := fs.ReadDir(files, defaultTemplatesDir)
	if err != nil {
		log.Panic().Err(err).Msg("could not read template files")
	}

	// Each template needs to be parsed independently to ensure that define directives
	// are not overwritten if they have the same name; e.g. to use the base template

	for _, file := range templateFiles {
		if file.IsDir() {
			continue
		}

		templates[file.Name()] = parseTemplate(file.Name())
	}
}

func parseTemplate(name string) *template.Template {
	// Each template will be accessible by its base name in the global map
	patterns := []string{}
	patterns = append(patterns, filepath.Join(defaultTemplatesDir, name))

	validExtensions := []string{".txt", ".html"}
	for _, ext := range validExtensions {
		if filepath.Ext(name) == ext {
			patterns = append(patterns, filepath.Join(
				defaultTemplatesDir,
				defaultPartialsDir, "*"+ext))
		}
	}

	tmpl, err := template.New(name).Funcs(fm).ParseFS(files, patterns...)
	if err != nil {
		log.Fatal().Err(err).Str("template", name).Msg("could not parse template")
	}

	return tmpl
}

// loadTemplate loads a template from the file system
func parseCustomTemplate(file os.DirEntry, path string, partials []string) (*template.Template, error) {
	customFiles := []string{filepath.Join(path, file.Name())}

	for _, partial := range partials {
		customFiles = append(customFiles, filepath.Join(path, partial))
	}

	tmpl, err := template.New(file.Name()).
		Funcs(fm).
		ParseFiles(customFiles...)
	if err != nil {
		return nil, fmt.Errorf("could not parse template %q: %w", file.Name(), err)
	}

	return tmpl, nil
}

// Render returns the text and html executed templates for the specified name and data
func Render(name string, data interface{}) (text, html string, err error) {
	if text, err = render(name+".txt", data); err != nil {
		return
	}

	if html, err = render(name+".html", data); err != nil {
		return
	}

	return
}

// render the provided template with the data
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
