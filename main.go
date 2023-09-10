package main

import (
	"fmt"
	"gorillaz-clean-v3/config"
	"gorillaz-clean-v3/helpers"
	"gorillaz-clean-v3/middleware"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	userHandler "gorillaz-clean-v3/user/handler"
	userRepo "gorillaz-clean-v3/user/repo"
	userUseCase "gorillaz-clean-v3/user/usecase"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		helpers.Logger("error", "Error getting env")
	}

	db := config.Init()
	defer db.Close()

	r := mux.NewRouter()
	r.Use(middleware.JwtAuth)

	//user
	userRepo := userRepo.NewUserRepo(db)
	userUsecase := userUseCase.NewUserUseCase(userRepo)
	userHandler.NewUserHandler(r, userUsecase)

	p := os.Getenv("PORT")
	h := r
	s := new(http.Server)
	s.Handler = h
	s.Addr = ":" + p
	fmt.Println("Server run in Port ", s.Addr)
	s.ListenAndServe()

}
