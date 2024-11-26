package htlog_test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
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

	r := httptest.NewRequest("GET", "/?access_token=SECRET", http.NoBody)
	_ = recordResp(h, r)
	got := buf.String()
	if err := contains(got, "GET", "/", "µs", "token=..."); err != nil {
		t.Error(got, err)
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
