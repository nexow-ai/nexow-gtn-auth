package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/nexow/nexow-gtn-auth/internal/gtn"
)

// CustomerTokenRequest is the body for POST /auth/customer/token.
type CustomerTokenRequest struct {
	CustomerNumber string `json:"customerNumber"`
}

// CustomerTokenResponse is the response for customer token endpoints.
type CustomerTokenResponse struct {
	AccessToken           string `json:"accessToken"`
	RefreshToken          string `json:"refreshToken"`
	AccessTokenExpiresAt  int64  `json:"accessTokenExpiresAt"`
	RefreshTokenExpiresAt int64  `json:"refreshTokenExpiresAt"`
	TokenType             string `json:"tokenType,omitempty"`
}

// CustomerTokenRefreshRequest is the body for POST /auth/customer/token/refresh.
type CustomerTokenRefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

// CustomerToken returns a customer access token (POST /auth/customer/token).
func (h *Handler) CustomerToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req CustomerTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid body")
		return
	}
	if req.CustomerNumber == "" {
		writeJSONError(w, http.StatusBadRequest, "customerNumber required")
		return
	}

	serverToken, err := h.getServerToken()
	if err != nil {
		slog.Error("get server token failed", "error", err)
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.gtnClient.GetCustomerToken(serverToken, req.CustomerNumber)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, CustomerTokenResponse{
		AccessToken:           resp.AccessToken,
		RefreshToken:          resp.RefreshToken,
		AccessTokenExpiresAt:  resp.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: resp.RefreshTokenExpiresAt,
		TokenType:             resp.TokenType,
	})
}

// CustomerTokenRefresh refreshes a customer token (POST /auth/customer/token/refresh).
func (h *Handler) CustomerTokenRefresh(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req CustomerTokenRefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid body")
		return
	}
	if req.RefreshToken == "" {
		writeJSONError(w, http.StatusBadRequest, "refreshToken required")
		return
	}

	resp, err := h.gtnClient.RefreshCustomerToken(req.RefreshToken)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, CustomerTokenResponse{
		AccessToken:           resp.AccessToken,
		RefreshToken:          resp.RefreshToken,
		AccessTokenExpiresAt:  resp.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: resp.RefreshTokenExpiresAt,
		TokenType:             resp.TokenType,
	})
}

// CreateCustomer proxies GTN create customer (POST /onboard/customer).
func (h *Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid body")
		return
	}
	var req gtn.CreateCustomerRequest
	if err := json.Unmarshal(body, &req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	if req.ReferenceNumber == "" || req.InstitutionCode == "" {
		writeJSONError(w, http.StatusBadRequest, "referenceNumber and institutionCode required")
		return
	}

	serverToken, err := h.getServerToken()
	if err != nil {
		slog.Error("get server token failed", "error", err)
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.gtnClient.CreateCustomer(serverToken, &req)
	if err != nil {
		if resp != nil {
			writeJSON(w, http.StatusOK, resp)
		} else {
			writeJSONError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	writeJSON(w, http.StatusOK, resp)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
