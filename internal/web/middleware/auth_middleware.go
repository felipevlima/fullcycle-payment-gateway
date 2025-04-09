package middleware

import (
	"net/http"

	"github.com/devfullcycle/imersao22/go-gateway/internal/domain"
	"github.com/devfullcycle/imersao22/go-gateway/internal/service"
)

type AuthMiddleware struct {
	accountService *service.AccountService
}

func NewAccountMiddleware(accountService *service.AccountService) *AuthMiddleware {
	return &AuthMiddleware{
		accountService: accountService,
	}
}

func (m *AuthMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" {
			http.Error(w, "X-API-KEY is required", http.StatusUnauthorized)
			return
		}

		_, err := m.accountService.FindByAPIKey(apiKey)
		if err != nil {
			if err == domain.ErrAccountNotFound {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
