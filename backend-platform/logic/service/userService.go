package service

import (
	"fmt"
	"lietcode/logic/auth"
	"lietcode/logic/constant"
	"lietcode/logic/dto"
	"lietcode/logic/entity"
	"lietcode/logic/repository"
	"log"
	"reflect"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) UserRegister(register dto.Register) (*dto.ApiResponse, error) {
	userRepo := service.repo

	exist, err := userRepo.ExsistUserEmail(register.Email)
	if err != nil {
		return nil, constant.ErrDatabaseAccess
	}
	if exist {
		return nil, constant.ErrEmailAlreadyExists
	}
	bytes, errToEncrypt := bcrypt.GenerateFromPassword([]byte(register.Password), 14)
	if errToEncrypt != nil {
		log.Printf("can complete encrypt code")
		return nil, constant.ErrRuntimeError
	}
	newUser, errToCreateData := userRepo.Create(&entity.User{
		Username: register.Username,
		Password: string(bytes),
		Email:    register.Email,
		Avatar:   "check",
	})
	if errToCreateData != nil {
		log.Printf("err to access database")
		return nil, constant.ErrDatabaseAccess
	}
	return &dto.ApiResponse{
		Message: "create data success",
		Data: map[string]interface{}{
			"id":       newUser.Id,
			"username": newUser.Username,
			"email":    newUser.Email,
		},
		Success: true,
	}, nil

}
func (service *UserService) UserLogin(email string, Password *string) (*dto.ApiResponse, error) {
	userRepo := service.repo
	user, errToFindUser := userRepo.FindOne(map[string]interface{}{
		"email": email,
	}, []string{})
	if errToFindUser != nil {
		log.Printf("err to access database")
		return nil, constant.ErrDatabaseAccess
	}
	if reflect.DeepEqual(user, entity.User{}) {
		return nil, constant.ErrNotFoundItem
	}
	if Password != nil {
		errComparePassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(*Password))
		if errComparePassword != nil {
			return nil, constant.ErrUnauthorizedClient
		}
	}
	helper := auth.TokenHelper{}
	token, errToGenerateToken := helper.GenerateToken(user.Username, []string{"USER"})
	if errToGenerateToken != nil {
		return nil, constant.ErrRuntimeError
	}

	mapper := dto.Mapper[entity.User, dto.UserResponse]{
		Fields: []string{
			"Id", "Username", "Email", "Avatar", "NumberHandle", "PointDaily"},
	}
	userResponse := mapper.EntityToResponse(user)
	dataResponse := map[string]interface{}{
		"token": token,
		"user":  userResponse,
	}
	return &dto.ApiResponse{
		Message: "login success",
		Data:    dataResponse,
		Success: true,
	}, nil

}

func (service *UserService) GetListUser() (*dto.ApiResponse, error) {
	userRepo := service.repo
	listUser, err := userRepo.FindAll(nil, []string{})

	if err != nil {
		return nil, constant.ErrDatabaseAccess
	}
	mapper := dto.Mapper[entity.User, dto.UserResponse]{
		Fields: []string{
			"Id", "Username", "Email", "Avatar", "NumberHandle", "PointDaily"},
	}
	userResponses := []dto.UserResponse{}

	for _, item := range listUser {
		userResponses = append(userResponses, mapper.EntityToResponse(item))
	}
	return &dto.ApiResponse{
		Message: "get list user success",
		Data:    userResponses,
		Success: true,
	}, nil
}

func (service *UserService) GetUserById(id uint) (*dto.ApiResponse, error) {
	userRepo := service.repo
	database := service.repo.DataAccess
	Lang := []string{"cpp",
		"java",
		"js",
		"python"}

	user, errToFindUser := userRepo.FindOne(map[string]interface{}{
		"id": id,
	}, []string{})
	if errToFindUser != nil {
		log.Printf("err to access database")
		return nil, constant.ErrDatabaseAccess
	}

	selectFields := []string{}
	accepted := constant.Status["Accepted"]

	for _, item := range Lang {
		selectFields = append(selectFields,
			fmt.Sprintf("COUNT(CASE WHEN submissions.lang = '%s' THEN 1 END) AS total_%s", item, item),
		)
	}
	selectFields = append(selectFields,
		fmt.Sprintf("COUNT(CASE WHEN submissions.status = '%s' THEN 1 END) AS accept_problem", accepted),
	)
	selectFields = append(selectFields,
		fmt.Sprintf("COUNT(CASE WHEN problems.difficult = 'easy' AND submissions.status = '%s' THEN 1 END) AS total_easy", accepted),
		fmt.Sprintf("COUNT(CASE WHEN problems.difficult = 'medium' AND submissions.status = '%s' THEN 1 END) AS total_medium", accepted),
		fmt.Sprintf("COUNT(CASE WHEN problems.difficult = 'hard' AND submissions.status = '%s' THEN 1 END) AS total_hard", accepted),
	)

	query := database.Table("submissions").
		Joins("JOIN problems ON submissions.problem_id = problems.id").
		Where("submissions.user_id = ?", id).
		Select(strings.Join(selectFields, ", "))

	type UserStats struct {
		TotalCpp    int64 `json:"total_cpp"`
		TotalJava   int64 `json:"total_java"`
		TotalJs     int64 `json:"total_js"`
		TotalPython int64 `json:"total_python"`

		TotalEasy   int64 `json:"total_easy"`
		TotalMedium int64 `json:"total_medium"`
		TotalHard   int64 `json:"total_hard"`
	}
	var result UserStats
	if err := query.Scan(&result).Error; err != nil {
		return nil, constant.ErrDatabaseAccess
	}
	mapper := dto.Mapper[entity.User, dto.UserResponse]{
		Fields: []string{
			"Id", "Username", "Email", "Avatar", "NumberHandle", "PointDaily"},
	}
	userResponse := mapper.EntityToResponse(user)
	return &dto.ApiResponse{
		Message: "get list user success",
		Data: map[string]interface{}{
			"user":      userResponse,
			"userStats": result,
		},
		Success: true,
	}, nil

}
