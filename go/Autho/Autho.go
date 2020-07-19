package Autho

import (
	"encoding/gob"
	"fmt"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	. "github.com/logrusorgru/aurora"
)

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
