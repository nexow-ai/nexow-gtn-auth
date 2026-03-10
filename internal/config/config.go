package config

import (
	"encoding/base64"
	"encoding/hex"
	"os"
	"strings"
)

type Config struct {
	GTNBaseURL        string
	AppKey            string
	AppSecret         string
	InstitutionCode   string
	UserID            string
	PrivateKeyDER     []byte
	ThrottleKey       string
	CORSAllowedOrigin string
}

func Load() (*Config, error) {
	gtnBase := os.Getenv("GTN_BASE_URL")
	if gtnBase == "" {
		gtnBase = "https://sandbox.globaltradingnetwork.com"
	}
	appKey := os.Getenv("GTN_APP_KEY")
	appSecret := os.Getenv("GTN_APP_SECRET")
	instCode := os.Getenv("GTN_INSTITUTION_CODE")
	userID := os.Getenv("GTN_USER_ID")
	if userID == "" {
		userID = "nexow-gtn-auth-1"
	}
	privateKeyEnc := os.Getenv("GTN_PRIVATE_KEY")
	throttleKey := os.Getenv("GTN_THROTTLE_KEY")
	corsOrigin := os.Getenv("CORS_ALLOWED_ORIGIN")
	if corsOrigin == "" {
		corsOrigin = "http://localhost:5173"
	}

	var der []byte
	if privateKeyEnc == "" {
		return nil, &MissingEnvError{Key: "GTN_PRIVATE_KEY"}
	}
	privateKeyEnc = strings.TrimSpace(privateKeyEnc)
	if isHex(privateKeyEnc) {
		der, _ = hex.DecodeString(privateKeyEnc)
	} else {
		der, _ = base64.StdEncoding.DecodeString(privateKeyEnc)
	}
	if len(der) == 0 {
		return nil, &MissingEnvError{Key: "GTN_PRIVATE_KEY (invalid encoding)"}
	}
	if appKey == "" {
		return nil, &MissingEnvError{Key: "GTN_APP_KEY"}
	}
	if appSecret == "" {
		return nil, &MissingEnvError{Key: "GTN_APP_SECRET"}
	}
	if instCode == "" {
		return nil, &MissingEnvError{Key: "GTN_INSTITUTION_CODE"}
	}
	if throttleKey == "" {
		return nil, &MissingEnvError{Key: "GTN_THROTTLE_KEY"}
	}

	return &Config{
		GTNBaseURL:        strings.TrimSuffix(gtnBase, "/"),
		AppKey:            appKey,
		AppSecret:         appSecret,
		InstitutionCode:   instCode,
		UserID:            userID,
		PrivateKeyDER:     der,
		ThrottleKey:       throttleKey,
		CORSAllowedOrigin: corsOrigin,
	}, nil
}

func isHex(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	for _, c := range s {
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F') {
			continue
		}
		return false
	}
	return len(s) > 0
}

type MissingEnvError struct{ Key string }

func (e *MissingEnvError) Error() string {
	return "missing or invalid env: " + e.Key
}
