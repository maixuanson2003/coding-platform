package controller

import (
	"encoding/json"
	"lietcode/logic/dto"
	"lietcode/logic/middleware"
	"lietcode/logic/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController struct {
	Serve *service.UserService
	*BaseController
}

func NewUserController(serve *service.UserService) *UserController {
	return &UserController{
		Serve:          serve,
		BaseController: &BaseController{},
	}
}

func (controller *UserController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/user/register", middleware.Handle(controller.Register)).Methods("POST")
	router.HandleFunc("/user/list", middleware.Handle(controller.GetListUser)).Methods("GET")
	router.HandleFunc("/user/{id}", middleware.Handle(controller.GetUserById)).Methods("GET")
}

// ========================
//
//	REGISTER USER
//
// ========================
func (controller *UserController) Register(w http.ResponseWriter, r *http.Request) error {
	var body dto.Register

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err
	}

	dataResp, err := controller.Serve.UserRegister(body)
	if err != nil {
		return err
	}

	controller.JSON(w, http.StatusOK, dataResp)
	return nil
}

func (controller *UserController) GetListUser(w http.ResponseWriter, r *http.Request) error {
	dataResp, err := controller.Serve.GetListUser()
	if err != nil {
		return err
	}

	controller.JSON(w, http.StatusOK, dataResp)
	return nil
}

func (controller *UserController) GetUserById(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	dataResp, err := controller.Serve.GetUserById(uint(id))
	if err != nil {
		return err
	}

	controller.JSON(w, http.StatusOK, dataResp)
	return nil
}
