package Home

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/golangast/go_sapper/go/Autho"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

// tpl holds all parsed templates
var tpl *template.Template

// index serves the index html file
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w.Header())
	fmt.Println("started")
	fmt.Println("Post is chosen")
	fmt.Println(r.Header.Get("Origin"))
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080/home")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	fmt.Println("index started ")

	session, _ := Autho.Store.Get(r, "cookie-name")
	fmt.Println(session)
	user := Autho.CheckUser(session)
	fmt.Println(user)

	if r.Method == "GET" {

		fmt.Println("this is the user: ", user)
		tpl.ExecuteTemplate(w, "index.html", user)
	}

	if r.Method == "POST" {

		fmt.Println("this is the user: ", user)
		tpl.ExecuteTemplate(w, "index.html", user)
	}

}

// logout revokes authentication for a user
func logout(w http.ResponseWriter, r *http.Request) {
	session, err := Autho.Store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user"] = Autho.User{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
