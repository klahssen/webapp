package api

import (
	"fmt"

	pb "github.com/klahssen/webapp/pkg/domain"
	handlers "github.com/klahssen/webapp/pkg/serve/json/handlers"
	//"github.com/klahssen/webapp/pkg/log"
)

//api holds the inner services
type API struct {
	accounts *handlers.Accounts
}

//NewAPI gets a new instance of the API
func NewAPI(accounts pb.AccountsServer) (*API, error) {
	if accounts == nil {
		return nil, fmt.Errorf("accounts is nil")
	}
	return &API{accounts: handlers.NewAccounts(accounts)}, nil
}

//Accounts returns handlers related to accounts manipulation
func (a *API) Accounts() *handlers.Accounts {
	return a.accounts
}
