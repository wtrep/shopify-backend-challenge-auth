package auth

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wtrep/shopify-backend-challenge-auth/common"
	"net/http"
	"os"
)

type Handler struct {
	db *sql.DB
}

func CheckEnvVariables() {
	env := []string{"DB_IP", "DB_PASSWORD", "DB_USERNAME", "DB_NAME", "JWT_KEY"}
	for _, e := range env {
		_, ok := os.LookupEnv(e)
		if !ok {
			panic("fatal: environment variable " + e + " is not set")
		}
	}
}

func SetupAndServeRoutes() {
	CheckEnvVariables()

	db, err := NewConnectionPool()
	if err != nil {
		panic(err)
	}
	handler := Handler{db: db}

	r := mux.NewRouter()
	r.HandleFunc("/user", handler.HandlePostUser).Methods("POST")
	r.HandleFunc("/key", handler.HandleGetKey).Methods("GET")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func (h *Handler) HandlePostUser(w http.ResponseWriter, r *http.Request) {
	var request UserRequest
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		common.RespondWithError(w, &common.InvalidRequestBodyError)
		return
	}

	if len(request.Password) > 32 {
		common.RespondWithError(w, &common.PasswordTooLongError)
		return
	}
	user, err := NewUser(request.Name, request.Password)
	if err != nil {
		common.RespondWithError(w, &common.DatabaseInsertionError)
		return
	}

	err = CreateUser(h.db, *user)
	if err == UserAlreadyExist {
		common.RespondWithError(w, &common.UserAlreadyExistError)
		return
	} else if err != nil {
		common.RespondWithError(w, &common.DatabaseInsertionError)
		return
	}

	sendJWTToken(w, user)
}

func (h *Handler) HandleGetKey(w http.ResponseWriter, r *http.Request) {
	var request UserRequest
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		common.RespondWithError(w, &common.InvalidRequestBodyError)
		return
	}

	currentUser, err := GetUser(h.db, request.Name)
	if err != nil {
		common.RespondWithError(w, &common.UserDoesNotExistError)
		return
	}

	if currentUser.goodPassword(request.Password) {
		sendJWTToken(w, currentUser)
	} else {
		common.RespondWithError(w, &common.WrongPasswordError)
	}
}

func sendJWTToken(w http.ResponseWriter, user *User) {
	token, err := common.GenerateJWT(user.Username)
	if err != nil {
		common.RespondWithError(w, &common.TokenGenerationError)
		return
	}
	response := UserResponse{Token: token}

	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		common.RespondWithError(w, &common.JSONEncoderError)
	}
}
