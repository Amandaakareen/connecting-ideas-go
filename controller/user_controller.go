package controller

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/example/dto"
	resterr "github.com/example/restErr"
	"github.com/example/usecase"
	"github.com/example/util"
	"github.com/gin-gonic/gin"
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

func (uc *UserController) Create(r *gin.Context) {
	var user dto.UserCreate

	err := r.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		resErr := resterr.NewBadRequestError("requisição inválida ")

		r.JSON(resErr.Code, resErr)
		return
	}

	_, err = uc.usecase.Save(user)
	if err != nil {
		resErr := resterr.NewBadRequestError(err.Error())

		r.JSON(resErr.Code, resErr)
		return
	}

}

func (uc *UserController) Login(r *gin.Context) {
	var user dto.UserLogin

	if err := r.ShouldBindJSON(user); err != nil {
		resErr := resterr.NewBadRequestError("requisição inválida ")

		r.JSON(resErr.Code, resErr)
		return
	}

	token, err := uc.usecase.Login(user)
	if err != nil {
		resErr := resterr.NewBadRequestError(err.Error())

		r.JSON(resErr.Code, resErr)
		return
	}

	r.JSON(http.StatusOK, token)
}

func (uc *UserController) AddPhoto(r *gin.Context) {
	file, _, err := r.Request.FormFile("photo")
	if err != nil {
		resErr := resterr.NewBadRequestError("requisição inválida")

		r.JSON(resErr.Code, resErr)
		return
	}

	//pendente; ver no Gin como fazer isso.
	fmt.Print(file)
	photo, _ := util.FileToBytes(file)

	fmt.Print(photo)
	path := r.Request.PathValue("id")

	parts := strings.Split(path, "/")

	id, err := strconv.Atoi(parts[2])

	fmt.Print(id)
	uc.usecase.AddPhotoMinio(id, photo)

	if err != nil {
		slog.Error(err.Error())
	}

}
