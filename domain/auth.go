package domain

import (
	"github.com/dgrijalva/jwt-go"
)

type permsClaims struct {
	Claims *jwt.StandardClaims
	Perms  *Permissions
}

//Valid method to implement the jwt.Claims interface
func (p *permsClaims) Valid() error {
	return nil
}

//GetPermissions extracts permissions from an access token
func (t *AccessToken) GetPermissions(keyFunc jwt.Keyfunc) (*Permissions, error) {
	claims := &permsClaims{}
	_, err := jwt.ParseWithClaims(t.Token, claims, keyFunc)
	if err != nil {
		return nil, &ErrInvalidAccessToken
	}
	return claims.Perms, nil
}

//GetToken returns an access token string from permissions
func (p *Permissions) GetToken(claims *jwt.StandardClaims, keyID string, signingKey []byte) (*AccessToken, error) {
	if p == nil {
		return nil, &ErrInvalidPermissions
	}
	if claims == nil {
		return nil, &ErrInvalidClaims
	}
	jwtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &permsClaims{Claims: claims, Perms: p})
	jwtoken.Header["kid"] = keyID
	tokenstr, err := jwtoken.SignedString(signingKey)
	if err != nil {
		return nil, &ErrFailedToGenerateJwtToken
	}
	return &AccessToken{Token: tokenstr}, nil
}
