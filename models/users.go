package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type S_user struct {
	Id           int    `form:"id" json:"id" pg:"id"`
	Name         string `form:"name" json:"name" pg:"name"`
	Email        string `form:"email" json:"email"`
	Phone        string `form:"phone" json:"phone"`
	Password     string `form:"password" json:"password"`
	Verified     bool   `form:"verified" json:"verified"`
	Ktp          string `form:"ktp" json:"ktp"`
	CreateDate   string `form:"create_date" json:"create_date"`
	VerifiedDate string `form:"verified_date" json:"verified_date"`
}

func init() {
	orm.RegisterModel(new(S_user))
}

func Register(u S_user) (us *S_user, err error) {
	o := orm.NewOrm()
	o.Using("default")

	var userAvailable S_user
	o.Raw("SELECT * FROM s_user WHERE phone = ?", u.Phone).QueryRow(&userAvailable)
	if userAvailable != (S_user{}) {
		return nil, errors.New("number phone already exists")
	}

	id, error := o.Insert(&u)
	if error != nil {
		return nil, errors.New(error.Error())
	}

	var user S_user
	errQuery := o.Raw("SELECT * FROM s_user WHERE id = ?", id).QueryRow(&user)
	if errQuery != nil {
		return nil, errors.New("Error while query " + errQuery.Error())
	}

	return &user, nil
}

func Login(phone string, password string) (isSuccess bool, msg string) {
	o := orm.NewOrm()
	o.Using("default")

	log.Print("he " + phone + "  " + password)

	var users S_user
	if err := o.Raw("SELECT * FROM s_user WHERE phone = ?", phone).QueryRow(&users); err != nil {
		log.Print(errors.New("Error while query " + err.Error()))
		return false, "Error while query " + err.Error()
	}

	if users == (S_user{}) {
		return false, "user not found please make sure your phone number is correct"
	}

	err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if err != nil {
		return false, "Password not match"
	} else {
		return true, "Success"
	}

}
