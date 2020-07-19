package main

import (
	"context"
	"log"
	"net/http"

	//only 3rd parties
	"github.com/rs/cors"

	//imported files
	Contextor "github.com/golangast/go_sapper/go/Contextor"
	API "github.com/golangast/go_sapper/go/Handlers/API"
	Forbidden "github.com/golangast/go_sapper/go/Handlers/Forbidden"
	DB "github.com/golangast/go_sapper/go/Handlers/Form"
	Home "github.com/golangast/go_sapper/go/Handlers/Home"
	Logins "github.com/golangast/go_sapper/go/Handlers/Login"
	Success "github.com/golangast/go_sapper/go/Handlers/Success"
	Spa "github.com/golangast/go_sapper/go/Spa"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/post", DB.POST)
	mux.HandleFunc("/api", API.GET)
	mux.HandleFunc("/login", Logins.Logins)
	mux.HandleFunc("/home", Home.Home)
	mux.HandleFunc("/forbidden", Forbidden.Forbidden)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	mux.HandleFunc("/secret", Success.Success)
	//uses the autho handler and wraps the public one
	mux.HandleFunc("/", Spa.SpaFileServeFunc("public"))
	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(CONN_HOST+":"+CONN_PORT, Contextor.AddContext(c, handler)))
	//log.Fatal(http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil))
}
