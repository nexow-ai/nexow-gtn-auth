package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nexow/nexow-gtn-auth/internal/config"
	"github.com/nexow/nexow-gtn-auth/internal/gtn"
	"github.com/nexow/nexow-gtn-auth/internal/handler"
	"github.com/nexow/nexow-gtn-auth/internal/tokenstore"
)

func main() {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		slog.Error("config load failed", "error", err)
		os.Exit(1)
	}

	gtnClient := gtn.NewClient(
		cfg.GTNBaseURL,
		cfg.AppKey,
		cfg.AppSecret,
		cfg.InstitutionCode,
		cfg.UserID,
		cfg.PrivateKeyDER,
		cfg.ThrottleKey,
	)
	store := tokenstore.NewStore(10 * time.Minute)
	h := handler.New(gtnClient, store, cfg.CORSAllowedOrigin)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /auth/customer/token", h.CORS(h.CustomerToken))
	mux.HandleFunc("POST /auth/customer/token/refresh", h.CORS(h.CustomerTokenRefresh))
	mux.HandleFunc("POST /onboard/customer", h.CORS(h.CreateCustomer))

	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	slog.Info("listening", "addr", addr)
	corsHandler := handler.CORSWrap(cfg.CORSAllowedOrigin, mux)
	if err := http.ListenAndServe(addr, corsHandler); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
