package domain

import (
	"fmt"
	"testing"
	"time"

	"github.com/klahssen/tester"

	"github.com/dgrijalva/jwt-go"
)

const (
	key1 = "1qsiaU8!dmJJ2)9J"
)

func stdKeyFunc(jwttoken *jwt.Token) (interface{}, error) {
	if jwttoken == nil {
		return nil, &ErrInvalidAccessToken
	}
	if _, ok := jwttoken.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, &ErrInvalidAccessToken //fmt.Errorf("Unexpected signing method: %v", jwttoken.Header["alg"])
	}
	typ := jwttoken.Header["typ"]
	alg := jwttoken.Header["alg"]
	keyID, ok := jwttoken.Header["kid"].(string)
	if typ != "JWT" {
		return nil, &ErrInvalidAccessToken
	}
	if alg != "HS256" {
		return nil, &ErrInvalidAccessToken
	}
	if !ok {
		return nil, &ErrInvalidAccessToken
	}
	switch keyID {
	case "key1":
		return key1, nil
	}
	return nil, &ErrInvalidAccessToken
}

func TestStdKeyFunc(t *testing.T) {
	te := tester.NewT(t)
	tests := []struct {
		t   *jwt.Token
		key interface{}
		err error
	}{
		{
			t:   &jwt.Token{Method: jwt.SigningMethodES384, Header: map[string]interface{}{"kid": "key1"}},
			key: nil,
			err: &ErrInvalidAccessToken,
		},
		{
			t:   &jwt.Token{Method: jwt.SigningMethodHS256, Header: map[string]interface{}{"kid": "key1", "typ": "JWT", "alg": "HS256"}},
			key: key1,
			err: nil,
		},
		{
			t:   &jwt.Token{Header: map[string]interface{}{"kid": "key2"}},
			key: nil,
			err: &ErrInvalidAccessToken,
		},
	}
	for ind, test := range tests {
		key, err := stdKeyFunc(test.t)
		te.CheckError(ind, test.err, err)
		if err != nil {
			continue
		}
		te.DeepEqual(ind, "key", test.key, key)
	}
}

func TestGetToken(t *testing.T) {
	te := tester.NewT(t)
	t0 := time.Unix(0, 0)
	t1 := time.Unix(1000000000, 0) //2001-09-09 01:46:40 +0000 UTC
	delay := time.Duration(0)
	dur := time.Hour * 24 * 365 * 100
	tests := []struct {
		t      time.Time
		perms  *Permissions
		claims *jwt.StandardClaims
		keyID  string
		key    []byte
		res    *AccessToken
		err    error
	}{
		{
			t: t0, perms: nil, claims: nil, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidPermissions,
		},
		{
			t: t0, perms: &Permissions{}, claims: nil, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidClaims,
		},
		{
			t: t0, perms: &Permissions{}, claims: &jwt.StandardClaims{}, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidClaims,
		},
		{
			t: t0, perms: &Permissions{}, claims: &jwt.StandardClaims{Audience: "audience"}, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidClaims,
		},
		{
			t: t0, perms: &Permissions{}, claims: &jwt.StandardClaims{Audience: "audience", Issuer: "issuer"}, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidClaims,
		},
		{
			t: t0, perms: &Permissions{}, claims: &jwt.StandardClaims{Id: "", Audience: "audience", Issuer: "issuer", Subject: "subject", IssuedAt: t0.Unix(), ExpiresAt: t0.Add(dur).Unix(), NotBefore: t0.Add(delay).Unix()}, keyID: "", key: []byte(""),
			res: &AccessToken{Token: "eyJhbGciOiJIUzI1NiIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJDbGFpbXMiOnsiYXVkIjoiYXVkaWVuY2UiLCJleHAiOjMxNTM2MDAwMDAsImlzcyI6Imlzc3VlciIsInN1YiI6InN1YmplY3QifSwiUGVybXMiOnsidWlkIjoiIiwidHlwZSI6MCwic3RhdHVzIjowLCJwZXJtaXNzaW9ucyI6bnVsbH19.cmbtzf6pklQVh8xqWaeHlHe21_YLMcHq1amnmWLRjps"},
			err: nil,
		},
		{
			t: t1, perms: &Permissions{Uid: "001", AccountType: AccountType_USER, AccountStatus: AccountStatus_ACTIVE, Permissions: []string{"abcd"}}, claims: &jwt.StandardClaims{Id: "", Audience: "audience", Issuer: "issuer", Subject: "subject", IssuedAt: t1.Unix(), ExpiresAt: t1.Add(dur).Unix(), NotBefore: t1.Add(delay).Unix()}, keyID: "", key: []byte(""),
			res: &AccessToken{Token: "eyJhbGciOiJIUzI1NiIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJDbGFpbXMiOnsiYXVkIjoiYXVkaWVuY2UiLCJleHAiOjQxNTM2MDAwMDAsImlhdCI6MTAwMDAwMDAwMCwiaXNzIjoiaXNzdWVyIiwibmJmIjoxMDAwMDAwMDAwLCJzdWIiOiJzdWJqZWN0In0sIlBlcm1zIjp7InVpZCI6IjAwMSIsInR5cGUiOjAsInN0YXR1cyI6MSwicGVybWlzc2lvbnMiOlsiYWJjZCJdfX0.Z9pAsH2pQF7Uqpr8NRk_-5_giMZLbogNvhGcLre6XmQ"},
			err: nil,
		},
	}
	for ind, test := range tests {
		res, err := test.perms.GetToken(test.claims, test.t, delay, dur, test.keyID, test.key)
		te.CheckError(ind, test.err, err)
		if test.err != nil {
			continue
		}
		te.DeepEqual(ind, "token", test.res, res)

		if (test.res == nil && res != nil) || (test.res != nil && res == nil) {
			t.Errorf("test [%d]: expected res %v received %v", ind, test.res, res)
			continue
		}
	}
}

