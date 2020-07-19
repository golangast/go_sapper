package Success

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/golangast/go_sapper/go/Autho"
)

// tpl holds all parsed templates
var tpl *template.Template

// secret displays the secret message for authorized users
func Success(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w.Header())
	fmt.Println("started")
	fmt.Println("Post is chosen")
	fmt.Println(r.Header.Get("Origin"))
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080/secret")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	if r.Method == "GET" {
		fmt.Println(r.Method)
		fmt.Println("you reached the secret get")
		session, _ := Autho.Store.Get(r, "cookie-name")
		fmt.Println(session)
		user := Autho.CheckUser(session)
		fmt.Println(user)
		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			fmt.Println("unautho")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		fmt.Println("about to send secret template")
		tpl.ExecuteTemplate(w, "secret.html", nil)
	}

	if r.Method == "POST" {
		fmt.Println(r.Method)
		fmt.Println("you reached the secret post")
		session, _ := Autho.Store.Get(r, "cookie-name")
		fmt.Println(session)
		user := Autho.CheckUser(session)
		fmt.Println(user)
		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			fmt.Println("unautho")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		fmt.Println("about to send secret template")
		tpl.ExecuteTemplate(w, "secret.html", nil)
	}

}
