package htlog

import (
	"log"
	"net/http"
	"net/url"
)

// Default middleware logging all requests using log.Println.
func Default(next http.Handler) http.HandlerFunc {
	m := &Middleware{
		Println: log.Println,
		Clean:   DefaultClean,
	}
	return m.Use(next)
}

// DefaultClean is used by Default constructor
var DefaultClean = QueryHide("access_token", "password", "secret")

// QueryHide returns a cleanup func for use in [Middleware].  Each
// word matching a query parameter is replaced with a "..." value.
func QueryHide(words ...string) func(u *url.URL) string {
	return func(u *url.URL) string {
		query := u.Query()
		if len(query) == 0 {
			return u.Path
		}
		// hide query arguments considered secret
		for _, k := range words {
			if query.Has(k) {
				query.Set(k, "...")
			}
		}
		// include query in path
		path := u.Path
		if v := query.Encode(); v != "" {
			path += "?" + v
		}
		return path
	}
}
