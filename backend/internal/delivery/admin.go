package delivery

import (
	"database/sql"
	"encoding/json"
	"errors"
	"main/internal/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) adminSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var input models.SignInInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.errPage(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Admin.SignIn(input.Username, input.Password)
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

func (h *Handler) adminGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	users, err := h.service.User.GetAll()
	if err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) adminCreateUser(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) adminDeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.errPage(w, http.StatusNotFound, err.Error())
		return
	}

	if err := h.service.User.DeleteUser(userId); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{"deleted"}); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) adminUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.errPage(w, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.errPage(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.User.UpdateUser(userId, user); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{"updated"}); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) adminGetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	products, err := h.service.Product.GetAll()
	if err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) adminCreateProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		h.errPage(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Product.Create(product, 0); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{"created"}); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) adminDeleteProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	productId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.errPage(w, http.StatusNotFound, err.Error())
		return
	}

	if err := h.service.Product.Delete(productId); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{"deleted"}); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) adminUpdateProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	productId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.errPage(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		h.errPage(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Product.Update(productId, product); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := json.NewEncoder(w).Encode(statusResponse{"updated"}); err != nil {
		h.errPage(w, http.StatusInternalServerError, err.Error())
		return
	}
}
