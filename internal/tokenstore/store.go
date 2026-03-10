package tokenstore

import (
	"sync"
	"time"
)

// ServerTokens holds server access and refresh token with expiry.
type ServerTokens struct {
	AccessToken          string
	RefreshToken         string
	AccessTokenExpiresAt time.Time
	RefreshTokenExpiresAt time.Time
}

// Store caches server token and refreshes before expiry.
type Store struct {
	mu       sync.RWMutex
	access   string
	refresh  string
	expiry   time.Time
	refreshExpiry time.Time
	refreshMargin time.Duration
}

// NewStore creates a store. refreshMargin is how long before expiry to refresh (e.g. 10*time.Minute).
func NewStore(refreshMargin time.Duration) *Store {
	if refreshMargin <= 0 {
		refreshMargin = 10 * time.Minute
	}
	return &Store{refreshMargin: refreshMargin}
}

// Get returns the current access token if still valid (or valid within refresh margin).
// Caller should call Refresh() when this returns empty.
func (s *Store) Get() (access string, needRefresh bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.access == "" {
		return "", true
	}
	refreshBy := s.expiry.Add(-s.refreshMargin)
	if time.Now().After(refreshBy) {
		return s.access, true
	}
	return s.access, false
}

// Set stores new server tokens.
func (s *Store) Set(access, refresh string, accessExpiry, refreshExpiry time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.access = access
	s.refresh = refresh
	s.expiry = accessExpiry
	s.refreshExpiry = refreshExpiry
}

// GetRefreshToken returns the refresh token for server token refresh.
func (s *Store) GetRefreshToken() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.refresh
}

// Invalidate clears the cache (e.g. after failed refresh).
func (s *Store) Invalidate() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.access = ""
	s.refresh = ""
	s.expiry = time.Time{}
	s.refreshExpiry = time.Time{}
}
