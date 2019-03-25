package main

import (
	"fmt"

	handlers "github.com/klahssen/webapp/app/json/handlers"
	pb "github.com/klahssen/webapp/pkg/domain"
	//"github.com/klahssen/webapp/pkg/log"
)

//api holds the inner services
type api struct {
	accounts *handlers.Accounts
}

//NewAPI gets a new instance of the API
func newAPI(accounts pb.AccountsServer) (*api, error) {
	if accounts == nil {
		return nil, fmt.Errorf("accounts is nil")
	}
	return &api{accounts: handlers.NewAccounts(accounts)}, nil
}

//Accounts returns handlers related to accounts manipulation
func (a *api) Accounts() *handlers.Accounts {
	return a.accounts
}
