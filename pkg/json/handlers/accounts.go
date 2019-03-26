package handlers

import (
	"net/http"

	"github.com/go-zoo/bone"
	pb "github.com/klahssen/webapp/pkg/domain"
	"github.com/klahssen/webapp/pkg/json/format"
)

//Accounts holds all handlers related to accounts manipulation
type Accounts struct {
	svc pb.AccountsServer
}

//NewAccounts returns an instance of Accounts handlers
func NewAccounts(svc pb.AccountsServer) *Accounts {
	return &Accounts{svc: svc}
}

//GetByID handler
func (a *Accounts) GetByID(w http.ResponseWriter, r *http.Request) {
	info, ok := getReqInfo(w, r)
	if !ok {
		return
	}
	accountUID := bone.GetValue(r, "id")
	acc, err := a.svc.Get(info.ctx, &pb.AccountID{Id: accountUID})
	status, err := analyzeError(err)
	if err != nil {
		format.WriteResponse(w, status, err, nil, info.t0)
		return
	}
	format.WriteResponse(w, http.StatusOK, nil, acc, info.t0)
}

//Create handler
func (a *Accounts) Create(w http.ResponseWriter, r *http.Request) {
	info, ok := getReqInfo(w, r)
	if !ok {
		return
	}
	params := &pb.AccountParams{}
	err := unmarshalJSON(r, params)
	if err != nil {
		format.WriteResponse(w, http.StatusBadRequest, err, nil, info.t0)
		return
	}
	acc, err := a.svc.Create(info.ctx, params)
	status, err := analyzeError(err)
	if err != nil {
		format.WriteResponse(w, status, err, nil, info.t0)
		return
	}
	format.WriteResponse(w, http.StatusOK, nil, acc, info.t0)
}
