package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	contx "github.com/klahssen/webapp/pkg/context"
	"github.com/klahssen/webapp/pkg/internal/errors"
	"github.com/klahssen/webapp/pkg/json/format"
)

var (
	errEmptyBody   = fmt.Errorf("empty JSON body")
	errInvalidBody = fmt.Errorf("invalid JSON body")
)

type reqInfo struct {
	t0    time.Time
	token string
	ctx   context.Context
}

func getReqInfo(w http.ResponseWriter, r *http.Request) (*reqInfo, bool) {
	t0 := time.Now()
	ctx := r.Context()
	token, ok := ctx.Value(contx.JwtToken).(string)
	t, ok2 := ctx.Value(contx.ReqTime).(time.Time)
	if ok2 {
		t0 = t
	}
	if !ok || token == "" {
		format.WriteResponse(w, http.StatusUnauthorized, fmt.Errorf("missing jwt token"), nil, t0)
		return nil, false
	}
	return &reqInfo{t0: t0, token: token, ctx: ctx}, true
}

func unmarshalJSON(r *http.Request, dest interface{}) error {
	if r.Body == nil {
		return errEmptyBody
	}
	err := json.NewDecoder(r.Body).Decode(dest)
	if err != nil {
		return errInvalidBody
	}
	return nil
}

func analyzeError(err error) (int, error) {
	if err == nil {
		return http.StatusOK, nil
	}
	if e, ok := err.(*errors.Error); ok {
		return e.GetStatusCode(), fmt.Errorf(e.Msg)
	}
	return http.StatusInternalServerError, err
}
