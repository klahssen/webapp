package middlewares

import (
	"net/http"
	"time"

	"github.com/klahssen/webapp/pkg/log"
)

//Log middleware logs new requests and how long they take
func Log(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()
		next.ServeHTTP(w, r)
		log.Infof("%s %s in %s", r.Method, r.URL.Path, time.Since(t0))
	}
	return http.HandlerFunc(fn)
}
