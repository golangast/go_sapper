package Success

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/golangast/go_sapper/go/Autho"
	Header "github.com/golangast/go_sapper/go/Handlers/Headers"
)

// tpl holds all parsed templates
var tpl *template.Template

// secret displays the secret message for authorized users
func Success(w http.ResponseWriter, r *http.Request) {
	//checking and authorizing headers
	Header.Headers(w, r)

	if r.Method == "GET" {
		fmt.Println(r.Header.Get("Origin"))
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}
		w.Write([]byte("Received a GET request\n"))
		fmt.Println("reqeusted boyd ", r.Body)
		w.WriteHeader(http.StatusOK)
	}

	if r.Method == "POST" {
		fmt.Println(r.Method)

		session, err := Autho.Store.Get(r, "cookie-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user := Autho.CheckUser(session)
		fmt.Println(user)
		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			fmt.Println("unautho")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		if user.Authenticated == true {
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		tpl.ExecuteTemplate(w, "secret.html", nil)
	}

}
