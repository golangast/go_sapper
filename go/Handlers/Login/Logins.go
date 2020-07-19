package Logins

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	//3rd party

	//imported files
	Autho "github.com/golangast/go_sapper/go/Autho"
	GetAllUsers "github.com/golangast/go_sapper/go/DB"
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

// tpl holds all parsed templates
var tpl *template.Template

// login authenticates the user
func Logins(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login started")
	session, err := Autho.Store.Get(r, "cookie-name")
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

func getDBUser(Username string, pass string) *Autho.User {
	fmt.Println("autho begin")
	fmt.Println("return request")
	//begininig authoa
	login := GetAllUsers.GetAllUsers()
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
				user := &Autho.User{
					Username:      Username,
					Authenticated: true,
				}
				fmt.Println("right before return ", user)
				return user
			}
		}
	}
	user := &Autho.User{
		Username:      Username,
		Authenticated: false,
	}
	return user

}
