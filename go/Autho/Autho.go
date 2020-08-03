package Autho

import (
	"encoding/gob"
	"fmt"

	//3rd party
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	. "github.com/logrusorgru/aurora"

	//imports
	GetAllUsers "github.com/golangast/go_sapper/go/DB"
)

func init() {
	//gernerate the random keys
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	//start a session
	Store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	//store the session
	Store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	//check the user
	gob.Register(User{})
}

//Login for the data
type Login struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Pass  string `json:"pass"`
}

// Store will hold all session data
var Store *sessions.CookieStore

//User wrapper for user data
type User struct {
	Username      string
	Email         string
	Authenticated bool
}

var user = User{}

// CheckUser on error returns an empty user
func CheckUser(s *sessions.Session) User {

	val := s.Values["user"]

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

//GetAuthoUser starts autho
func GetAuthoUser(Username string, pass string) *User {
	//begininig authoa
	login := GetAllUsers.GetAllUsers()
	//begin scanning
	for _, value := range login {
		//comparing values
		success := CompareUser(value.Email, value.Name, Username, pass)
		if success == true {
			return &user
		} else {
			user := &User{
				Username:      Username,
				Authenticated: false,
			}
			spew.Dump(user)
		}
	}
	return &user
}

//SignupUser starts autho
func SignupUser(email string, Username string, pass string) *User {
	//begininig authoa
	login := GetAllUsers.GetAllUsers()
	//begin scanning
	for _, value := range login {
		//comparing values
		success := CompareSignupUser(value.Email, value.Name, value.Pass, Username, pass, email)
		if success == true {
			return &user
		} else {
			user := &User{
				Email:         email,
				Username:      Username,
				Authenticated: false,
			}
			spew.Dump(user)
		}
	}
	return &user
}

//CompareUser compares the user data
func CompareUser(u string, p string, Username string, pass string) bool {
	//comparing usernames
	if u != Username {
		fmt.Println("user not found")
	} else {
		fmt.Println("user found! ", u)
		//check passwords
		if pass != p {
			fmt.Println("password not found")
		} else {
			fmt.Println("password found! ", p)
			user := &User{

				Username:      u,
				Authenticated: true,
			}
			spew.Dump(user)
			return true
		}
	}
	return true
}

//CompareSignupUser compares the user data
func CompareSignupUser(u string, p string, e string, Username string, pass string, email string) bool {
	//comparing usernames
	if e != email {
		fmt.Println("email not found")
	} else {
		fmt.Println("email is found")
	}
	if u != Username {
		fmt.Println("user not found")
	} else {
		fmt.Println("user found! ", u)
		//check passwords
		if pass != p {
			fmt.Println("password not found")
		} else {
			fmt.Println("password found! ", p)
			user := &User{
				Email:         e,
				Username:      u,
				Authenticated: true,
			}
			spew.Dump(user)
			return true
		}
	}
	return true
}

//StartCookie logs the user in
func StartCookie() {

}
