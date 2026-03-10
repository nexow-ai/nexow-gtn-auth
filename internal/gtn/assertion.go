package gtn

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// AssertionClaims is the JWT payload for GTN server token request.
type AssertionClaims struct {
	jwt.RegisteredClaims
	InstCode string `json:"instCode"`
	UserID   string `json:"userId"`
}

// BuildAssertion creates a signed JWT (RS256) for GTN auth/token.
// privateKeyDER is PKCS#8 DER-encoded RSA private key.
func BuildAssertion(privateKeyDER []byte, appKey, institutionCode, userID string, expirySeconds int64) (string, error) {
	if expirySeconds <= 0 {
		expirySeconds = 3600
	}
	now := time.Now()
	exp := now.Add(time.Duration(expirySeconds) * time.Second)

	claims := AssertionClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    appKey,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		InstCode: institutionCode,
		UserID:   userID,
	}

	key, err := parsePrivateKeyDER(privateKeyDER)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(key)
}

func parsePrivateKeyDER(der []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(der)
	var keyBytes []byte
	if block != nil {
		keyBytes = block.Bytes
	} else {
		keyBytes = der
	}
	key, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		rsaKey, err2 := x509.ParsePKCS1PrivateKey(keyBytes)
		if err2 != nil {
			return nil, errors.New("invalid private key: " + err.Error())
		}
		return rsaKey, nil
	}
	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key is not RSA")
	}
	return rsaKey, nil
}
