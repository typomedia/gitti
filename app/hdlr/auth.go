package hdlr

import (
	"github.com/spf13/viper"
	"github.com/typomedia/gitti/app/msg"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if !validate(token) {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("Unauthorized"))
			msg.Check(err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func validate(token string) bool {
	return token == viper.GetString("auth.token")
}
