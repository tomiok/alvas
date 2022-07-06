package render

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/tomiok/alvas/pkg/config"
)

var functions = template.FuncMap{}

type TemplateData struct {
	CustomerName string
	AdminName    string
	Err          string
	CSRFToken    string
	Data         map[string]interface{}
	IsLoginReq   bool
	IsLogged     bool
}

func TemplateRender(w http.ResponseWriter, r *http.Request, tmpl string, td *TemplateData) {
	var t *template.Template
	if config.AppCfg.UseCache {
		var ok = true
		t, ok = config.AppCfg.Cache[tmpl]
		if !ok {
			log.Fatal().Msg("cache is not working")
		}
	} else {
		cache, err := TemplateRenderCache()

		if err != nil {
			log.Fatal().Msg("cache is not working")
		}

		t = cache[tmpl]
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, td)

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
