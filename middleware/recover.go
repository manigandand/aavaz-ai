package middleware

import (
	"aavaz/errors"
	"aavaz/respond"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/go-chi/chi/middleware"
)

// Recoverer is a middleware that recovers from panics, logs the panic (and a
// backtrace), and returns a HTTP 500 (Internal Server Error) status if
// possible. Recoverer prints a request ID if one is provided.
func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				logEntry := middleware.GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rvr, debug.Stack())
				} else {
					fmt.Fprintf(os.Stderr, "Panic: %+v\n", rvr)
					debug.PrintStack()
				}

				// TODO: Let devloper know the panic
				// Log into sentry/..
				respond.Fail(w, errors.InternalServerStd())
				return
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
