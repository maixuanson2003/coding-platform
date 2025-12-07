package controller

import (
	"context"
	"encoding/json"
	"io"
	"lietcode/logic/config"
	"lietcode/logic/dto"
	"lietcode/logic/middleware"
	"lietcode/logic/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

type AuthController struct {
	GitHubOAuthConfig *oauth2.Config
	GoogleOAuthConfig *oauth2.Config
	Serve             *service.UserService
	*BaseController
}

func NewAuthController(serve *service.UserService) *AuthController {
	return &AuthController{
		GitHubOAuthConfig: config.GitHubOAuthConfig,
		GoogleOAuthConfig: config.GoogleOAuthConfig,
		Serve:             serve,
		BaseController:    &BaseController{},
	}
}

func (controller *AuthController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/auth/google/login", middleware.Handle(controller.GoogleLogin)).Methods("GET")
	router.HandleFunc("/auth/google/callback", middleware.Handle(controller.GoogleCallback)).Methods("GET")
	router.HandleFunc("/auth/github/login", middleware.Handle(controller.GithubLogin)).Methods("GET")
	router.HandleFunc("/auth/github/callback", middleware.Handle(controller.GitHubCallback)).Methods("GET")
	router.HandleFunc("/auth/basic/login", middleware.Handle(controller.UserLogin)).Methods("POST")
}
func (controller *AuthController) GoogleLogin(w http.ResponseWriter, r *http.Request) error {
	urlLogin := controller.GoogleOAuthConfig.AuthCodeURL("state-token")
	log.Println("Redirecting to:", urlLogin)
	controller.JSON(w, http.StatusAccepted, map[string]interface{}{
		"url": urlLogin,
	})
	return nil
}
func (controller *AuthController) GoogleCallback(w http.ResponseWriter, r *http.Request) error {
	codeGoogleSend := r.URL.Query().Get("code")

	token, err := controller.GoogleOAuthConfig.Exchange(context.Background(), codeGoogleSend)
	if err != nil {
		return err
	}
	client := controller.GoogleOAuthConfig.Client(context.Background(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)

	var res map[string]interface{}
	json.Unmarshal(data, &res)

	ok, _ := controller.Serve.UserRegister(dto.Register{
		Username: res["name"].(string),
		Password: "randomPassword",
		Email:    res["email"].(string),
	})
	log.Print(ok)

	dataResp, errToLogin := controller.Serve.UserLogin(res["email"].(string), nil)
	if errToLogin != nil {
		return errToLogin
	}
	dataMap, _ := dataResp.Data.(map[string]interface{})

	jwtToken, _ := dataMap["token"].(string)
	userResponse := dataMap["user"].(dto.UserResponse)
	userIDStr := strconv.Itoa(int(userResponse.Id))
	redirectURL := "http://localhost:3000/login?token=" + jwtToken + "&" + "user_id=" + userIDStr
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
	return nil

}
func (controller *AuthController) GithubLogin(w http.ResponseWriter, r *http.Request) error {
	urlLogin := controller.GitHubOAuthConfig.AuthCodeURL("state-token")
	log.Println("Redirecting to:", urlLogin)
	controller.JSON(w, http.StatusAccepted, map[string]interface{}{
		"url": urlLogin,
	})
	return nil
}
func (controller *AuthController) GitHubCallback(w http.ResponseWriter, r *http.Request) error {
	code := r.URL.Query().Get("code")

	token, err := controller.GitHubOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return err
	}

	client := controller.GitHubOAuthConfig.Client(context.Background(), token)

	resp, err := client.Get("https://api.github.com/user")
	log.Print(resp)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)

	var res map[string]interface{}
	json.Unmarshal(data, &res)
	// Auto register user
	ok, _ := controller.Serve.UserRegister(dto.Register{
		Username: res["login"].(string),
		Password: "randomPassword",
		Email:    res["notification_email"].(string),
	})
	log.Print(ok)

	dataResp, errLogin := controller.Serve.UserLogin(res["notification_email"].(string), nil)
	if errLogin != nil {
		return errLogin
	}

	dataMap, _ := dataResp.Data.(map[string]interface{})

	jwtToken, _ := dataMap["token"].(string)

	userResponse := dataMap["user"].(dto.UserResponse)
	userIDStr := strconv.Itoa(int(userResponse.Id))

	redirectURL := "http://localhost:3000/login?token=" + jwtToken + "&" + "user_id=" + userIDStr
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)

	return nil
}
func (controller *AuthController) UserLogin(w http.ResponseWriter, r *http.Request) error {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")

	resp, err := controller.Serve.UserLogin(email, &password)
	if err != nil {
		return err
	}
	controller.JSON(w, 200, resp)
	return nil
}
