package render

import (
	"bytes"
	"github.com/rs/zerolog/log"
	"github.com/tomiok/alvas/pkg/config"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func TemplateRender(w http.ResponseWriter, tmpl string) {
	t, ok := config.AppCfg.Cache[tmpl]

	if !ok {
		log.Fatal().Msg("cache is not working")
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, nil)

	if err != nil {
		log.Fatal().Msgf("cannot execute %s", err.Error())
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Error().Err(err)
	}
}

func TemplateRenderCache() (map[string]*template.Template, error) {
	pages, err := filepath.Glob("./pkg/templates/*.page.tmpl")
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

		matches, err := filepath.Glob("./pkg/templates/*.layout.tmpl")

		if err != nil {
			return templateCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./pkg/templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
			templateCache[name] = ts
		}
	}

	return templateCache, nil
}
