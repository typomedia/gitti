package hdlr

import (
	"log"
	"net/http"
	"os"
)

func Logger(next http.Handler) http.Handler {
	// Create a logger that writes to standard output
	logger := log.New(os.Stdout, "", log.LstdFlags)
	// logger.SetOutput(io.MultiWriter(os.Stdout, file))

	// Set up a handler function that will log each incoming request in the common log format
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request in the common log format
		logger.Printf("%s [%s] \"%s %s %s\" %d\n",
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			r.Proto,
			r.Header.Get("User-Agent"),
			http.StatusOK,
		)

		// Write a response back to the client
		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
	})
}
