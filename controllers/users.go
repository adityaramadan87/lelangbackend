package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"lelangbackend/models"
	"time"
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

	users.CreateDate = time.Now().Format("02/01/2006 15:04:05")
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	users.Password = string(hashPassword)

	us, err := models.Register(users)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = us
	}
	u.ServeJSON()
}

func (u *UsersController) Login() {
	//var phone string
	//var password string

	phone := u.GetString("phone")
	password := u.GetString("password")

	//u.Ctx.Input.Bind(phone, "phone")
	//u.Ctx.Input.Bind(password, "password")

	isSuccess, msg := models.Login(phone, password)

	if isSuccess {
		u.Data["json"] = map[string]string{"rc": "0", "response": msg}
	} else {
		u.Data["json"] = map[string]string{"rc": "1", "response": msg}
	}

	u.ServeJSON()

}
