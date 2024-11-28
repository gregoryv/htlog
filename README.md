[gregoryv/htlog](https://pkg.go.dev/github.com/gregoryv/htlog)
provides a log middleware for Go http.Handler

## Quick start

     go get github.com/gregoryv/htlog
   
then use it

     var h http.Handler
   
     // use default middleware
     router := htlog.UseDefault(h)

Depending on the configured Println a request can be logged as

    2024/11/28 20:25:00 GET /?password=... 200 7.544µs

