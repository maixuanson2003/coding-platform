package middleware

import (
	"encoding/json"
	"lietcode/logic/constant"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func Handle(h Handler) func(http.ResponseWriter, *http.Request) {
	return (func(w http.ResponseWriter, r *http.Request) {

		if err := h(w, r); err != nil {
			code, ok := constant.StatusCodes[err]
			if !ok || code == 0 {
				code = http.StatusInternalServerError
			}
			description, ok := constant.Descriptions[err]
			if !ok || description == "" {
				description = "Internal Server Error"
			}
			errorResponse := map[string]interface{}{
				"message": err.Error(),
				"success": false,
				"data": map[string]interface{}{
					"description": description,
					"error_code":  code,
				},
			}

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(code)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}
	})
}
