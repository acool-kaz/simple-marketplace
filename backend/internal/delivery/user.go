package delivery

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) profilePage(w http.ResponseWriter, r *http.Request) {
	userId := h.ctx.Value(userCtx).(int)
	user, err := h.service.User.GetUserById(userId)
	if err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}

type userFind struct {
	Username string `json:"username"`
}

func (h *Handler) findUsers(w http.ResponseWriter, r *http.Request) {
	var find userFind
	if err := json.NewDecoder(r.Body).Decode(&find); err != nil {
		h.errPage(w, http.StatusBadRequest, err.Error())
		return
	}
	users, err := h.service.User.GetUsers(find.Username)
	if err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(&users); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}
