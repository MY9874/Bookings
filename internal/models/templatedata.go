package models

import "github.com/moyu/bookings/internal/forms"

//this file will never import other files, only being imported by others

// to store data that sent from handlers to templates, we do not know the type so need a struct
type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{} //we do not know the type, uncertain
	CSRFToken       string                 // a security token for every page, crossed site request forgery token
	Flash           string                 // a flash message
	Warning         string
	Error           string
	Form            *forms.Form
	IsAuthenticated int
}
