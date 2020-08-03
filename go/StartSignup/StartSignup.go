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

// Logins handler authenticates the user
func Logins(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login started")
	Autho.StartCookie()
	//checks is cookie is there
	session, err := Autho.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//grab the form data
	username := r.FormValue("username")
	pass := r.FormValue("password")
	email := r.FormValue("email")

	if pass == "" {
		session.AddFlash("Must enter a password")
	}

	//GetAuthoUser checks username/pass if correct returns authorized user
	user := Autho.SignupUser(username, pass, email)

	//create session
	session.Values["user"] = user
	if user.Authenticated == true {
		fmt.Println("authorized sessions started")
		session.Values["authenticated"] = true
	} else {
		session.Values["authenticated"] = false
	}

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

}
