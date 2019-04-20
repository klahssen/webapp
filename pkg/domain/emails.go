package domain

import (
	"fmt"
	"strings"

	"github.com/klahssen/webapp/internal/validators"
)

//EmailConfirmationParams data
type EmailConfirmationParams struct {
	Addr string `json:"addr"`
	Lang string `json:"lang"`
}

//Format inner data
func (ep *EmailConfirmationParams) Format() {
	if ep == nil {
		return
	}
	ep.Addr = strings.ToLower(strings.Replace(ep.Addr, " ", "", -1))
	ep.Lang = strings.ToLower(strings.Replace(ep.Lang, " ", "", -1))
}

//Validate inner data
func (ep *EmailConfirmationParams) Validate() error {
	if ep == nil {
		return fmt.Errorf("nil pointer")
	}
	var err error
	if err = ValidateEmail(ep.Addr); err != nil {
		return err
	}
	if !validators.IsValidLang(ep.Lang) {
		ep.Lang = validators.LangENG
	}
	return nil
}
