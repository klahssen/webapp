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
	t0 := time.Now().Unix()
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
			perms: &Permissions{}, claims: &jwt.StandardClaims{Audience: "audience", Issuer: "issuer", IssuedAt: t0}, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidClaims,
		},
		{
			perms: &Permissions{}, claims: &jwt.StandardClaims{Audience: "audience", Issuer: "issuer", IssuedAt: t0, Id: "id"}, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidClaims,
		},
		{
			perms: &Permissions{}, claims: &jwt.StandardClaims{Audience: "audience", Issuer: "issuer", IssuedAt: t0, Id: "id", Subject: ""}, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidClaims,
		},
		{
			perms: &Permissions{}, claims: &jwt.StandardClaims{Audience: "audience", Issuer: "issuer", IssuedAt: t0, Id: "id", ExpiresAt: t0 + 1000, Subject: ""}, keyID: "", key: []byte(""), res: nil, err: &ErrInvalidClaims,
		},
		{
			perms: &Permissions{}, claims: &jwt.StandardClaims{Audience: "audience", Issuer: "issuer", IssuedAt: t0, Id: "id", ExpiresAt: t0 + 1000, NotBefore: t0, Subject: ""}, keyID: "", key: []byte(key1), res: nil, err: &ErrInvalidClaims,
		},
	}
	for ind, test := range tests {
		res, err := test.perms.GetToken(test.claims, test.keyID, test.key)
		te.CheckError(ind, test.err, err)
		if err != nil {
			continue
		}
		te.DeepEqual(ind, "res", test.res, res)
	}
}
