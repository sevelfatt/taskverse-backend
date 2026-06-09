package task

import (
	"net/http"
	"time"

	"github.com/sevelfatt/taskverse-backend/utils"
)

func DeleteTaskByUUIDController(w http.ResponseWriter, r *http.Request) {
	taskUUID := r.URL.Query().Get("task_uuid")
	if taskUUID == "" {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Task UUID Required",
		})
		return
	}

	err := DeleteTaskByUUIDService(taskUUID)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{
		"message": "task deleted successfully",
	})
}

func GetAllTasksByUserUUIDController(w http.ResponseWriter, r *http.Request) {
	tokenString, err := utils.GetTokenFromHeader(r)
	if err != nil {
		utils.RespondJSON(w, http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		return
	}

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

	tasks, err := GetAllTasksByUserUUIDService(userUUID)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]any{
		"message": "Tasks fetched successfully",
		"tasks":   tasks,
	})
}

func CreateTaskController(w http.ResponseWriter, r *http.Request) {
	tokenString, err := utils.GetTokenFromHeader(r)
	if err != nil {
		utils.RespondJSON(w, http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		return
	}

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

	var body struct {
		Title string `json:"title"`
		Type string `json:"type"`
		DaysInWeek []time.Weekday `json:"days_in_week"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
	}

	if err := utils.DecodeJSON(r, &body); err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	task, err := CreateTaskService(userUUID, body.Title, body.Type, body.DaysInWeek, body.StartDate, body.EndDate)
	if err != nil {
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]any{
		"message": "Task created successfully",
		"task":    task,
	})
}