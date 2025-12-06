package main

import (
	"lietcode/database"
	"lietcode/logic/controller"
	"lietcode/logic/execute"
	"lietcode/logic/repository"
	"lietcode/logic/service"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	godotenv.Load()
	config := "root:root123@tcp(127.0.0.1:3307)/appdb?charset=utf8mb4&parseTime=True&loc=Local"
	database.InitDatabase(config)
	mainRouter := router.PathPrefix("/api").Subrouter()
	Cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:3001"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With", "Accept"},
		AllowCredentials: true,
		Debug:            true,
	})

	// ✅ Đúng cách để áp dụng CORS
	handler := Cors.Handler(router)
	log.Print(mainRouter)
	mainRouter.HandleFunc("/event/{user_id}/{problem_id}/{submiss_id}", execute.EventsHandler)
	userRepo := repository.NewUserRepository(database.DatabaseInstance)
	submissionRepo := repository.NewSubmissionRepository(database.DatabaseInstance)
	testCaseRepo := repository.NewTestcaseRepository(database.DatabaseInstance)
	problemRepo := repository.NewProblemRepository(database.DatabaseInstance)

	userService := service.NewUserService(userRepo)
	submissService := service.NewUserSubmissionService(submissionRepo, userRepo, problemRepo)
	problemService := service.NewProblemService(problemRepo, testCaseRepo, submissionRepo)

	authController := controller.NewAuthController(service.NewUserService(userRepo))
	authController.RegisterRoutes(mainRouter)

	userController := controller.NewUserController(userService)
	userController.RegisterRoutes(mainRouter)
	submissionController := controller.NewSubmissionController(submissService)
	submissionController.RegisterRoutes(mainRouter)
	problemController := controller.NewProblemController(problemService)
	problemController.RegisterRoutes(mainRouter)
	execute.Create(submissionRepo)

	go func() {
		worker := execute.WorkerMQInstance
		Consummer := worker.HandleQueueSubmiss()
		for _, function := range Consummer {
			go function()
		}
		select {}
	}()

	log.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", handler)
}
