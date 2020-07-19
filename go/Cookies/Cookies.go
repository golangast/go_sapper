package Cookies

import (
	"fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

// Hash keys should be at least 32 bytes long
var hashKey = []byte("very-secret")

// Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
// Shorter keys may weaken the encryption used.
var blockKey = []byte("a-lot-secret")
var S = securecookie.New(hashKey, blockKey)

func ReadingCookie(w http.ResponseWriter, r *http.Request) string {
	if cookie, err := r.Cookie("cookie-name"); err == nil {
		value := make(map[string]string)
		if err = S.Decode("cookie-name", cookie.Value, &value); err == nil {
			fmt.Fprintf(w, "The value of foo is %q", value["foo"])
		}
		return cookie.Value
	}
	return "no cookie"
}

func SetCookieHandler(s *securecookie.SecureCookie, w http.ResponseWriter, r *http.Request) {
	value := map[string]string{
		"foo": "bar",
	}
	if encoded, err := s.Encode("cookie-name", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie-name",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}
