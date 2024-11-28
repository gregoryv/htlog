package htlog_test

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gregoryv/htlog"
)

func ExampleDefault() {
	http.Handle("/", htlog.UseDefault(serveHello()))
	http.ListenAndServe(":8080", nil)
}

func ExampleMiddleware() {
	// define a custom
	customLog := htlog.Middleware{
		Println: func(a ...any) {
			fmt.Println(a...)
		},
		Clean: func(u *url.URL) string {
			return u.Path
		},
	}

	http.Handle("/", customLog.Use(serveHello()))
	http.ListenAndServe(":8080", nil)
}

func serveHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	}
}
