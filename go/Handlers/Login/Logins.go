package Logins

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	//imported files
	Autho "github.com/golangast/go_sapper/go/Autho"
)
//template stuff
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
var tpl *template.Template

// login handler authenticates the user
func Logins(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login started")

	//checks is cookie is there
	session, err := Autho.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//grab the form data
	username := r.FormValue("username")
	pass := r.FormValue("code")
	if pass == "" {
		session.AddFlash("Must enter a code")
	}
	return
}

	//authorize
	user := Autho.GetAuthoUser(username, pass)
    //create session
	session.Values["user"] = user
	session.Values["authenticated"] = true
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("didnt make it to redirect")
		return
	}
	fmt.Println("Your request method:", r.Method)
	err = tpl.ExecuteTemplate(w, "success.html", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}


