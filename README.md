[gregoryv/htlog](https://pkg.go.dev/github.com/gregoryv/htlog)
provides a log middleware for Go http.Handler

## Quick start

     go get github.com/gregoryv/htlog
   
then use it

     var h http.Handler
   
     // use default middleware
     router := htlog.Default(h)

