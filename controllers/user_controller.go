package controllers

import (
	"fmt"
	"github.com/nikola43/testsocket/models"
	"github.com/nikola43/testsocket/utils"
	"net/http"
)

type UserResult struct {
	User chan models.User `json:"user"`
}

var num = 0

func Hi(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, models.User{FingerPrint: "hola"})
	fmt.Println(num)
}

func Login(w http.ResponseWriter, r *http.Request) {
	num++

	defer func() {
		_ = r.Body.Close()
	}()

	userResult := UserResult{
		User: make(chan models.User),
	}


	err := myPool.Submit(func() {
		user := models.User{FingerPrint: "hola"}
		userResult.User <- user

	})

	utils.HandleError(err)
	utils.RespondWithJSON(w, http.StatusOK, <-userResult.User)
	fmt.Println(num)
}
