package config

import "text/template"

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}

func (a *AppConfig) SetUseCache(flag bool) {
	a.UseCache = flag
}
