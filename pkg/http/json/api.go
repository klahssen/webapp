package json

import (
	"fmt"

	pb "github.com/klahssen/webapp/pkg/domain"
)

//API holds the inner services
type API struct {
	accounts *accountsHandlers
}

//NewAPI gets a new instance of the API
func NewAPI(accounts pb.AccountsServer) (*API, error) {
	if accounts == nil {
		return nil, fmt.Errorf("accounts is nil")
	}
	return &API{accounts: &accountsHandlers{svc: accounts}}, nil
}

//Accounts returns handlers related to accounts manipulation
func (a *API) Accounts() *accountsHandlers {
	return a.accounts
}
