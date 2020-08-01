package Cookies

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"
)

const cookieName = "mycookiename"

var hashKey = []byte(securecookie.GenerateRandomKey(32))
var blockKey = []byte(securecookie.GenerateRandomKey(32))

var sc = securecookie.New(hashKey, blockKey)

func checkAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if ReadCookieHandler(r) {
			h.ServeHTTP(w, r)
		} else {
			log.Printf("Not authorized %s", 401)
		}

		h.ServeHTTP(w, r)
	}
}

func identifyUserFromForm(r *http.Request) User {
	err := r.ParseForm()
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	du := User{}

	u, err := du.FindOneUserByEmail(email)

	if err != nil {
		log.Fatal(err)
	} else {
		passwordPlainText := []byte(fmt.Sprintf("%s%s", u.Password_salt, password))

		err = bcrypt.CompareHashAndPassword([]byte(u.Password_hash), []byte(passwordPlainText))
		if err == nil {
			log.Printf("user %s auth ok\n", u.Email)
			return u
		}
	}

	return du
}

//GenerateBcryptHash Generates the hash
func GenerateBcryptHash(s string, p string) ([]byte, error) {
	fmt.Printf("user auth: %s %s", s, p)
	password := []byte(fmt.Sprintf("%s%s", s, p))
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//CompareBcryptHash compares the hashes
func CompareBcryptHash(hash []byte, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), hash)
}

//SetCookieHandler  set the cookie
func SetCookieHandler(w http.ResponseWriter, r *http.Request, u User) {
	value := map[string]string{
		"id": u.Id,
	}

	log.Printf("set cookie: %s\n", u.Id)
	if encoded, err := sc.Encode(cookieName, value); err == nil {
		cookie := &http.Cookie{
			Name:  cookieName,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

//ReadCookieHandler reads the cookie
func ReadCookieHandler(r *http.Request) bool {

	log.Printf("cookie name: %#s\n", cookieName)
	if cookie, err := r.Cookie(cookieName); err == nil {
		value := make(map[string]string)
		if err = sc.Decode(cookieName, cookie.Value, &value); err == nil {

			log.Printf("cookie: %#s\n", cookie)
			return true
		}
		log.Printf("cookie: %#s\n", cookie)
	}
	return false
}
