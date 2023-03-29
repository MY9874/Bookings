// this configuration file might be accessed from any part of the application
// avoid importing extra packages here to avoid import package cycle, which will make it can't compile
// config will be imported by other files but won't import any other files
package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/moyu/bookings/internal/models"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
}
