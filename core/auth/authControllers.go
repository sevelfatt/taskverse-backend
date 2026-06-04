package auth

import (
	"net/http"
	"strings"

	"github.com/sevelfatt/taskverse-backend/utils"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := utils.DecodeJSON(r, &body); err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	if body.Username == "" || body.Email == "" || body.Password == "" {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{
			"error": "All fields are required",
		})
		return
	}

	message, err := RegisterService(body.Username, body.Email, body.Password)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{
		"message": message,
	})

}

func LoginController(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := utils.DecodeJSON(r, &body); err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	if body.Email == "" || body.Password == "" {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{
			"error": "All fields are required",
		})
		return
	}

	tokenString, err := LoginService(body.Email, body.Password)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]any{
		"message": "User logged in successfully",
		"token":    tokenString,
	})
}

func GetUserController(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		utils.RespondJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "Authorization token is required",
		})
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims, err := utils.ValidateAndGetJwtTokenClaims(tokenString)
	if err != nil {
		utils.RespondJSON(w, http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		return
	}

	userUUID, ok := claims["sub"].(string)
	if !ok {
		utils.RespondJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "Invalid token claims",
		})
		return
	}

	user, err := GetUserService(userUUID)
	if err != nil {
		utils.RespondJSON(w, http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]any{
		"message": "User fetched successfully",
		"user":    user,
	})
}