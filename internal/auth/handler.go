package auth

import (
	"awesomeproject/configs"
	"awesomeproject/pkg/req"
	"awesomeproject/pkg/res"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func SetupHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyLogin, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		res.Json(w, bodyLogin, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyRegister, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		res.Json(w, bodyRegister, 200)
	}
}
