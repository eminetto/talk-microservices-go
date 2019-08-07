package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/eminetto/talk-microservices-go/pkg/security"
)

func IsJWTAuthenticated() negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		errorMessage := "Erro na autenticação"
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			err := errors.New("Unauthorized")
			respondWithError(rw, http.StatusUnauthorized, err.Error(), errorMessage)
			return
		}
		token, err := security.ParseToken(tokenString)
		if err != nil {
			respondWithError(rw, http.StatusUnauthorized, err.Error(), errorMessage)
			return
		}

		claims, err := security.GetClaims(token)
		if err != nil {
			err := errors.New("Unauthorized")
			respondWithError(rw, http.StatusUnauthorized, err.Error(), errorMessage)
			return
		}
		r.Header.Add("email", claims["email"])

		next(rw, r)
	}
}

//RespondWithError return a http error
func respondWithError(w http.ResponseWriter, code int, e string, message string) {
	respondWithJSON(w, code, map[string]string{"code": strconv.Itoa(code), "error": e, "message": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
