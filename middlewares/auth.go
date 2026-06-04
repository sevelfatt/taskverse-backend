package middlewares

import (
	"net/http"

	"github.com/sevelfatt/taskverse-backend/utils"
)
func AuthMiddelware(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := utils.GetTokenFromHeader(r)
		if err != nil {
			utils.RespondJSON(w, http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
			return
		}

		_, err = utils.ValidateAndGetJwtTokenClaims(tokenString)
		if err != nil {
			utils.RespondJSON(w, http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
			return
		}
		next(w, r)
	}

	return fn
}