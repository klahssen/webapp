package main

import (
	"fmt"

	pb "github.com/klahssen/webapp/pkg/domain"
	handlers "github.com/klahssen/webapp/pkg/json/handlers"
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
