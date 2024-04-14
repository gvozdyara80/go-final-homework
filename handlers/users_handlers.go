package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-final-homework/models"
	"github.com/spf13/viper"

	"github.com/go-final-homework/service"
	"github.com/go-final-homework/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	repository models.UserRepository
}

func NewUserHandler(repository models.UserRepository) *UserHandler {
	return &UserHandler{repository: repository}
}

func (uh *UserHandler) InitUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", uh.RegisterUser).Methods("POST")
	router.HandleFunc("/users/{id}", uh.LoginUser).Methods("POST")
}

func (uh *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserRegisterCredentials
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid credentials: %v", errors))
		return
	}

	// check if user exists
	_, err := uh.repository.GetUserByEmail(user.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	// hash password
	hashedPassword, err := service.EncryptPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = uh.repository.CreateUser(models.User{
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (uh *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserLoginCredentials
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid credentials: %v", errors))
		return
	}

	u, err := uh.repository.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	if !service.ComparePasswords(u.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(viper.GetString("jwt_secret"))
	token, err := service.CreateJWT(secret, u.Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}
