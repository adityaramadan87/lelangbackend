package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"image/jpeg"
	"image/png"
	"lelangbackend/helper"
	"lelangbackend/models"
	"log"
	"strconv"
	"strings"
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
	var users models.S_user
	json.Unmarshal(u.Ctx.Input.RequestBody, users)
	//
	u.Ctx.Input.Bind(&users.Phone, "phone")
	u.Ctx.Input.Bind(&users.Password, "password")
	//if err := u.ParseForm(&users); err != nil {
	//	log.Print(err.Error())
	//	helper.Response(1, err.Error(), nil, u.Controller)
	//	return
	//}

	rc, msg, ss := models.Login(users.Phone, users.Password)

	helper.Response(rc, msg, ss, u.Controller)
}

func (u *UsersController) Update() {
	var users models.S_user
	json.Unmarshal(u.Ctx.Input.RequestBody, &users)

	u.Ctx.Input.Bind(&users.Id, "id")
	u.Ctx.Input.Bind(&users.Avatar, "avatar")
	u.Ctx.Input.Bind(&users.Email, "email")
	u.Ctx.Input.Bind(&users.Name, "name")

	us, msg, rc := models.UpdateUsers(&users)
	helper.Response(rc, msg, us, u.Controller)
}

func (u *UsersController) GetUserAvatar() {
	imgID, _ := strconv.Atoi(u.GetString(":avatarid"))

	o := orm.NewOrm()
	o.Using("default")

	var img string
	if err := o.Raw("SELECT avatar FROM user_avatar WHERE id = ?", imgID).QueryRow(&img); err != nil {
		helper.Response(1, "Error while get Avatar "+err.Error(), nil, u.Controller)
		return
	}

	coI := strings.Index(img, ",")
	rawImage := img[coI+1:]

	// Encoded Image DataUrl //
	unbased, _ := base64.StdEncoding.DecodeString(string(rawImage))
	res := bytes.NewReader(unbased)

	switch strings.TrimSuffix(img[5:coI], ";base64") {
	case "image/png":
		pngImg, _ := png.Decode(res)

		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, pngImg, nil); err != nil {
			log.Println("unable to encode image.")
		}

		u.Ctx.Request.Header.Set("Content-Type", "image/png")
		u.Ctx.Request.Header.Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

		if _, err := u.Ctx.ResponseWriter.Write(buffer.Bytes()); err != nil {
			log.Print("errrrrorr", err.Error())
		}
	case "image/jpeg":
		jpgImg, _ := jpeg.Decode(res)

		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, jpgImg, nil); err != nil {
			log.Println("unable to encode image.")
		}

		u.Ctx.Request.Header.Set("Content-Type", "image/png")
		u.Ctx.Request.Header.Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

		if _, err := u.Ctx.ResponseWriter.Write(buffer.Bytes()); err != nil {
			log.Print("errrrrorr", err.Error())
		}
	}
}
