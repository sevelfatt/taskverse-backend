package auth

import (
	"net/http"

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

	user, err := LoginService(body.Email, body.Password)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]any{
		"message": "User logged in successfully",
		"user":    user,
	})
}

func GetUserController(w http.ResponseWriter, r *http.Request) {

	var params struct {
		UserId string `json:"user_id"`
	}

	if err := utils.DecodeJSON(r, &params); err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	if params.UserId == "" {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{
			"error": "User ID is required",
		})
		return
	}

	user, err := GetUserService(params.UserId)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]any{
		"message": "User fetched successfully",
		"user":    user,
	})
}