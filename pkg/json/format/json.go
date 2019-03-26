package format

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
)

//WriteResponse write a formated json response
func WriteResponse(w http.ResponseWriter, code int, err error, data interface{}, t0 time.Time) {
	w.WriteHeader(code)
	resp := &Response{Data: data, Dur: fmt.Sprint(time.Since(t0)), OK: false}
	if code < 300 {
		resp.OK = true
	}
	if err != nil {
		resp.Err = err.Error()
	}
	b, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte(spew.Sdump(resp)))
		return
	}
	w.Write(b)
}

//Response for json request
type Response struct {
	Data interface{} `json:"data"`
	Err  interface{} `json:"err"`
	Dur  string      `json:"dur"`
	OK   bool        `json:"ok"`
}
