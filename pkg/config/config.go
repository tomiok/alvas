package config

import "html/template"

var AppCfg *AppConfig

type AppConfig struct {
	Cache map[string]*template.Template
}

func Init(fn func() (map[string]*template.Template, error)) {
	m, _ := fn()

	AppCfg = &AppConfig{
		Cache: m,
	}
}
