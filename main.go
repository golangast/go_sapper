package main

import (
	"context"
	"database/sql"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"text/template"
	"time"

	//only 3rd parties
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	. "github.com/logrusorgru/aurora"
	"github.com/rs/cors"
	"gitlab.com/zendrulat123/gow/Handlers/Clients"

	//imported files
	API "github.com/golangast/go_sapper/go/Handlers/API"
	DB "github.com/golangast/go_sapper/go/Handlers/Form"
	DBAll "github.com/golangast/go_sapper/go/DB"
	
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Login struct {
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pass"`
}

// User holds a users account information https://curtisvermeeren.github.io/2018/05/13/Golang-Gorilla-Sessions
type User struct {
	Username      string
	Authenticated bool
}

// store will hold all session data
var store *sessions.CookieStore

// tpl holds all parsed templates
var tpl *template.Template

var login []Login

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(User{})

	tpl = template.Must(template.ParseGlob("templates/*.html"))

}

// func final(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Executing finalHandler")
// 	w.Header().Set("Content-Type", "text/html")
// 	tpl.ExecuteTemplate(w, "next.html", nil)
// }
func main() {

	mux := http.NewServeMux() //used for cors

	//uses the autho handler and wraps the public one
	//https://subscription.packtpub.com/book/web_development/9781787286740/1/ch01lvl1sec13/implementing-basic-authentication-on-a-simple-http-server
	mux.HandleFunc("/post", DB.POST)
	mux.HandleFunc("/api", API.GET)
	mux.HandleFunc("/login", logins)
	mux.HandleFunc("/home", index)

	// finalHandler := http.HandlerFunc(final)
	// mux.Handle("/next", exampleMiddleware(finalHandler))
	mux.HandleFunc("/forbidden", forbidden)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	mux.HandleFunc("/secret", secret)
	//uses the autho handler and wraps the public one
	mux.HandleFunc("/", SpaFileServeFunc("public"))

	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(CONN_HOST+":"+CONN_PORT, AddContext(c, handler)))
	//log.Fatal(http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil))
}

func SpaFileServeFunc(dir string) func(http.ResponseWriter, *http.Request) {
	fileServer := http.FileServer(http.Dir(dir))
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")
		fmt.Println(session)
		user := getUser(session)
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

var err error

//used for printing time of request
var Start = time.Now()
var Durations = time.Now().Sub(Start)

//getting context
type Contexter struct {
	M      string
	S      int
	Co     string
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

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

// index serves the index html file
func index(w http.ResponseWriter, r *http.Request) {
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

	session, _ := store.Get(r, "cookie-name")
	fmt.Println(session)
	user := getUser(session)
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
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user"] = User{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// secret displays the secret message for authorized users
func secret(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w.Header())
	fmt.Println("started")
	fmt.Println("Post is chosen")
	fmt.Println(r.Header.Get("Origin"))
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080/secret")
	//w.Header().Set("Access-Control-Allow-Origin", "http://b6f2509b93bd.ngrok.io")

	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	if r.Method == "GET" {
		fmt.Println(r.Method)
		fmt.Println("you reached the secret get")
		session, _ := store.Get(r, "cookie-name")
		fmt.Println(session)
		user := getUser(session)
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
		session, _ := store.Get(r, "cookie-name")
		fmt.Println(session)
		user := getUser(session)
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

func forbidden(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	flashMessages := session.Flashes()
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "forbidden.html", flashMessages)
}

// getUser returns a user from session s
// on error returns an empty user
func getUser(s *sessions.Session) User {
	val := s.Values["user"]
	var user = User{}

	fmt.Println(BgRed("/ʕ◔ϖ◔ʔ/ Cookies````````````````````````````````````````````"))
	fmt.Printf("Options:%v\n - Values:%v\n - Path:%s\n - ID:%v\n - Name:%d\n - IsNew:%v\n - Domain:%s\n - MaxAge:%d\n - User:%v/n",
		Cyan(s.Options),
		Brown(s.Values),
		Red(s.Options.Path),
		Blue(s.ID),
		Yellow(s.Name),
		BgRed(s.IsNew),
		BgGreen(s.Options.Domain),
		BgBrown(s.Options.MaxAge),
		BgMagenta(user))

	user, ok := val.(User)
	if !ok {
		fmt.Println("user not autho")
		return User{Authenticated: false}
	}
	return user
}

func getDBUser(Username string, pass string) *User {

	fmt.Println("autho begin")

	fmt.Println("return request")

	//begininig authoa
	login := DBAll()
	//begin scanning
	for key, value := range login {
		fmt.Println(key, " username is: ", value.Username, "password is: ", value.Password)
		//comparing usernames
		if value.Username != Username {
			fmt.Println("user not found")
		} else {
			fmt.Println("user found! ", value.Username)

			//check passwords
			if pass != value.Password {
				fmt.Println("password not found")
			} else {
				fmt.Println("password found! ", value.Password)

				user := &User{
					Username:      Username,
					Authenticated: true,
				}
				fmt.Println("right before return ", user)
				return user

			}
		}

	}
	user := &User{
		Username:      Username,
		Authenticated: false,
	}
	return user

}
func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		next.ServeHTTP(w, r)
	})
}

// login authenticates the user
func logins(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login started")
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username := r.FormValue("username")
	pass := r.FormValue("code")
	user := getDBUser(username, pass)
	fmt.Println(user, username, pass)

	session.Values["user"] = user

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("didnt make it to redirect")
		return
	}
	fmt.Println("you are authoed and now being redirected to secret", user)

	session.Values["authenticated"] = true
	session.Save(r, w)
	fmt.Println("Your request method:", r.Method)
	err = tpl.ExecuteTemplate(w, "secret.html", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}

	if r.FormValue("code") == "" {
		session.AddFlash("Must enter a code")
	}

	return

}
