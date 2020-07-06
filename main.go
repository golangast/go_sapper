package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	//only 3rd parties

	. "github.com/logrusorgru/aurora"
	"github.com/rs/cors"
	"gitlab.com/zendrulat123/gow/Handlers/Clients"

	API "github.com/golangast/go_sapper/go/Handlers/API"
	DB "github.com/golangast/go_sapper/go/Handlers/Form"
)

const (
	CONN_HOST      = "localhost"
	CONN_PORT      = "8080"
	ADMIN_USER     = "admin"
	ADMIN_PASSWORD = "admin"
)

type Login struct {
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pass"`
}

var login []Login

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

func spaFileServeFunc(dir string) func(http.ResponseWriter, *http.Request) {
	fileServer := http.FileServer(http.Dir(dir))
	return func(w http.ResponseWriter, r *http.Request) {
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

//end of spa handler code
func main() {

	mux := http.NewServeMux() //used for cors

	//uses the autho handler and wraps the public one
	//https://subscription.packtpub.com/book/web_development/9781787286740/1/ch01lvl1sec13/implementing-basic-authentication-on-a-simple-http-server
	http.HandleFunc("/show", BasicAuth(helloWorld, "Please enter your username and password"))
	mux.HandleFunc("/post", DB.POST)
	mux.HandleFunc("/api", API.GET)
	//uses the autho handler and wraps the public one
	mux.HandleFunc("/", spaFileServeFunc("public"))
	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(CONN_HOST+":"+CONN_PORT, AddContext(c, handler)))
	//log.Fatal(http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
func BasicAuth(handler http.HandlerFunc, realm string) http.HandlerFunc {
	fmt.Println("autho begin")
	return func(w http.ResponseWriter, r *http.Request) {
		//checking

		// This is a dummy check for credentials.
		//go get html file
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
		// Read cookie
		cookie, err := r.Cookie("mine")
		if err != nil {
			fmt.Printf("Cant find cookie :/\r\n")
			return
		}

		fmt.Printf("%s=%s\r\n", cookie.Name, cookie.Value)

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
			//	fmt.Println("database ", username, password)
			login = append(login, Login{Username: username, Email: email, Password: password})
			//	fmt.Println("before marshal ", login)

		}

		// user, pass, ok := r.BasicAuth()
		// //if not autho
		// if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(ADMIN_USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(ADMIN_PASSWORD)) != 1 {
		// 	w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
		// 	w.WriteHeader(401)
		// 	w.Write([]byte("You are Unauthorized to access the application.\n"))
		// 	return
		// }
		// var users = user
		// for _, value := range login {

		// 	fmt.Println(login)
		// 	if user == string(login. {

		// 	} else {
		// 		unauthorised(w)
		// 		return
		// 	}
		// }

		defer rows.Close()
		//	spew.Dump(login)
		for key, value := range login {
			fmt.Println(key, value.Email)

			if value.Email == cookie.Name {
				fmt.Println("yes found")
			}

		}
		handler(w, r)
	}
}
func unauthorised(rw http.ResponseWriter) {
	rw.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
	rw.WriteHeader(http.StatusUnauthorized)
}

var err error

//used for printing time of request
var Start = time.Now()
var Durations = time.Now().Sub(Start)

//getting context
type Contexter struct {
	M      string
	S      int
	U      *url.URL
	P      string
	B      io.ReadCloser
	Gb     func() (io.ReadCloser, error)
	Host   string
	Form   url.Values
	Cancel <-chan struct{}
	R      *http.Response
	H      http.Header
	D      time.Duration
	I      string
}

//used to shorten use of Contexter
var CC Contexter

//initializing context
func AddContext(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		Start := time.Now()
		Duration := time.Now().Sub(Start)

		CC = Contexter{
			M:      r.WithContext(ctx).Method,
			S:      http.StatusBadRequest,
			U:      r.WithContext(ctx).URL,
			P:      r.WithContext(ctx).Proto,
			B:      r.WithContext(ctx).Body,
			Host:   r.WithContext(ctx).Host,
			Form:   r.WithContext(ctx).Form,
			Cancel: r.WithContext(ctx).Cancel,
			R:      r.WithContext(ctx).Response,
			D:      Duration,
			H:      r.WithContext(ctx).Header,
			I:      Clients.ReadUserIP(r),
		}

		fmt.Println(Blue("/ʕ◔ϖ◔ʔ/````````````````````````````````````````````"))
		fmt.Printf("Method:%s\n - Status:%d\n - URL:%s - Body:%v\n - Host:%s\n - Form:%v\n - Cancel:%d\n - Response:%d\n - Dur:%02d-00:00\n - Cache-Control:%s - Accept:%s\n - IP:%s\n",
			Cyan(CC.M),
			Brown(CC.S),
			Red(CC.U),
			Blue(CC.B),
			Yellow(CC.Host),
			BgRed(CC.Form),
			BgGreen(CC.Cancel),
			BgBrown(CC.R),
			BgMagenta(CC.D),
			Red(CC.H.Get("Cache-Control")),
			Blue(CC.H.Get("Accept")),
			Yellow(CC.I),
		)
		//this is to spit out the ctx.header in pieces cause its exhaustive
		// for k, v := range CC.H {
		// 	fmt.Println("\n", k, v)
		// }
		cookie, _ := r.Cookie("username")

		if cookie != nil {
			//Add data to context
			ctx := context.WithValue(r.Context(), "Username", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {

			if err != nil {
				// Error occurred while parsing request body
				w.WriteHeader(http.StatusBadRequest)

				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