func TestPermsValidate(t *testing.T) {
	te := tester.NewT(t)
	tests := []struct {
		perms *permsClaims
		err   error
	}{
		{
			perms: &permsClaims{},
			err:   &ErrInvalidClaims,
		},
		{
			perms: &permsClaims{Claims: &jwt.StandardClaims{}},
			err:   &ErrInvalidPermissions,
		},
		{
			perms: &permsClaims{Claims: &jwt.StandardClaims{}, Perms: &Permissions{}},
			err:   &ErrInvalidClaims,
		},
		{
			perms: &permsClaims{Claims: &jwt.StandardClaims{Audience: "audience"}, Perms: &Permissions{}},
			err:   &ErrInvalidClaims,
		},
		{
			perms: &permsClaims{Claims: &jwt.StandardClaims{Audience: "audience", Issuer: "issuer"}, Perms: &Permissions{}},
			err:   &ErrInvalidClaims,
		},
		{
			perms: &permsClaims{Claims: &jwt.StandardClaims{Audience: "audience", Issuer: "issuer", Subject: "subject"}, Perms: &Permissions{}},
			err:   nil,
		},
	}

	for ind, test := range tests {
		err := test.perms.validate()
		te.CheckError(ind, test.err, err)
	}
}

func TestGetPermissions(t *testing.T) {
	te := tester.NewT(t)
	f := func(issuer, audience, id, subject string) error {
		if issuer != "issuer" {
			return fmt.Errorf("invalid issuer '%s'", issuer)
		}
		if audience != "audience" {
			return fmt.Errorf("invalid audience '%s'", audience)
		}
		if subject != "subject" {
			return fmt.Errorf("invalid subject '%s'", subject)
		}
		return nil
	}
	f1 := func(issuer, audience, id, subject string) error {
		if issuer != "issuer1" {
			return fmt.Errorf("invalid issuer '%s'", issuer)
		}
		if audience != "audience1" {
			return fmt.Errorf("invalid audience '%s'", audience)
		}
		if id != "id1" {
			return fmt.Errorf("invalid id '%s'", id)
		}
		if subject != "subject1" {
			return fmt.Errorf("invalid subject '%s'", subject)
		}
		return nil
	}
	tests := []struct {
		a       *AccessToken
		keyFunc jwt.Keyfunc
		f       ClaimsValidator
		res     *Permissions
		err     error
	}{
		{
			a: nil, err: &ErrInvalidAccessToken,
		},
		{
			a: &AccessToken{Token: ""}, err: fmt.Errorf("token contains an invalid number of segments"),
		},
		{
			a: &AccessToken{
				Token: "eyJhbGciOiJIUzI1NiIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJDbGFpbXMiOnsiYXVkIjoiYXVkaWVuY2UiLCJleHAiOjQxNTM2MDAwMDAsImlhdCI6MTAwMDAwMDAwMCwiaXNzIjoiaXNzdWVyIiwibmJmIjoxMDAwMDAwMDAwLCJzdWIiOiJzdWJqZWN0In0sIlBlcm1zIjp7InVpZCI6IjAwMSIsInR5cGUiOjAsInN0YXR1cyI6MSwicGVybWlzc2lvbnMiOlsiYWJjZCJdfX0.Z9pAsH2pQF7Uqpr8NRk_-5_giMZLbogNvhGcLre6XmQ",
			},
			keyFunc: nil,
			res:     nil,
			err:     fmt.Errorf("no Keyfunc was provided."),
		},
		{
			a: &AccessToken{
				Token: "eyJhbGciOiJIUzI1NiIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJDbGFpbXMiOnsiYXVkIjoiYXVkaWVuY2UiLCJleHAiOjQxNTM2MDAwMDAsImlhdCI6MTAwMDAwMDAwMCwiaXNzIjoiaXNzdWVyIiwibmJmIjoxMDAwMDAwMDAwLCJzdWIiOiJzdWJqZWN0In0sIlBlcm1zIjp7InVpZCI6IjAwMSIsInR5cGUiOjAsInN0YXR1cyI6MSwicGVybWlzc2lvbnMiOlsiYWJjZCJdfX0.Z9pAsH2pQF7Uqpr8NRk_-5_giMZLbogNvhGcLre6XmQ",
			},
			f:       f,
			keyFunc: func(a *jwt.Token) (interface{}, error) { return []byte(""), nil },
			res:     &Permissions{Uid: "001", AccountType: AccountType_USER, AccountStatus: AccountStatus_ACTIVE, Permissions: []string{"abcd"}},
			err:     nil,
		},
		{
			a: &AccessToken{
				Token: "eyJhbGciOiJIUzI1NiIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJDbGFpbXMiOnsiYXVkIjoiYXVkaWVuY2UiLCJleHAiOjQxNTM2MDAwMDAsImlhdCI6MTAwMDAwMDAwMCwiaXNzIjoiaXNzdWVyIiwibmJmIjoxMDAwMDAwMDAwLCJzdWIiOiJzdWJqZWN0In0sIlBlcm1zIjp7InVpZCI6IjAwMSIsInR5cGUiOjAsInN0YXR1cyI6MSwicGVybWlzc2lvbnMiOlsiYWJjZCJdfX0.Z9pAsH2pQF7Uqpr8NRk_-5_giMZLbogNvhGcLre6XmQ",
			},
			f:       f1,
			keyFunc: func(a *jwt.Token) (interface{}, error) { return []byte(""), nil },
			res:     &Permissions{Uid: "001", AccountType: AccountType_USER, AccountStatus: AccountStatus_ACTIVE, Permissions: []string{"abcd"}},
			err:     fmt.Errorf("invalid issuer 'issuer'"),
		},
	}

	for ind, test := range tests {
		res, err := test.a.GetPermissions(test.keyFunc, test.f)
		te.CheckError(ind, test.err, err)
		if err != nil {
			continue
		}
		if res != nil && test.res == nil {
			t.Errorf("test [%d]: unexpected nil res", ind)
			continue
		} else if res == nil && test.res != nil {
			t.Errorf("test [%d]: expected nil res", ind)
			continue
		}
		if test.res == nil {
			continue
		}
		te.DeepEqual(ind, "res.Uid", test.res.Uid, res.Uid)
		te.DeepEqual(ind, "res.AccountStatus", test.res.AccountStatus, res.AccountStatus)
		te.DeepEqual(ind, "res.AccountType", test.res.AccountType, res.AccountType)
		te.DeepEqual(ind, "res.Permissions", test.res.Permissions, res.Permissions)
	}
}
