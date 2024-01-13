package handler

import (
	"context"
	"forum/internal/models"
	"net/http"
	"time"
)

func (h *Handler) middleWareGetUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		c, err := r.Cookie("token")
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", models.User{})))
			return
		}
		user, err = h.Service.GetUserByToken(c.Value)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", models.User{})))
			return
		}
		if user.ExpiresAt.Before(time.Now()) {
			if err := h.Service.DeleteToken(c.Value); err != nil {
				h.ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}
		user.IsAuth = true
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", user)))
	}
}

// func (h *Handler) middleWareGetUser(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var user models.User
// 		c, err := r.Cookie("token")
// 		if err != nil {
// 			// Токен отсутствует или ошибка при извлечении
// 			log.Println("Token not found or error extracting token:", err)
// 			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 			return
// 		}

// 		user, err = h.Service.GetUserByToken(c.Value)
// 		if err != nil {
// 			// Ошибка при извлечении пользователя по токену
// 			log.Println("Error retrieving user by token:", err)
// 			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 			return
// 		}

// 		if user.ExpiresAt.Before(time.Now()) {
// 			if err := h.Service.DeleteToken(c.Value); err != nil {
// 				// Ошибка при удалении устаревшего токена
// 				log.Println("Error deleting expired token:", err)
// 				h.ErrorPage(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 				return
// 			}

// 			// Токен устарел, перенаправляем на страницу входа
// 			http.Redirect(w, r, "/signin", http.StatusSeeOther)
// 			return
// 		}

// 		user.IsAuth = true
// 		// Помечаем пользователя как авторизованного и добавляем его в контекст
// 		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", user)))
// 	}
// }
