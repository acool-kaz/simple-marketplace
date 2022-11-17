package delivery

import (
	"context"
	"log"
	"net/http"
	"strings"
)

func (h *Handler) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// w.Header().Set("content-type", "application/json;charset=UTF-8")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

const userCtx = "userId"

func (h *Handler) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header, ok := r.Header["Authorization"]
		if !ok {
			h.errPage(w, http.StatusUnauthorized, "empty auth header")
			return
		}
		headerParts := strings.Split(header[0], " ")
		if len(headerParts) != 2 {
			h.errPage(w, http.StatusUnauthorized, "invalid auth header")
			return
		}
		userId, err := h.service.Auth.ParseToken(headerParts[1])
		if err != nil {
			h.errPage(w, http.StatusUnauthorized, err.Error())
			return
		}
		h.ctx = context.WithValue(h.ctx, userCtx, userId)
		next.ServeHTTP(w, r)
	})
}

const adminCtx = "userId"

func (h *Handler) adminIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header, ok := r.Header["Authorization"]
		if !ok {
			h.errPage(w, http.StatusUnauthorized, "empty auth header")
			return
		}
		headerParts := strings.Split(header[0], " ")
		if len(headerParts) != 2 {
			h.errPage(w, http.StatusUnauthorized, "invalid auth header")
			return
		}
		adminId, err := h.service.Admin.ParseToken(headerParts[1])
		if err != nil {
			h.errPage(w, http.StatusUnauthorized, err.Error())
			return
		}
		h.ctx = context.WithValue(h.ctx, adminCtx, adminId)
		next.ServeHTTP(w, r)
	})
}
