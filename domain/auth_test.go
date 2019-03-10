package domain

import (
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
	delay := time.Duration(0)
	dur := time.Minute * 30
	tests := []struct {
		perms  *Permissions
		claims *jwt.StandardClaims
		keyID  string
		key    []byte
		res    *AccessToken
		err    error
	}{
		{
			perms: nil, claims: nil, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidPermissions,
		},
		{
			perms: &Permissions{}, claims: nil, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidClaims,
		},
		{
			perms: &Permissions{}, claims: &jwt.StandardClaims{}, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidClaims,
		},
		{
			perms: &Permissions{}, claims: &jwt.StandardClaims{Audience: "audience"}, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidClaims,
		},
		{
			perms: &Permissions{}, claims: &jwt.StandardClaims{Audience: "audience", Issuer: "issuer"}, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidClaims,
		},
		{
			perms: &Permissions{}, claims: &jwt.StandardClaims{Id: "", Audience: "audience", Issuer: "issuer", Subject: "subject", IssuedAt: t0.Unix(), ExpiresAt: t0.Add(dur).Unix(), NotBefore: t0.Add(delay).Unix()}, keyID: "", key: []byte(""),
			res: &AccessToken{Token: "eyJhbGciOiJIUzI1NiIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJDbGFpbXMiOnsiYXVkIjoiYXVkaWVuY2UiLCJleHAiOjE4MDAsImlzcyI6Imlzc3VlciIsInN1YiI6InN1YmplY3QifSwiUGVybXMiOnsidWlkIjoiIiwidHlwZSI6MCwic3RhdHVzIjowLCJwZXJtaXNzaW9ucyI6bnVsbH19.szVOm2HhuXCLq14lbzkXMIv5GZRS3YKq1zdWUp-dbSQ"},
			err: nil,
		},
	}
	for ind, test := range tests {
		res, err := test.perms.GetToken(test.claims, t0, delay, dur, test.keyID, test.key)
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
	}
	for ind, test := range tests {
		err := test.perms.validate()
		te.CheckError(ind, test.err, err)
	}
}
