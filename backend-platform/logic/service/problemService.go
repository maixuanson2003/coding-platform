package service

import (
	"lietcode/logic/constant"
	"lietcode/logic/dto"
	"lietcode/logic/entity"
	"lietcode/logic/repository"

	"gorm.io/gorm"
)

type ProblemService struct {
	ProblemRepo    *repository.ProblemRepository
	TestCaseRepo   *repository.TestcaseRepository
	SubmissionRepo *repository.SubmissionRepository
}

func NewProblemService(repo *repository.ProblemRepository, testCaseRepo *repository.TestcaseRepository, submissionRepo *repository.SubmissionRepository) *ProblemService {
	return &ProblemService{
		ProblemRepo:    repo,
		TestCaseRepo:   testCaseRepo,
		SubmissionRepo: submissionRepo,
	}
}

func (service *ProblemService) CreateProblem(payLoad dto.ProblemCreate) (*dto.ApiResponse, error) {
	var result entity.Problem
	if len(payLoad.TestCase) <= 0 {
		return nil, constant.ErrInvalidRequest
	}
	if _, ok := constant.Difficult[payLoad.Difficult]; !ok {
		return nil, constant.ErrInvalidRequest
	}
	err := service.ProblemRepo.DataAccess.Transaction(func(db *gorm.DB) error {
		res, errToCreate := service.ProblemRepo.Create(&entity.Problem{
			Category:    payLoad.Category,
			Difficult:   constant.Difficult[payLoad.Difficult],
			Title:       payLoad.Title,
			Content:     payLoad.Content,
			MemoryLimit: payLoad.MemoryLimit,
			TimeLimit:   payLoad.TimeLimit,
		})

		if errToCreate != nil {
			return errToCreate
		}
		testCase := []entity.TestCase{}
		for _, item := range payLoad.TestCase {
			testCase = append(testCase, entity.TestCase{
				ProblemId: res.Id,
				Input:     item.Input,
				Output:    item.Output,
				RuntimeMS: item.RuntimeMS,
				MemoryKB:  item.MemoryKB,
			})
		}
		errCreateTestCase := db.Create(&testCase).Error
		if errCreateTestCase != nil {
			return errCreateTestCase
		}
		result = *res
		return nil
	})
	if err != nil {
		return nil, constant.ErrDatabaseAccess
	}

	return &dto.ApiResponse{
		Message: "add problem success",
		Data: map[string]interface{}{
			"problem_id": result.Id,
		},
		Success: true,
	}, nil

}
func (service *ProblemService) GetListProblem(userId *uint, category *string, difficult *string, title *string) (*dto.ApiResponse, error) {

	problemList, errToGetList := service.ProblemRepo.GetListProblem(category, difficult, title, []string{})
	if errToGetList != nil {
		return nil, constant.ErrDatabaseAccess
	}
	var submissList []entity.Submission
	problemArray := map[uint]bool{}
	if userId != nil {
		submissionList, errToGetListSubmission := service.SubmissionRepo.FindAll(map[string]interface{}{
			"user_id": userId,
			"status":  constant.Status["Accepted"],
		}, []string{})
		if errToGetListSubmission != nil {
			return nil, constant.ErrDatabaseAccess
		}
		submissList = submissionList
		for _, submiss := range submissList {
			problemArray[submiss.ProblemId] = true
		}
	}
	mapper := dto.Mapper[entity.Problem, dto.ProblemResponse]{
		Fields: []string{
			"Id",
			"Category",
			"Difficult",
			"Title",
			"Content",
			"IsDeleted",
			"IsDailyToday",
			"PointDaily",
			"MemoryLimit",
			"TimeLimit",
		},
	}
	problemResponse := []map[string]interface{}{}

	for _, item := range problemList {
		isAccept := false
		if v, ok := problemArray[item.Id]; ok && v {
			isAccept = true
		}
		problemResponse = append(problemResponse, map[string]interface{}{
			"problem":  mapper.EntityToResponse(item),
			"isAccept": isAccept,
		})
	}
	return &dto.ApiResponse{
		Message: "success",
		Data:    problemResponse,
		Success: true,
	}, nil

}
func (servive *ProblemService) GetProblemDetail(id uint) (*dto.ApiResponse, error) {
	problem, errToGetDetail := servive.ProblemRepo.FindOne(map[string]interface{}{
		"id": id,
	}, []string{})
	if errToGetDetail != nil {
		return nil, constant.ErrDatabaseAccess
	}
	mapper := dto.Mapper[entity.Problem, dto.ProblemResponse]{
		Fields: []string{
			"Id",
			"Category",
			"Difficult",
			"Title",
			"Content",
			"IsDeleted",
			"IsDailyToday",
			"PointDaily",
			"MemoryLimit",
			"TimeLimit",
		},
	}
	problemResponse := mapper.EntityToResponse(problem)
	return &dto.ApiResponse{
		Message: "success",
		Data:    problemResponse,
		Success: true,
	}, nil
}
func (service *ProblemService) AddTestCaseToProblem(id uint, testCase dto.TestCaseData) (*dto.ApiResponse, error) {
	resp, errToCreate := service.TestCaseRepo.Create(&entity.TestCase{
		ProblemId: id,
		Input:     testCase.Input,
		Output:    testCase.Output,
		RuntimeMS: testCase.RuntimeMS,
		MemoryKB:  testCase.MemoryKB,
	})
	if errToCreate != nil {
		return nil, constant.ErrDatabaseAccess
	}
	return &dto.ApiResponse{
		Message: "success",
		Data:    resp,
		Success: true,
	}, nil

}
