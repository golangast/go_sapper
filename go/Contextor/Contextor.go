package Contextor

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	. "github.com/logrusorgru/aurora"
	"gitlab.com/zendrulat123/gow/Handlers/Clients"
)

var err error

//used for printing time of request
var Start = time.Now()
var Durations = time.Now().Sub(Start)

//getting context
type Contexter struct {
	M      string
	S      int
	Co     string
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
			S:      http.StatusBadRequest,
			U:      r.WithContext(ctx).URL,
			P:      r.WithContext(ctx).Proto,
			B:      r.WithContext(ctx).Body,
			Host:   r.WithContext(ctx).Host,
			Form:   r.WithContext(ctx).Form,
			Cancel: r.WithContext(ctx).Cancel,
			R:      r.WithContext(ctx).Response,
			D:      Duration,
			H:      r.WithContext(ctx).Header,
			I:      Clients.ReadUserIP(r),
		}

		fmt.Println(Blue("/ʕ◔ϖ◔ʔ/````````````````````````````````````````````"))
		fmt.Printf("Method:%s\n - Status:%d\n - URL:%s - Body:%v\n - Host:%s\n - Form:%v\n - Cancel:%d\n - Response:%d\n - Dur:%02d-00:00\n - Cache-Control:%s - Accept:%s\n - IP:%s\n",
			Cyan(CC.M),
			Brown(CC.S),
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

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
