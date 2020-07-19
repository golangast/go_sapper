package Spa

import (
	"fmt"
	"net/http"

	"github.com/golangast/go_sapper/go/Autho"
)

func SpaFileServeFunc(dir string) func(http.ResponseWriter, *http.Request) {
	fileServer := http.FileServer(http.Dir(dir))
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Autho.Store.Get(r, "cookie-name")
		fmt.Println(session)
		user := Autho.CheckUser(session)
		fmt.Println(user)
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			fmt.Println("unautho")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		wt := &intercept404{ResponseWriter: w}
		fileServer.ServeHTTP(wt, r)
		fmt.Println(w.Header())
		fmt.Println("started")
		fmt.Println("Post is chosen")
		fmt.Println(r.Header.Get("Origin"))

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
