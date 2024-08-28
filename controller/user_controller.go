package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/example/dto"
	"github.com/example/usecase"
	"github.com/example/util"
)

type UserController struct {
	usecase *usecase.UserUseCase
}

type Photo struct {
	Id     int           `json:"id"`
	Buffer *bytes.Buffer `json:"buffer"`
}

func NewUserController(us *usecase.UserUseCase) *UserController {
	return &UserController{usecase: us}
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.UserCreate

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	fmt.Print(user)
	result, err := uc.usecase.Save(user)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var user dto.UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(400)
		return
	}

	token, err := uc.usecase.Login(user)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (uc *UserController) AddPhoto(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("photo")

	photo, _ := util.FileToBytes(file)

	path := r.URL.Path

	parts := strings.Split(path, "/")

	id, err := strconv.Atoi(parts[2])

	uc.usecase.AddPhotoMinio(id, photo)

	if err != nil {
		slog.Error(err.Error())
		return
	}

	return
}
