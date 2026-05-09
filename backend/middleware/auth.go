package middleware

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"
)

/* gerado com auxilio do gemini */
/* TEMPORÁRIO: Mudar regras para diferenciar diferentes usuários */
func AuthMiddleware(app *firebase.App, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// client de autenticação do firebase
		authClient, err := app.Auth(context.Background())
		if err != nil {
			http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
			return
		}

		// extrai o header "Authorization: Bearer <TOKEN>"
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return
		}

		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		// validação da assinatura do token
		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			http.Error(w, "Sessão expirada ou token inválido", http.StatusUnauthorized)
			return
		}

		// usuario legitimo
		// injeta o UID no context para que o handler saiba quem tá pedindo
		ctx := context.WithValue(r.Context(), "user_id", token.UID)
		
		// deixa a requisição passar
		next(w, r.WithContext(ctx))
	}
}