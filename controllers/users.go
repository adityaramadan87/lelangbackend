package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"lelangbackend/helper"
	"lelangbackend/models"
)

type UsersController struct {
	beego.Controller
}

func (u *UsersController) Register() {
	var users models.S_user
	json.Unmarshal(u.Ctx.Input.RequestBody, &users)

	u.Ctx.Input.Bind(&users.Name, "name")
	u.Ctx.Input.Bind(&users.Phone, "phone")
	u.Ctx.Input.Bind(&users.Password, "password")

	users.CreateDate = helper.TimeNow()
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	users.Password = string(hashPassword)

	us, err := models.Register(users)
	if err != nil {
		helper.Response(1, err.Error(), u.Controller)
	} else {
		helper.Response(0, us, u.Controller)
	}
	u.ServeJSON()
}

func (u *UsersController) Login() {
	phone := u.GetString("phone")
	password := u.GetString("password")

	isSuccess, msg := models.Login(phone, password)

	if isSuccess {
		helper.Response(0, msg, u.Controller)
	} else {
		helper.Response(1, msg, u.Controller)
	}

	u.ServeJSON()

}

func (u *UsersController) Update() {
	var users models.S_user
	json.Unmarshal(u.Ctx.Input.RequestBody, &users)

	u.Ctx.Input.Bind(&users.Id, "id")
	u.Ctx.Input.Bind(&users.Phone, "phone")
	u.Ctx.Input.Bind(&users.Email, "email")
	u.Ctx.Input.Bind(&users.Name, "name")

	us, err := models.UpdateUsers(&users)
	if err != nil {
		helper.Response(1, err.Error(), u.Controller)
	} else {
		helper.Response(0, us, u.Controller)
	}

	u.ServeJSON()
}
