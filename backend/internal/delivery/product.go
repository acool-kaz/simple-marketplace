package delivery

import (
	"database/sql"
	"encoding/json"
	"errors"
	"main/internal/models"
	"net/http"
)

func (h *Handler) productGetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	products, err := h.service.Product.GetAll()
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			h.errPage(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if err := json.NewEncoder(w).Encode(products); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) productFind(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	query := r.URL.Query()["info"][0]

	products, err := h.service.Product.Find(query)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			h.errPage(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if err := json.NewEncoder(w).Encode(products); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) productCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	userId := h.ctx.Value(userCtx).(int)
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		h.errPage(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.Product.Create(product, userId); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(statusResponse{"ok"}); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) productUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// userId := h.ctx.Value(userCtx).(int)
	// productId, ok := mux.Vars(r)["id"]
	// if !ok {
	// 	h.errPage(w, http.StatusNotFound, fmt.Sprintf("id field not found"))
	// 	return
	// }
	if err := json.NewEncoder(w).Encode(statusResponse{"ok"}); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) productDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// userId := h.ctx.Value(userCtx).(int)
	// productId, ok := mux.Vars(r)["id"]
	// if !ok {
	// 	h.errPage(w, http.StatusNotFound, fmt.Sprintf("id field not found"))
	// 	return
	// }
	if err := json.NewEncoder(w).Encode(statusResponse{"ok"}); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}
