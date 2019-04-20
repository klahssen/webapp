package format

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/klahssen/webapp/internal/log"
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
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Infof("failed to json encode response: %v", err)
		if _, err = w.Write([]byte(spew.Sdump(resp))); err != nil {
			log.Infof("failed to write dump of response: %v", err)
		}
	}
}

//Response for json request
type Response struct {
	Data interface{} `json:"data"`
	Err  interface{} `json:"err"`
	Dur  string      `json:"dur"`
	OK   bool        `json:"ok"`
}
