package delivery

import (
	"database/sql"
	"encoding/json"
	"errors"
	"main/internal/models"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.errPage(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Auth.SignUp(user); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{"created"}); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var input signInInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errPage(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Auth.SignIn(input.Username, input.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.errPage(w, http.StatusBadRequest, err.Error())
			return
		}
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]interface{}{"token": token}); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}
