package delivery

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) profilePage(w http.ResponseWriter, r *http.Request) {
	userId := h.ctx.Value(ctx).(int)
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
