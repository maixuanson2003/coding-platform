package service

import (
	"fmt"
	"lietcode/logic/constant"
	"lietcode/logic/dto"
	"lietcode/logic/entity"
	"lietcode/logic/execute"
	"lietcode/logic/repository"
	"log"
	"reflect"
	"time"
)

type SubmissionService struct {
	SubmissRepo *repository.SubmissionRepository
	UserRepo    *repository.UserRepository
	ProblemRepo *repository.ProblemRepository
}

var SubmisService *SubmissionService

func NewUserSubmissionService(repo *repository.SubmissionRepository, UserRepo *repository.UserRepository,
	ProblemRepo *repository.ProblemRepository) *SubmissionService {
	SubmisService = &SubmissionService{
		SubmissRepo: repo,
		UserRepo:    UserRepo,
		ProblemRepo: ProblemRepo,
	}
	return &SubmissionService{
		SubmissRepo: repo,
		UserRepo:    UserRepo,
		ProblemRepo: ProblemRepo,
	}
}

var LangMap = map[string]bool{
	"cpp":    true,
	"java":   true,
	"js":     true,
	"python": true,
}

func (service *SubmissionService) SaveSubmissionRecord(userId uint, problemId uint, submission dto.Submission) (*dto.ApiResponse, error) {
	submissionRepo := service.SubmissRepo
	userRepo := service.UserRepo
	problemRepo := service.ProblemRepo

	if !LangMap[submission.Lang] {
		return nil, constant.ErrNotFoundItem
	}
	user, errGetUser := userRepo.FindOne(map[string]interface{}{
		"id": userId,
	}, []string{})
	if errGetUser != nil {
		return nil, constant.ErrDatabaseAccess
	}
	if reflect.DeepEqual(user, entity.User{}) {
		return nil, constant.ErrNotFoundItem
	}
	problem, errToGetProblem := problemRepo.FindOne(map[string]interface{}{
		"id": problemId,
	}, []string{"Testcases"})
	if errToGetProblem != nil {
		return nil, constant.ErrDatabaseAccess
	}
	if reflect.DeepEqual(problem, entity.Problem{}) {
		return nil, constant.ErrNotFoundItem
	}
	res, errToCreateSubmiss := submissionRepo.Create(&entity.Submission{
		UserId:    userId,
		ProblemId: problemId,
	})
	if errToCreateSubmiss != nil {
		return nil, constant.ErrDatabaseAccess
	}
	fileName := fmt.Sprintf("submission_%d_%d.txt", res.Id, time.Now().Unix())
	codeExcute := execute.CodeExecuteWorker{}
	file, errSetUp := codeExcute.SetUpToWorker(execute.CodeExecuteConfig{
		Lang:     submission.Lang,
		Code:     submission.Code,
		FileName: fileName,
	})
	if errSetUp != nil {
		log.Print(errSetUp)
		return nil, constant.ErrServerError
	}

	errToPublishmsg := execute.WorkerMQInstance.PublishMessageJSON(map[string]interface{}{
		"UserId":       userId,
		"SubmissionID": res.Id,
		"ProblemId":    problemId,
		"Language":     submission.Lang,
		"Testcase":     problem.Testcases,
		"FileName":     file,
		"SourceCode":   submission.Code,
	})
	if errToPublishmsg != nil {
		return nil, constant.ErrServerError
	}
	return &dto.ApiResponse{
		Message: "submiss success",
		Data: map[string]interface{}{
			"submission": res.Id,
		},
		Success: true,
	}, nil

}
func (service *SubmissionService) GetListSubmission(userId *uint, problemId *uint, lang *string) (*dto.ApiResponse, error) {
	searchData := map[string]interface{}{}

	if userId != nil {
		searchData["user_id"] = *userId
	}

	if problemId != nil {
		searchData["problem_id"] = *problemId
	}

	if lang != nil {
		searchData["lang"] = *lang
	}
	mapper := dto.Mapper[entity.Submission, dto.SubmissionResponse]{
		Fields: []string{
			"Lang",
			"Code",
			"Status",
			"RuntimeMS",
			"MemoryKB",
		},
	}

	submission, errToGetListSubmiss := service.SubmissRepo.FindAll(searchData, []string{})
	if errToGetListSubmiss != nil {
		return nil, constant.ErrDatabaseAccess
	}
	responseList := []dto.SubmissionResponse{}

	for _, item := range submission {
		responseList = append(responseList, mapper.EntityToResponse(item))
	}
	return &dto.ApiResponse{
		Message: "get list success",
		Data:    responseList,
		Success: true,
	}, nil
}
func (service *SubmissionService) GetSubmissionDetail(Id uint) (*dto.ApiResponse, error) {
	submission, errToGetSubmission := service.SubmissRepo.FindOne(map[string]interface{}{
		"id": Id,
	}, []string{})
	if errToGetSubmission != nil {
		return nil, constant.ErrDatabaseAccess
	}
	mapper := dto.Mapper[entity.Submission, dto.SubmissionResponse]{
		Fields: []string{
			"Lang",
			"Code",
			"Status",
			"RuntimeMS",
			"MemoryKB",
		},
	}
	submissResponse := mapper.EntityToResponse(submission)

	return &dto.ApiResponse{
		Message: "get list success",
		Data:    submissResponse,
		Success: true,
	}, nil
}
func (service *SubmissionService) GetSubmissionHistory(userId uint) (*dto.ApiResponse, error) {
	response := map[time.Time]int{}

	Submission, errToGetListSubmission := service.SubmissRepo.FindAll(map[string]interface{}{
		"user_id": userId,
	}, []string{})
	if errToGetListSubmission != nil {
		return nil, constant.ErrDatabaseAccess
	}
	for _, item := range Submission {
		response[item.CreatedAt]++
	}
	handle := []map[string]interface{}{}
	for key, count := range response {
		handle = append(handle, map[string]interface{}{
			"date":          key,
			"numberSubmiss": count,
		})
	}
	return &dto.ApiResponse{
		Message: "get submiss history",
		Data:    handle,
		Success: true,
	}, nil

}
