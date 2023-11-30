package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

type Appconfig struct {
	TemplateCache map[string]*template.Template
	InProduction bool
	Session *scs.SessionManager
}