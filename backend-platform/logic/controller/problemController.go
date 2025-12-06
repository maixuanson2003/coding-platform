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

type ProblemController struct {
	Serve *service.ProblemService
	*BaseController
}

func NewProblemController(serve *service.ProblemService) *ProblemController {
	return &ProblemController{
		Serve:          serve,
		BaseController: &BaseController{},
	}
}

func (controller *ProblemController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/problem", middleware.Handle(controller.CreateProblem)).Methods("POST")
	router.HandleFunc("/problem", middleware.Handle(controller.GetListProblem)).Methods("GET")
	router.HandleFunc("/problem/{id}", middleware.Handle(controller.GetProblemDetail)).Methods("GET")
	router.HandleFunc("/problem/{id}/testcase", middleware.Handle(controller.AddTestCaseToProblem)).Methods("POST")
}
func (controller *ProblemController) CreateProblem(w http.ResponseWriter, r *http.Request) error {
	var body dto.ProblemCreate

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err
	}

	dataResp, err := controller.Serve.CreateProblem(body)
	if err != nil {
		return err
	}

	controller.JSON(w, http.StatusOK, dataResp)
	return nil
}
func (controller *ProblemController) GetListProblem(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()

	var userId *uint
	var category *string
	var difficult *string
	var title *string

	// userId = ?user_id=123
	if v := query.Get("user_id"); v != "" {
		parsed, _ := strconv.Atoi(v)
		val := uint(parsed)
		userId = &val
	}

	if v := query.Get("category"); v != "" {
		category = &v
	}

	if v := query.Get("difficult"); v != "" {
		difficult = &v
	}

	if v := query.Get("title"); v != "" {
		title = &v
	}

	// ⭐ truyền đúng thứ tự param
	dataResp, err := controller.Serve.GetListProblem(userId, category, difficult, title)
	if err != nil {
		return err
	}

	controller.JSON(w, http.StatusOK, dataResp)
	return nil
}
func (controller *ProblemController) GetProblemDetail(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	dataResp, err := controller.Serve.GetProblemDetail(uint(id))
	if err != nil {
		return err
	}

	controller.JSON(w, http.StatusOK, dataResp)
	return nil
}
func (controller *ProblemController) AddTestCaseToProblem(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	var body dto.TestCaseData
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return err
	}

	dataResp, err := controller.Serve.AddTestCaseToProblem(uint(id), body)
	if err != nil {
		return err
	}

	controller.JSON(w, http.StatusOK, dataResp)
	return nil
}
