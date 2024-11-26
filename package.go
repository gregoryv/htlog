package htlog

import (
	"log"
	"net/http"
	"net/url"
)

// Default logs all requests using log.Println with basic cleanup of
// query parameters considered secret.
func Default(next http.Handler) http.HandlerFunc {
	m := &Middleware{
		Println: log.Println,
		Clean: QueryHide(
			"access_token",
			"password",
			"secret",
		),
	}
	return m.Use(next)
}

// QueryHide returns a cleanup func for use in [Middleware].
func QueryHide(words ...string) func(u *url.URL) string {
	return func(u *url.URL) string {
		query := u.Query()
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
