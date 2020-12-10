package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"lelangbackend/helper"
	"lelangbackend/models"
	"log"
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
	if err != nil {
		log.Print(err)
	}
	users.Password = string(hashPassword)

	us, msg, rc := models.Register(users)
	helper.Response(rc, msg, us, u.Controller)
}

func (u *UsersController) Login() {
	phone := u.GetString("phone")
	password := u.GetString("password")

	rc, msg := models.Login(phone, password)

	helper.Response(rc, msg, nil, u.Controller)
}

func (u *UsersController) Update() {
	var users models.S_user
	json.Unmarshal(u.Ctx.Input.RequestBody, &users)

	u.Ctx.Input.Bind(&users.Id, "id")
	u.Ctx.Input.Bind(&users.Phone, "phone")
	u.Ctx.Input.Bind(&users.Email, "email")
	u.Ctx.Input.Bind(&users.Name, "name")

	us, msg, rc := models.UpdateUsers(&users)
	helper.Response(rc, msg, us, u.Controller)
}
