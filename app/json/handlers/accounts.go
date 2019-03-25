package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-zoo/bone"
	"github.com/klahssen/webapp/app/json/format"
	pb "github.com/klahssen/webapp/pkg/domain"
	h "github.com/klahssen/webapp/pkg/http"
	"github.com/klahssen/webapp/pkg/log"
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
	t0 := time.Now()
	ctx := r.Context()
	token, ok := ctx.Value(h.JwtTokenInCtx).(string)
	t, ok2 := ctx.Value(h.ReqTimeInCtx).(time.Time)
	if ok2 {
		t0 = t
	}
	log.Infof("jwttoken,ok from ctx: '%s', %v", token, ok)
	if !ok || token == "" {
		format.WriteResponse(w, http.StatusUnauthorized, fmt.Errorf("missing jwt token"), nil, t0)
		return
	}
	log.Info("received new GET by ID")
	accountID := bone.GetValue(r, "id")
	acc, err := a.svc.Get(ctx, &pb.AccountID{Id: accountID})
	if err != nil {
		format.WriteResponse(w, http.StatusInternalServerError, err, nil, t0)
		return
	}
	format.WriteResponse(w, http.StatusOK, nil, acc, t0)
}
