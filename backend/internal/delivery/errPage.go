package delivery

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func (h *Handler) errPage(w http.ResponseWriter, status int, text string) {
	w.Header().Set("Content-Type", "application/json")
	e := errorResponse{
		Status: status,
		Msg:    text,
	}
	w.WriteHeader(e.Status)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
