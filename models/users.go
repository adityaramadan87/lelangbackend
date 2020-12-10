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

func Register(u S_user) (us *S_user, msg string, rc int) {
	o := orm.NewOrm()
	o.Using("default")

	var userAvailable S_user
	o.Raw("SELECT * FROM s_user WHERE phone = ?", u.Phone).QueryRow(&userAvailable)
	if userAvailable != (S_user{}) {
		return nil, "number phone already exists", 1
	}

	id, error := o.Insert(&u)
	if error != nil {
		return nil, error.Error(), 1
	}

	var user S_user
	errQuery := o.Raw("SELECT * FROM s_user WHERE id = ?", id).QueryRow(&user)
	if errQuery != nil {
		return nil, "Error while query " + errQuery.Error(), 1
	}

	return &user, "Success", 0
}

func Login(phone string, password string) (rc int, msg string) {
	o := orm.NewOrm()
	o.Using("default")

	log.Print("he " + phone + "  " + password)

	var users S_user
	if err := o.Raw("SELECT * FROM s_user WHERE phone = ?", phone).QueryRow(&users); err != nil {
		log.Print(errors.New("Error while query " + err.Error()))
		return 1, "Error while query " + err.Error()
	}

	if users == (S_user{}) {
		return 1, "user not found please make sure your phone number is correct"
	}

	err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if err != nil {
		return 1, "Password not match"
	} else {
		return 0, "Success"
	}

}

func UpdateUsers(users *S_user) (data *S_user, msg string, rc int) {
	if string(users.Id) != "" {
		o := orm.NewOrm()
		o.Using("default")

		var us S_user
		if er := o.Raw("SELECT * FROM s_user WHERE id = ?", users.Id).QueryRow(&us); er != nil {
			return nil, "Error while query " + er.Error(), 1
		}

		if us == (S_user{}) {
			return nil, "Data with id " + string(users.Id) + " not found", 1
		}

		if users.Phone != "" {
			us.Phone = users.Phone
		}
		if users.Name != "" {
			us.Name = users.Name
		}
		if users.Email != "" {
			us.Email = users.Email
		}

		_, erro := o.Update(&us)
		if erro != nil {
			return nil, "Error when Update data " + erro.Error(), 1
		}

		var userDone S_user
		if err := o.Raw("SELECT * FROM s_user WHERE id = ?", users.Id).QueryRow(&userDone); err != nil {
			return nil, "Error when select user " + err.Error(), 1
		}

		return &userDone, "Success", 0

	}

	return nil, "Id not null", 1

}
