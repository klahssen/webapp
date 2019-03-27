package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	dt "cloud.google.com/go/datastore"
	"github.com/go-zoo/bone"
	"github.com/justinas/alice"

	//"github.com/justinas/nosurf"
	"github.com/klahssen/webapp/pkg/json/api"
	"github.com/klahssen/webapp/pkg/json/middlewares"
	"github.com/klahssen/webapp/pkg/log"
	"github.com/klahssen/webapp/pkg/repos/gcloud/datastore"
	"github.com/klahssen/webapp/pkg/services/accounts"
	emails "github.com/klahssen/webapp/pkg/services/emails/local"
	"github.com/pkg/errors"
	"github.com/throttled/throttled"
	"github.com/throttled/throttled/store/memstore"
)

/*
WARNING:
The default http server has NO TIMEOUTS
*/

func main() {
	port := "8080"
	pport := flag.String("p", "", "port on which the server will listen")
	flag.Parse()
	if *pport == "" {
		if p := os.Getenv("SERVER_PORT"); p != "" {
			//using SERVER_PORT environment variable
			port = p
		}
		//using default port
	} else {
		//using port from flag
		port = *pport
	}
	p := 0
	var err error
	if p, err = strconv.Atoi(port); err != nil {
		log.Fatalf("invalid port '%s': %v", port, err)
	}
	projectID := dt.DetectProjectID
	namespace := "v1"
	mux, err := getMux(projectID, namespace)
	if err != nil {
		log.Fatalf("failed to get mux: %v", err)
	}
	//middlewares
	store, err := memstore.New(1000)
	if err != nil {
		log.Fatal(err)
	}
	th := throttled.RateLimit(throttled.PerSec(10), &throttled.VaryBy{Path: true}, store)
	chain := alice.New(th.Throttle, timeoutHandler, middlewares.Log, middlewares.TokenFromHeader) //nosurf.NewPure()

	server := http.Server{
		Addr:              fmt.Sprintf(":%d", p),
		ReadTimeout:       time.Second * 3,
		WriteTimeout:      time.Second * 3,
		ReadHeaderTimeout: time.Second * 3,
		IdleTimeout:       time.Second * 3,
		Handler:           chain.Then(mux),
	}
	log.Infof("Listening on port %d", p)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func timeoutHandler(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 1*time.Second, "timed out")
}

//getMux returns the
func getMux(projectID, namespace string) (http.Handler, error) {
	mux := bone.New()
	apiV1, err := getAPIV1(projectID, namespace)
	if err != nil {
		return nil, errors.Wrap(err, "get API V1")
	}
	mux.SubRoute("/v1", apiV1)
	return mux, nil
}

func getAPIV1(projectID, namespace string) (http.Handler, error) {
	mux := bone.New()
	accRepo, err := datastore.NewAccountsRepo(projectID, namespace)
	if err != nil {
		return nil, errors.Wrap(err, "accounts: set repo")
	}
	emailSrv := emails.New()
	accSrv, err := accounts.New(accRepo, emailSrv)
	if err != nil {
		return nil, errors.Wrap(err, "accounts: set service")
	}
	api, err := api.NewAPI(accSrv)
	if err != nil {
		return nil, errors.Wrap(err, "set api")
	}
	mux.GetFunc("/accounts/:id", api.Accounts().GetByID)
	return mux, nil
}
