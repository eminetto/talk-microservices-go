package main

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/eminetto/talk-microservices-go/auth/pkg/security"
	"github.com/eminetto/talk-microservices-go/auth/pkg/user"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	uService := user.NewUserService()
	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.NewLogger(),
	)
	r.Handle("/v1/auth", n.With(
		negroni.Wrap(userAuth(uService)),
	)).Methods("POST", "OPTIONS")
	r.Handle("/v1/validate-token", n.With(
		negroni.Wrap(validateToken()),
	)).Methods("POST", "OPTIONS")
	http.Handle("/", r)
	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":8081",
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func userAuth(uService user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		var param parameters
		err := json.NewDecoder(r.Body).Decode(&param)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		type result struct {
			Token string `json:"token"`
		}
		err = uService.ValidateUser(param.Email, param.Password)
		if err != nil{
			w.WriteHeader(http.StatusForbidden)
			return
		}
		t, err := security.NewToken(param.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res := &result{
			Token: t,
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		return
	})
}

func validateToken() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Token    string `json:"token"`
		}
		var param parameters
		err := json.NewDecoder(r.Body).Decode(&param)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		type result struct {
			Email string `json:"email"`
		}
		t, err := security.ParseToken(param.Token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tData, err := security.GetClaims(t)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res := &result{
			Email: tData["email"].(string),
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	})
}
