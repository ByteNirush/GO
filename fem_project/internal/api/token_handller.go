package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ByteNirush/fem_project/internal/store"
	"github.com/ByteNirush/fem_project/internal/tokens"
	"github.com/ByteNirush/fem_project/internal/utils"
)

type TokenHandler struct {
	tokenStore store.TokenStore
	userStore store.UserStore
	logger *log.Logger
}

type CreateTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewTokenHandler(tokenStore store.TokenStore, userStore store.UserStore, logger *log.Logger) *TokenHandler {
	return &TokenHandler{
		tokenStore: tokenStore,
		userStore: userStore,
		logger: logger,
	}
}

func (h *TokenHandler) HandleCreateToken(w http.ResponseWriter, r *http.Request) {
	var req CreateTokenRequest
	
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Printf("ERROR: createTokenRequest: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request sent"})
		return
	}

	// lets get the user by username
	user, err := h.userStore.GetUserByUsername(req.Username)
	if err != nil || user == nil {
		h.logger.Printf("ERROR: getUserByUsername: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	PasswordsDoMatch, err := user.PasswordHash.Matches(req.Password)
	if err != nil {
		h.logger.Printf("ERROR: PasswordHash.Match %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	if !PasswordsDoMatch {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "invalid credentials"})
		return
	}

	token, err := h.tokenStore.CreateNewToken(user.ID, 24*time.Hour, tokens.ScopeAuth)
	if err != nil {
		h.logger.Printf("ERROR: Creating %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"auth_token": token})
	
}