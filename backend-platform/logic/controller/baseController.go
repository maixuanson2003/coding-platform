package controller

import (
	"encoding/json"
	"net/http"
)

type BaseController struct{}

func (bc *BaseController) JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
func (bc *BaseController) Error(statusCode int, message string) {

}
