package Post

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	Header "github.com/golangast/go_sapper/go/Handlers/Headers"
)

func UnmarshalLogin(data []byte) (Login, error) {
	var r Login
	fmt.Print("starting unmarshal")
	err := json.Unmarshal(data, &r)
	fmt.Print("is starting", data)

	return r, err
}

func (r *Login) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Login struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

func (p Login) Sanitize() {
	p.Name = html.EscapeString(p.Name)
	p.Email = html.EscapeString(p.Email)
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

func POST(w http.ResponseWriter, r *http.Request) {

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
		fmt.Println(r.Header.Get("Origin"))

		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", reqBody)
		w.Write([]byte("Received a POST request\n"))
		fmt.Println("reqeusted boyd ", r.Body)
		l, err := UnmarshalLogin(reqBody)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("opening database")

		//database beginsssssss
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

		// query
		stmt, err := db.Prepare("INSERT INTO users(name, email, pass) VALUES(?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}

		userstemp := Data{L: Login{Name: l.Name, Email: l.Email, Pass: l.Pass}}
		fmt.Println(userstemp)

		u := userstemp
		s, err := Save(u)
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec(s.L.Name, s.L.Email, s.L.Pass)
		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
		fmt.Println("reached query")

		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}
