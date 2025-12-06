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

type SubmissionController struct {
	Serve *service.SubmissionService
	*BaseController
}

func NewSubmissionController(serve *service.SubmissionService) *SubmissionController {
	return &SubmissionController{
		Serve:          serve,
		BaseController: &BaseController{},
	}
}

func (controller *SubmissionController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/submission", middleware.Handle(controller.CreateSubmission)).Methods("POST")
	router.HandleFunc("/submission", middleware.Handle(controller.GetListSubmission)).Methods("GET")
	router.HandleFunc("/submission/{id}", middleware.Handle(controller.GetSubmissionDetail)).Methods("GET")
}
func (controller *SubmissionController) CreateSubmission(write http.ResponseWriter, req *http.Request) error {
	userId := req.URL.Query().Get("user_id")
	problemId := req.URL.Query().Get("problem_id")
	parseUserId, _ := strconv.Atoi(userId)
	parseProblemId, _ := strconv.Atoi(problemId)

	var body dto.Submission

	errToParseJson := json.NewDecoder(req.Body).Decode(&body)
	if errToParseJson != nil {
		return errToParseJson
	}
	response, err := controller.Serve.SaveSubmissionRecord(uint(parseUserId), uint(parseProblemId), body)
	if err != nil {
		return err
	}
	controller.JSON(write, 200, response)
	return nil
}
func (controller *SubmissionController) GetListSubmission(w http.ResponseWriter, r *http.Request) error {

	query := r.URL.Query()

	var userId *uint
	var problemId *uint
	var lang *string

	if q := query.Get("user_id"); q != "" {
		parsed, _ := strconv.Atoi(q)
		val := uint(parsed)
		userId = &val
	}

	if q := query.Get("problem_id"); q != "" {
		parsed, _ := strconv.Atoi(q)
		val := uint(parsed)
		problemId = &val
	}

	if q := query.Get("lang"); q != "" {
		lang = &q
	}

	dataResp, err := controller.Serve.GetListSubmission(userId, problemId, lang)
	if err != nil {
		return err
	}

	controller.JSON(w, http.StatusOK, dataResp)
	return nil
}
func (controller *SubmissionController) GetSubmissionDetail(w http.ResponseWriter, r *http.Request) error {

	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	dataResp, err := controller.Serve.GetSubmissionDetail(uint(id))
	if err != nil {
		return err
	}

	controller.JSON(w, http.StatusOK, dataResp)
	return nil
}
