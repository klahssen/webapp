package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type permsClaims struct {
	Claims *jwt.StandardClaims
	Perms  *Permissions
}

//ClaimsValidator is a function that validates claims
type ClaimsValidator func(issuer, audience, id, subject string) error

//AuthorizeFunc checks if an action is authorized for a set of permissions
type AuthorizeFunc func(target *AuthParams) (*AuthResp, error)

//Valid method to implement the jwt.Claims interface
func (p *permsClaims) Valid() error {
	if err := p.validate(); err != nil {
		return err
	}
	return p.Claims.Valid()
}

//validate method checks if not nil pointers and if no empty fields
func (p *permsClaims) validate() error {
	if p == nil {
		return &ErrInvalidClaims
	}
	if p.Claims == nil {
		return &ErrInvalidClaims
	}
	if p.Perms == nil {
		return &ErrInvalidPermissions
	}
	if p.Claims.Audience == "" || p.Claims.Issuer == "" || p.Claims.Subject == "" {
		return &ErrInvalidClaims
	}

	return nil
}

//GetPermissions extracts permissions from an access token
func (t *AccessToken) GetPermissions(keyFunc jwt.Keyfunc, f ClaimsValidator) (*Permissions, error) {
	if t == nil {
		return nil, &ErrInvalidAccessToken
	}
	claims := &permsClaims{}
	_, err := jwt.ParseWithClaims(t.Token, claims, keyFunc)
	if err != nil {
		//return nil, &ErrInvalidAccessToken
		return nil, err
	}
	if err = f(claims.Claims.Issuer, claims.Claims.Audience, claims.Claims.Id, claims.Claims.Subject); err != nil {
		return nil, err
	}
	return claims.Perms, nil
}

//GetToken returns an access token string from permissions. delay is used in not before, t is used for issued at and validity for expiration
func (p *Permissions) GetToken(claims *jwt.StandardClaims, t time.Time, delay, validity time.Duration, keyID string, signingKey []byte) (*AccessToken, error) {
	if p == nil {
		return nil, &ErrInvalidPermissions
	}
	if claims == nil {
		return nil, &ErrInvalidClaims
	}
	if validity < 0 {
		validity *= -1
	}
	if delay < 0 {
		delay *= -1
	}
	t.UTC()
	claims.IssuedAt = t.Unix()
	claims.ExpiresAt = t.Add(validity).Unix()
	claims.NotBefore = t.Add(delay).Unix()
	pc := &permsClaims{Claims: claims, Perms: p}
	if err := pc.validate(); err != nil {
		return nil, err
	}
	jwtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, pc)
	jwtoken.Header["kid"] = keyID
	tokenstr, err := jwtoken.SignedString(signingKey)
	if err != nil {
		return nil, &ErrFailedToGenerateJwtToken
	}
	return &AccessToken{Token: tokenstr}, nil
}
