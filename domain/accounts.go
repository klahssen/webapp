package domain

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/klahssen/webapp/internal/errors"
	"github.com/klahssen/webapp/internal/validators"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost        = 12
	minPasswordLength = 8
	maxPasswordLength = 62
)

//Format inner data
func (ap *AccountParams) Format() {
	if ap == nil {
		return
	}
	ap.Uid = strings.Replace(ap.Uid, " ", "", -1)
	ap.Em = strings.Replace(ap.Em, " ", "", -1)
	ap.Pw = strings.Replace(ap.Pw, " ", "", -1)
}

//IsEmpty checks if there is data
func (ap *AccountParams) IsEmpty() bool {
	if ap == nil {
		return true
	}
	return ap.Uid == "" || (ap.Em == "" && ap.Pw == "")
}

//Format inner data
func (a *AccountEntity) Format() {
	if a == nil {
		return
	}
	a.Em = strings.ToLower(strings.Replace(a.Em, " ", "", -1))
}

//Hash will hash the password
func (a *AccountEntity) Hash() error {
	if a == nil {
		return fmt.Errorf("nil pointer")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(a.Pw), bcryptCost)
	if err != nil {
		return &errors.Error{Code: http.StatusInternalServerError, Msg: err.Error()}
	}
	a.Pw = string(hashed)
	return nil
}

//ComparePassword checks if the provided password matches accounts password
func (a *AccountEntity) ComparePassword(pw string) error {
	if a == nil {
		return fmt.Errorf("nil pointer")
	}
	err := bcrypt.CompareHashAndPassword([]byte(a.Pw), []byte(pw))
	if err != nil {
		return &ErrAuthenticationFailed
	}
	return nil
}

//ValidateNew AccountEntity data
func (a *AccountEntity) ValidateNew() error {
	if a == nil {
		return fmt.Errorf("empty pointer")
	}
	var err error
	if err = ValidateEmail(a.Em); err != nil {
		return err
	}
	if err = ValidatePassword(a.Pw); err != nil {
		return err
	}
	a.CreatedAt = time.Now().UTC().Unix()
	a.RecordUpdate()
	return nil
}

//Validate account info
func (a *AccountEntity) Validate() error {
	if a == nil {
		return fmt.Errorf("empty pointer")
	}
	var err error
	if err = validators.EmailAddress(a.Em); err != nil {
		return err
	}
	if err = ValidatePassword(a.Pw); err != nil {
		return err
	}
	a.RecordUpdate()
	return nil
}

//RecordUpdate stores timestamp of the update
func (a *AccountEntity) RecordUpdate() {
	a.UpdatedAt = time.Now().UTC().Unix()
}

//ValidatePassword validator
func ValidatePassword(pw string) error {
	l := len(pw)
	if l < minPasswordLength || l > maxPasswordLength {
		return &ErrAccountInvalidPassword
	}
	return nil
}

//ValidateEmail validator
func ValidateEmail(email string) error {
	if err := validators.EmailAddress(email); err != nil {
		return &ErrAccountInvalidPassword
	}
	return nil
}
