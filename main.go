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
	Logout "github.com/golangast/go_sapper/go/Handlers/Logout"
	Success "github.com/golangast/go_sapper/go/Handlers/Success"
	Spa "github.com/golangast/go_sapper/go/Spa"
)

func main() {
	// m := autocert.Manager{
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("www.checknu.de"),
	// 	Cache:      autocert.DirCache("/home/letsencrypt/"),
	// }https://medium.com/@ScullWM/golang-http-server-for-pro-69034c276355

	// server := &http.Server{
	// 	Addr: ":443",
	// 	TLSConfig: &tls.Config{
	// 		GetCertificate: m.GetCertificate,
	// 	},
	// }
	//err := server.ListenAndServeTLS("", "") if err != nil {         log.Fatal("ListenAndServe: ", err) }
	mux := http.NewServeMux()
	mux.HandleFunc("/post", DB.POST)
	mux.HandleFunc("/api", API.GET)
	mux.HandleFunc("/logout", Logout.Logout)
	mux.HandleFunc("/login", Logins.Logins)
	mux.HandleFunc("/home", Home.Home)
	mux.HandleFunc("/forbidden", Forbidden.Forbidden)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	mux.HandleFunc("/secret", Success.Success)
	//uses the autho handler and wraps the public one
	mux.HandleFunc("/", Spa.SpaFileServeFunc("public"))
	handler := cors.Default().Handler(mux)
	c := context.Background()
	log.Fatal(http.ListenAndServe(":8081", Contextor.AddContext(c, handler)))
	//log.Fatal(http.ListenAndServe(CONN_HOST+":"+CONN_PORT, Contextor.AddContext(c, handler)))
	//log.Fatal(http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil))
}
