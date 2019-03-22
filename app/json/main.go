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
	"github.com/klahssen/webapp/pkg/http/json"
	"github.com/klahssen/webapp/pkg/repos/gcloud/datastore"
	"github.com/klahssen/webapp/pkg/services/accounts"
	emails "github.com/klahssen/webapp/pkg/services/emails/local"
	"github.com/pkg/errors"
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
		logger.Fatalf("invalid port '%s': %v", port, err)
	}
	projectID := dt.DetectProjectID
	namespace := ""
	mux, err := getMux(projectID, namespace)
	if err != nil {
		logger.Fatalf("failed to get mux: %v", err)
	}
	server := http.Server{
		Addr:              fmt.Sprintf(":%d", p),
		ReadTimeout:       time.Second * 3,
		WriteTimeout:      time.Second * 3,
		ReadHeaderTimeout: time.Second * 3,
		IdleTimeout:       time.Second * 3,
		Handler:           mux,
	}
	logger.Infof("Listening on port %d", p)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

//getMux returns the
func getMux(projectID, namespace string) (http.Handler, error) {
	mux := bone.New()
	apiV1, err := getAPIV1(projectID, namespace)
	if err != nil {
		return nil, errors.Wrap(err, "get API V1")
	}
	mux.SubRoute("/v1/", apiV1)
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
	api, err := json.NewAPI(accSrv)
	if err != nil {
		return nil, errors.Wrap(err, "set api")
	}
	mux.GetFunc("accounts/:id", api.Accounts().GetByID)
	return mux, nil
}
