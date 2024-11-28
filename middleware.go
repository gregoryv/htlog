package htlog

import (
	"net/http"
	"net/url"
	"time"
)

type Middleware struct {
	// Println is used to print request method, path, response status
	// and duration
	Println func(...any)

	// Clean is used to hide query parameters if needed
	Clean func(u *url.URL) string
}

func (m *Middleware) Use(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// default status is 200, incase WriteHeader is not called
		rec := statusRecorder{w, 200}

		// do
		next.ServeHTTP(&rec, r)

		path := r.URL.Path
		if m.Clean != nil {
			path = m.Clean(r.URL)
		}
		if m.Println != nil {
			m.Println(r.Method, path, rec.status, time.Since(start))
		}
	}
}
