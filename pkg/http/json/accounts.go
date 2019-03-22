package json

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-zoo/bone"
	pb "github.com/klahssen/webapp/pkg/domain"
	h "github.com/klahssen/webapp/pkg/http"
)

//accountsHandlers holds all handlers related to accounts manipulation
type accountsHandlers struct {
	svc pb.AccountsServer
}

//GetByID handler
func (a *accountsHandlers) GetByID(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()
	ctx := r.Context()
	token, ok := ctx.Value(h.JwtTokenInCtx).(string)
	t, ok2 := ctx.Value(h.ReqTimeInCtx).(time.Time)
	if ok2 {
		t0 = t
	}
	if ok || token == "" {
		WriteResponse(w, http.StatusUnauthorized, fmt.Errorf("missing jwt token"), nil, t0)
		return
	}
	accountID := bone.GetValue(r, "id")
	acc, err := a.svc.Get(ctx, &pb.AccountID{Id: accountID})
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, err, nil, t0)
		return
	}
	WriteResponse(w, http.StatusOK, nil, acc, t0)
}
