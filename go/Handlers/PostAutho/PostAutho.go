package Post

import (
	//"bytes"
	//"compress/zlib"

	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	//imported files
	Autho "github.com/golangast/go_sapper/go/Autho"
	Header "github.com/golangast/go_sapper/go/Handlers/Headers"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

// tpl holds all parsed templates
var tpl *template.Template

func UnmarshalLogin(data []byte) (Login, error) {
	var r Login
	fmt.Println("starting unmarshal ", string(data))
	err := json.Unmarshal(data, &r)
	fmt.Println("is starting", string(data))

	return r, err
}

func (r *Login) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Login struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Pass  string `json:"pass"`
}

func (p Login) Sanitize() {
	p.Email = html.EscapeString(p.Email)
	p.Name = html.EscapeString(p.Name)
	p.Pass = html.EscapeString(p.Pass)

}

type Data struct {
	L Login
	S Sanitizer
}
type Sanitizer interface {
	Sanitize()
}

func Save(l Data) (Data, error) {
	var err error

	// type assertion for Sanitizer (could also use a type switch)
	s, ok := l.S.(Sanitizer)

	if !ok {
		if err != nil {
			log.Fatal(err)
		}
		// ... save without sanitization
		return l, err
	}

	s.Sanitize()
	return l, err
}

type ContactDetails struct {
	Email string
	Name  string
	Pass  string
}

// PostAutho saves in the database
func PostAutho(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AuthoPost started")
	//checking and authorizing headers
	Header.Headers(w, r)

	switch r.Method {
	case "GET":
		fmt.Println(r.Header.Get("Origin"))
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}
		w.Write([]byte("Received a GET request\n"))
		fmt.Println("reqeusted boyd ", r.Body)
		w.WriteHeader(http.StatusOK)

	case "POST":

		fmt.Println("starting PostAutho , ", r.Header.Get("Origin"))
		details := ContactDetails{
			Email: r.FormValue("email"),
			Name:  r.FormValue("name"),
			Pass:  r.FormValue("pass"),
		}

		fmt.Println("opening database")

		//database beginsssssss
		db, err := sql.Open("mysql", "root:@/user")
		if err != nil {
			fmt.Println("there was an error after opening db")
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

		// query
		stmt, err := db.Prepare("INSERT INTO users(email, name, pass) VALUES(?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}

		userstemp := Data{L: Login{Email: details.Email, Name: details.Name, Pass: details.Pass}}
		fmt.Println(userstemp)

		u := userstemp
		s, err := Save(u)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec(s.L.Email, s.L.Name, s.L.Pass)
		if err != nil {
			log.Fatal(err)
		}

		lastID, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
		fmt.Println("reached query")

		//w.WriteHeader(http.StatusOK)
		fmt.Println("login started")
		Autho.StartCookie()
		//checks is cookie is there
		session, err := Autho.Store.Get(r, "cookie-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//GetAuthoUser checks username/pass if correct returns authorized user
		user := Autho.GetAuthoUser(details.Name, details.Pass)
		//create session
		session.Values["user"] = user
		if user.Authenticated == true {
			fmt.Println("authorized sessions started")
			session.Values["authenticated"] = true
		} else {
			session.Values["authenticated"] = false
		}

		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("didnt make it to redirect")
			return
		}

		tpl.ExecuteTemplate(w, "success.html", nil)

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}
