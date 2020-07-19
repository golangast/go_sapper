package Login

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

//spa handler code
type intercept404 struct {
	http.ResponseWriter
	statusCode int
}

func (w *intercept404) Write(b []byte) (int, error) {
	if w.statusCode == http.StatusNotFound {
		return len(b), nil
	}
	if w.statusCode != 0 {
		w.WriteHeader(w.statusCode)
	}
	return w.ResponseWriter.Write(b)
}

func (w *intercept404) WriteHeader(statusCode int) {
	if statusCode >= 300 && statusCode < 400 {
		w.ResponseWriter.WriteHeader(statusCode)
		return
	}
	w.statusCode = statusCode
}

func SpaFileServeFunc(dir string) func(http.ResponseWriter, *http.Request) {
	fileServer := http.FileServer(http.Dir(dir))
	return func(w http.ResponseWriter, r *http.Request) {
		var hashKey = []byte("very-secret")
		var blockKey = []byte("a-lot-secret")
		var s = securecookie.New(hashKey, blockKey)

		ss := ReadingCookie(s, w, r)
		dlogin := GetData()

		//TODO COMPARE DATABASE LOGIN WITH COOKIE
		//TODO CHANGE COOKIE VALUE TO COMPARE TO DATABASE
		fmt.Println("start comparing")
		fmt.Println("cookie is ", ss)
		for key, value := range dlogin {
			fmt.Println("database is ", key, value.Email)
			//value is database data and s is cookie
			if value.Email == ss {
				fmt.Println("yes found")
			}
		}
		fmt.Println("done comparing")

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

func GetData() []Login {

	fmt.Println("autho begin")
	fmt.Println("db begin")
	db, err := sql.Open("mysql", "root:@/user")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("open ")
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ping ")
	}
	//opening database

	var (
		id       int
		email    string
		username string
		password string
	)
	i := 0

	//	rows, err := db.Query("select * from users where id = ?", 1)
	rows, err := db.Query("select * from users")
	for rows.Next() {
		err := rows.Scan(&id, &username, &email, &password)
		if err != nil {
			fmt.Println(err)
		} else {
			i++
			fmt.Println("scan ", i)
		}
		login = append(login, Login{Username: username, Email: email, Password: password})

	}
	defer rows.Close()
	fmt.Println("return request")

	return login
}
func unauthorised(rw http.ResponseWriter) {
	rw.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
	rw.WriteHeader(http.StatusUnauthorized)
}
