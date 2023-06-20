package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool // use cache or not
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger // save logs
	InProduction  bool
	Session       *scs.SessionManager
}
