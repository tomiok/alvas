package render

import (
	"github.com/rs/zerolog/log"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func TemplateRender(w http.ResponseWriter, tmpl string) {
	_, err := TemplateRenderCache(w)

	if err != nil {
		log.Error().Msgf("cannot render: %s", err.Error())
	}

	parsed, err := template.ParseFiles("../templates/" + tmpl)

	if err != nil {
		log.Error().Err(err)
		return
	}

	err = parsed.Execute(w, nil)

	if err != nil {
		log.Error().Err(err)
		return
	}
}

func TemplateRenderCache(w http.ResponseWriter) (map[string]*template.Template, error) {
	pages, err := filepath.Glob("../templates/*.page.tmpl")
	var templateCache = make(map[string]*template.Template)

	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("../templates/*.layout.tmpl")

		if err != nil {
			return templateCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
			templateCache[name] = ts
		}
	}

	return templateCache, nil
}
