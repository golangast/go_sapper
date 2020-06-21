package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	//only 3rd parties

	. "github.com/logrusorgru/aurora"
	"github.com/rs/cors"

	API "github.com/golangast/go_sapper/go/Handlers/API"
	DB "github.com/golangast/go_sapper/go/Handlers/Form"
)

type intercept404 struct {
	http.ResponseWriter
	statusCode int
}

func (w *intercept404) Write(b []byte) (int, error) {
	if w.statusCode == http.StatusNotFound {
		return len(b), nil
	}
	if w.statusCode != 0 {
		w.WriteHeader(w.statusCode)
	}
	return w.ResponseWriter.Write(b)
}

func (w *intercept404) WriteHeader(statusCode int) {
	if statusCode >= 300 && statusCode < 400 {
		w.ResponseWriter.WriteHeader(statusCode)
		return
	}
	w.statusCode = statusCode
}

func spaFileServeFunc(dir string) func(http.ResponseWriter, *http.Request) {

	fileServer := http.FileServer(http.Dir(dir))
	return func(w http.ResponseWriter, r *http.Request) {
		wt := &intercept404{ResponseWriter: w}
		fileServer.ServeHTTP(wt, r)
		fmt.Println(w.Header())
		fmt.Println("started")
		fmt.Println("Post is chosen")
		fmt.Println(r.Header.Get("Origin"))

		if wt.statusCode == http.StatusNotFound {
			r.URL.Path = "/"
			w.Header().Set("Content-Type", "text/html")
			fileServer.ServeHTTP(w, r)
		}
	}
}

func main() {
	mux := http.NewServeMux() //used for cors
	mux.HandleFunc("/", spaFileServeFunc("public"))
	mux.HandleFunc("/post", DB.POST)
	mux.HandleFunc("/api", API.GET)
	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(":8081", AddContext(c, handler)))

}

var err error

//used for printing time of request
var Start = time.Now()
var Durations = time.Now().Sub(Start)

//getting context
type Contexter struct {
	M      string
	U      *url.URL
	P      string
	B      io.ReadCloser
	Gb     func() (io.ReadCloser, error)
	Host   string
	Form   url.Values
	Cancel <-chan struct{}
	R      *http.Response
	H      http.Header
	D      time.Duration
	I      string
}

//used to shorten use of Contexter
var CC Contexter

//initializing context
func AddContext(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Start := time.Now()
		Duration := time.Now().Sub(Start)

		CC = Contexter{
			M:      r.WithContext(ctx).Method,
			U:      r.WithContext(ctx).URL,
			P:      r.WithContext(ctx).Proto,
			B:      r.WithContext(ctx).Body,
			Host:   r.WithContext(ctx).Host,
			Form:   r.WithContext(ctx).Form,
			Cancel: r.WithContext(ctx).Cancel,
			R:      r.WithContext(ctx).Response,
			D:      Duration,
			H:      r.WithContext(ctx).Header,
			//I:      Clients.ReadUserIP(r),
		}

		fmt.Println(Blue("/ʕ◔ϖ◔ʔ/````````````````````````````````````````````"))
		fmt.Printf("Method:%s\n - Status:%s\n - URL:%s - Body:%v\n - Host:%s\n - Form:%v\n - Cancel:%d\n - Response:%d\n - Dur:%02d-00:00\n - Cache-Control:%s - Accept:%s\n - IP:%s\n",
			Cyan(CC.M),
			Brown(r.Header.Get("X-Forwarded-Port")),
			Red(CC.U),
			Blue(CC.B),
			Yellow(CC.Host),
			BgRed(CC.Form),
			BgGreen(CC.Cancel),
			BgBrown(CC.R),
			BgMagenta(CC.D),
			Red(CC.H.Get("Cache-Control")),
			Blue(CC.H.Get("Accept")),
			Yellow(CC.I),
		)
		//this is to spit out the ctx.header in pieces cause its exhaustive
		// for k, v := range CC.H {
		// 	fmt.Println("\n", k, v)
		// }
		cookie, _ := r.Cookie("username")

		if cookie != nil {
			//Add data to context
			ctx := context.WithValue(r.Context(), "Username", cookie.Value)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {

			if err != nil {
				// Error occurred while parsing request body
				w.WriteHeader(http.StatusBadRequest)

				return
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
