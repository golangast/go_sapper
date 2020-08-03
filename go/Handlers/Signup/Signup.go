package Home

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/golangast/go_sapper/go/Autho"
	Header "github.com/golangast/go_sapper/go/Handlers/Headers"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

// tpl holds all parsed templates
var tpl *template.Template

// Signup returns signup.html
func Signup(w http.ResponseWriter, r *http.Request) {
	Header.Headers(w, r)

	session, _ := Autho.Store.Get(r, "cookie-name")
	fmt.Println(session)
	user := Autho.CheckUser(session)
	fmt.Println(user)

	if r.Method == "GET" {

		fmt.Println("this is the user: ", user)
		tpl.ExecuteTemplate(w, "signup.html", user)
	}

	if r.Method == "POST" {

		fmt.Println("this is the user: ", user)
		tpl.ExecuteTemplate(w, "signup.html", user)
	}

}
