package htlog_test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gregoryv/htlog"
)

func Test_Default(t *testing.T) {
	var router http.Handler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			fmt.Fprint(w, "hello")
		},
	)
	h := htlog.Default(router)
	var buf bytes.Buffer
	log.SetOutput(&buf)

	r := httptest.NewRequest("GET", "/?password=SECRET", http.NoBody)
	_ = recordResp(h, r)
	got := buf.String()
	err := contains(got, "GET", "/", "µs", "password=...")
	if err != nil {
		t.Error(got, err)
	}
}

func Test_QueryHide(t *testing.T) {
	r := httptest.NewRequest("GET", "/some/secret", http.NoBody)
	got := htlog.QueryHide()(r.URL)
	if err := contains(got, "/some/secret"); err != nil {
		t.Error(err)
	}
}

func Test_DefaultClean(t *testing.T) {
	cases := []string{
		"access_token",
		"secret",
		"password",
	}
	for _, c := range cases {
		t.Run(c, func(t *testing.T) {
			u, _ := url.Parse("http://example.com/?" + c + "=SECRET")
			got := htlog.DefaultClean(u)
			if err := contains(got, "SECRET"); err == nil {
				t.Errorf("expected %q value to be hidden", c)
			}
		})
	}
}

func recordResp(h http.Handler, r *http.Request) *http.Response {
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w.Result()
}

func contains(got string, expect ...string) error {
	var miss []string
	for _, exp := range expect {
		if !strings.Contains(got, exp) {
			miss = append(miss, exp)
		}
	}
	if len(miss) > 0 {
		return fmt.Errorf("missing %q", miss)
	}
	return nil
}
