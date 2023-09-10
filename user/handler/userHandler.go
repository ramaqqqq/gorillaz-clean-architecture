package handler

import (
	"encoding/json"
	"gorillaz-clean-v3/helpers"
	"gorillaz-clean-v3/models"
	"gorillaz-clean-v3/user"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userUseCase user.UserUseCase
}

func NewUserHandler(ctm *mux.Router, userUseCase user.UserUseCase) {

	middleUrl := os.Getenv("MIDDLE_URL")
	userHandlers := UserHandler{userUseCase}
	ctm.HandleFunc(middleUrl+"/login", userHandlers.H_Login).Methods("POST")
	ctm.HandleFunc(middleUrl+"/register", userHandlers.H_Register).Methods("POST")

}

func (e *UserHandler) H_Login(w http.ResponseWriter, r *http.Request) {
	datum := models.User{}
	err := json.NewDecoder(r.Body).Decode(&datum)
	if err != nil {
		helpers.Logger("error", "In Server: handler user, "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	result, err := e.userUseCase.Login(&datum)
	if err != nil {
		helpers.Logger("error", "In Server: handler user, "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "user: "+string(logger))
	helpers.Response(w, http.StatusOK, rMsg)
}

func (e *UserHandler) H_Register(w http.ResponseWriter, r *http.Request) {
	datum := models.User{}
	err := json.NewDecoder(r.Body).Decode(&datum)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	result, err := e.userUseCase.Create(&datum)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "creates user: "+string(logger))
	helpers.Response(w, http.StatusCreated, rMsg)
}
