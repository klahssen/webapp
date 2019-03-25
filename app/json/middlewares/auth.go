package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	contx "github.com/klahssen/webapp/pkg/context"
)

type ctxKey string

//AuthHeader checks if Authorization header is set
func AuthHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token, err := getAuthHeader(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized: %s\n", err.Error())
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, contx.JwtToken, token)
		ctx = context.WithValue(ctx, contx.ReqTime, token)
		r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func getAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("missing 'Authorization' header")
	}
	if !strings.HasPrefix(authHeader, "Bearer") {
		return "", fmt.Errorf("must satisfy Authorization Bearer scheme")
	}
	token := strings.TrimSpace(strings.TrimLeft(authHeader, "Bearer "))
	return token, nil
}
