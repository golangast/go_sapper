package Spa

import (
	"fmt"
	"net/http"

	"github.com/golangast/go_sapper/go/Autho"
)

//used to let through the sveltejs
func SpaFileServeFunc(dir string) func(http.ResponseWriter, *http.Request) {
	//take from the sveltejs files
	fileServer := http.FileServer(http.Dir(dir))
	return func(w http.ResponseWriter, r *http.Request) {
		//autho
		session, _ := Autho.Store.Get(r, "cookie-name")
		fmt.Println(session)
		user := Autho.CheckUser(session)
		fmt.Println(user)
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			fmt.Println("unautho")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		//check if okay
		wt := &intercept404{ResponseWriter: w}
		fileServer.ServeHTTP(wt, r)
		fmt.Println(w.Header())
		if wt.statusCode == http.StatusNotFound {
			r.URL.Path = "/"
			w.Header().Set("Content-Type", "text/html")
			fileServer.ServeHTTP(w, r)
		}
	}
}

//spa handler code
type intercept404 struct {
	http.ResponseWriter
	statusCode int
}
