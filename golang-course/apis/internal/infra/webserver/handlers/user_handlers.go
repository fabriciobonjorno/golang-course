package handlers

import (
	"apis/internal/dto"
	"apis/internal/entity"
	"apis/internal/infra/database"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
)

type UserHandlers struct {
	// Define user handler methods here
	UserDB database.UserInterface
}

func NewUserHandlers(userDB database.UserInterface) *UserHandlers {
	return &UserHandlers{
		UserDB: userDB,
	}
}

func (h *UserHandlers) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)    // after value add type of context
	jwtExpireIn := r.Context().Value("jwtExpireIn").(int) // after value add type of context
	var user dto.GetJwtInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		http.Error(w, "User or Password invalid", http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(user.Password) {
		http.Error(w, "User or Password invalid", http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpireIn)).Unix(),
	})

	// Function anonymous
	accessToken := struct {
		AcessToken string `json:"access_token"`
	}{
		AcessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

	// _, token, _ := h.Jwt.Encode(map[string]interface{}{"sub": u.ID})
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(map[string]string{"access_token": token})
}

func (h *UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a user
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}
