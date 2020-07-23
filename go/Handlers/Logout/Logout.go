package Logins

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	
)

//template stuff
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

var tpl *template.Template

// login handler authenticates the user
func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout started")
c := &http.Cookie{
    Name:     "cookie-name",
    Value:    "",
    Path:     "/",
    MaxAge:   -1,
    HttpOnly: true,
}

http.SetCookie(w, c)
	
	fmt.Println("Your request method:", r.Method)
  err := tpl.ExecuteTemplate(w, "logout.html", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}

}
