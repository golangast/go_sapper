package Forbidden

import (
	"net/http"
	"text/template"

	"github.com/golangast/go_sapper/go/Autho"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

// tpl holds all parsed templates
var tpl *template.Template

func Forbidden(w http.ResponseWriter, r *http.Request) {
	session, err := Autho.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	flashMessages := session.Flashes()
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "forbidden.html", flashMessages)
}
