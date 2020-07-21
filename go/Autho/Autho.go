package Autho

import (
	"encoding/gob"
	"fmt"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	. "github.com/logrusorgru/aurora"

	//import
	GetAllUsers "github.com/golangast/go_sapper/go/DB"
)

//initialize for autho
func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)
	Store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)
	Store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}
	gob.Register(User{})
}

// store will hold all session data
var Store *sessions.CookieStore

type User struct {
	Username      string
	Authenticated bool
}

// on error returns an empty user
func CheckUser(s *sessions.Session) User {
	val := s.Values["user"]
	var user = User{}

	//prints the cookie data
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

	//checking user autho
	user, ok := val.(User)
	if !ok {
		fmt.Println("user not autho")
		return User{Authenticated: false}
	}
	return user
}

//where the actual Autho begins
func GetAuthoUser(Username string, pass string) *User {
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
