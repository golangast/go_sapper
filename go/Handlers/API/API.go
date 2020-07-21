package API

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	//import
	GetAllUsers "github.com/golangast/go_sapper/go/DB"
	Header "github.com/golangast/go_sapper/go/Handlers/Headers"
)

//used to unmarshal data
func UnmarshalLogin(data []byte) (Login, error) {
	var r Login
	fmt.Print("starting unmarshal")
	err := json.Unmarshal(data, &r)
	fmt.Print("is starting", data)

	return r, err
}

//needed to add method on package type
type Login GetAllUsers.Login

func (r *Login) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func GET(w http.ResponseWriter, r *http.Request) {

	//checking and authorizing headers
	Header.Headers(w, r)

	//switch statement for get or post
	switch r.Method {

	case "GET":
		fmt.Println(r.Header.Get("Origin"))
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}
		fmt.Println("reqeusted boyd ", r.Body)

		//get all users
		login := GetAllUsers.GetAllUsers()

		//turn into json
		json.NewEncoder(w).Encode(login)
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
