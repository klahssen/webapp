package json

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
)

//WriteResponse write a formated json response
func WriteResponse(w http.ResponseWriter, code int, err error, data interface{}, t0 time.Time) {
	w.WriteHeader(code)
	resp := &Response{Data: data, Err: err, Dur: time.Since(t0), OK: false}
	if code < 300 {
		resp.OK = true
	}
	b, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte(spew.Sdump(resp)))
		return
	}
	w.Write(b)
}

//Response for json request
type Response struct {
	Data interface{}   `json:"data"`
	Err  error         `json:"err"`
	Dur  time.Duration `json:"dur"`
	OK   bool          `json:"ok"`
}
