package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	contx "github.com/klahssen/webapp/pkg/context"
	"github.com/klahssen/webapp/pkg/json/format"
)

type ctxKey string

//TokenFromHeader checks if Authorization header is set
func TokenFromHeader(next http.Handler) http.Handler {
	t0 := time.Now()
	fn := func(w http.ResponseWriter, r *http.Request) {
		token, err := getAuthHeader(r)
		if err != nil {
			format.WriteResponse(w, http.StatusUnauthorized, err, nil, t0)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, contx.JwtToken, token)
		r = r.WithContext(ctx)
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
