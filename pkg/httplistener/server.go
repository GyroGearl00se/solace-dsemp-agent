package httplistener

import (
	"context"
	"crypto/subtle"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/GyroGearl00se/solace-dsemp-agent/pkg/state"
	"github.com/sirupsen/logrus"
)

const (
	maxBodyBytes    = 10 << 20 // 10 MB
	readTimeout     = 30 * time.Second
	writeTimeout    = 30 * time.Second
	shutdownTimeout = 10 * time.Second
)

// Config holds the HTTP listener configuration.
type Config struct {
	Enabled     bool
	Address     string // e.g. ":9001"
	TLSEnabled  bool
	TLSCertFile string
	TLSKeyFile  string
	Username    string
	Password    string
	AESKey      string
}

// Server is an HTTP server that accepts TargetState payloads via POST.
type Server struct {
	cfg      Config
	callback func(*state.TargetState)
}

// NewServer creates a new HTTP listener Server.
func NewServer(cfg Config, callback func(*state.TargetState)) *Server {
	return &Server{cfg: cfg, callback: callback}
}

// Start runs the HTTP server until ctx is cancelled.
func (s *Server) Start(ctx context.Context) {
	mux := http.NewServeMux()
	mux.HandleFunc("/state", s.handleState)
	mux.HandleFunc("/health", s.handleHealth)

	srv := &http.Server{
		Addr:         s.cfg.Address,
		Handler:      mux,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	if s.cfg.TLSEnabled {
		srv.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	// Shut down gracefully when context is cancelled
	go func() {
		<-ctx.Done()
		shutCtx, shutCancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer shutCancel()
		if err := srv.Shutdown(shutCtx); err != nil {
			logrus.WithField("category", "HTTPListener").Warnf("Shutdown error: %v", err)
		}
	}()

	var err error
	if s.cfg.TLSEnabled {
		logrus.WithField("category", "HTTPListener").Infof("Listening on %s (TLS)", s.cfg.Address)
		err = srv.ListenAndServeTLS(s.cfg.TLSCertFile, s.cfg.TLSKeyFile)
	} else {
		logrus.WithField("category", "HTTPListener").Infof("Listening on %s (plain HTTP)", s.cfg.Address)
		err = srv.ListenAndServe()
	}

	if err != nil && err != http.ErrServerClosed {
		logrus.WithField("category", "HTTPListener").Errorf("Server error: %v", err)
	}
}

// handleState accepts a TargetState payload via POST.
// Authentication is mandatory — requests without valid Basic Auth credentials
// are rejected with 401 before any body is read.
func (s *Server) handleState(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "only POST is supported")
		return
	}

	if !s.authenticate(r) {
		w.Header().Set("WWW-Authenticate", `Basic realm="solace-dsemp-agent"`)
		writeError(w, http.StatusUnauthorized, "authentication required")
		return
	}

	body, err := io.ReadAll(io.LimitReader(r.Body, maxBodyBytes))
	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("failed to read request body: %v", err))
		return
	}

	target, err := state.ParseTargetState(body, s.cfg.AESKey)
	if err != nil {
		writeError(w, http.StatusUnprocessableEntity, fmt.Sprintf("invalid target state: %v", err))
		return
	}

	logrus.WithFields(logrus.Fields{
		"category": "HTTPListener",
		"version":  target.Version,
		"remote":   r.RemoteAddr,
	}).Info("Received target state via HTTP POST")

	s.callback(target)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "accepted", "version": target.Version})
}

// handleHealth returns a simple 200 OK for liveness probes.
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// authenticate validates HTTP Basic Auth credentials using constant-time comparison.
func (s *Server) authenticate(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
		return false
	}

	decoded, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
	if err != nil {
		return false
	}

	parts := strings.SplitN(string(decoded), ":", 2)
	if len(parts) != 2 {
		return false
	}

	userMatch := subtle.ConstantTimeCompare([]byte(parts[0]), []byte(s.cfg.Username))
	passMatch := subtle.ConstantTimeCompare([]byte(parts[1]), []byte(s.cfg.Password))
	return userMatch == 1 && passMatch == 1
}

// writeError sends a JSON error response.
func writeError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

