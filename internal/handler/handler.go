package handler

import (
	"log/slog"
	"sync"
	"time"

	"github.com/nexow/nexow-gtn-auth/internal/gtn"
	"github.com/nexow/nexow-gtn-auth/internal/tokenstore"
)

// Handler exposes HTTP endpoints for auth and onboarding.
type Handler struct {
	gtnClient  *gtn.Client
	store      *tokenstore.Store
	corsOrigin string
	mu         sync.Mutex
}

// New creates a new Handler.
func New(gtnClient *gtn.Client, store *tokenstore.Store, corsOrigin string) *Handler {
	return &Handler{
		gtnClient:  gtnClient,
		store:      store,
		corsOrigin: corsOrigin,
	}
}

// getServerToken returns a valid server token, refreshing if needed.
func (h *Handler) getServerToken() (string, error) {
	access, needRefresh := h.store.Get()
	if !needRefresh && access != "" {
		return access, nil
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	access, needRefresh = h.store.Get()
	if !needRefresh && access != "" {
		return access, nil
	}

	if refreshToken := h.store.GetRefreshToken(); refreshToken != "" {
		resp, err := h.gtnClient.RefreshServerToken(refreshToken)
		if err == nil {
			h.store.Set(
				resp.AccessToken,
				resp.RefreshToken,
				time.Unix(resp.AccessTokenExpiresAt, 0),
				time.Unix(resp.RefreshTokenExpiresAt, 0),
			)
			return resp.AccessToken, nil
		}
		slog.Warn("server token refresh failed, will get new token", "error", err)
		h.store.Invalidate()
	}

	resp, err := h.gtnClient.GetServerToken()
	if err != nil {
		return "", err
	}
	h.store.Set(
		resp.AccessToken,
		resp.RefreshToken,
		time.Unix(resp.AccessTokenExpiresAt, 0),
		time.Unix(resp.RefreshTokenExpiresAt, 0),
	)
	return resp.AccessToken, nil
}
