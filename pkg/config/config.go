package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application configuration
type AppConfig struct {
	TemplateCache map[string]*template.Template
	Logger        *log.Logger
	UseCache      bool
}
