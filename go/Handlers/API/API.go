package API

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
	Username string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pass"`
}

func GET(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w.Header())
	fmt.Println("started")
	fmt.Println("Post is chosen")
	fmt.Println(r.Header.Get("Origin"))
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081/get")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	w.WriteHeader(http.StatusOK)

	//switch statement for get or post
	switch r.Method {

	case "GET":
		fmt.Println(r.Header.Get("Origin"))
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}
		//w.Write([]byte("Received a GET request\n"))
		fmt.Println("reqeusted boyd ", r.Body)

		//go get html file
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
		var login []Login

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
			fmt.Println("before marshal ", login)

		}
		json.NewEncoder(w).Encode(login) //remember to encode it

		defer rows.Close()
		w.Header().Set("Content-Type", "application/json")

	case "POST":
		fmt.Println(r.Header.Get("Origin"))

		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", reqBody)
		w.Write([]byte("Received a POST request\n"))

		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}
