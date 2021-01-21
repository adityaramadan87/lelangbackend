package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Session struct {
	Id        int
	Token     string    `json:"token"`
	IsValid   bool      `json:"is_valid"`
	ExpiresAt time.Time `json:"expires_at"`
	UserId    int       `json:"user_id"`
}

type SessionClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func init() {
	orm.RegisterModel(new(Session))
}

func CreateToken(userId int) (*Session, error) {
	o := orm.NewOrm()
	o.Using("default")

	var expAT = 24 * 30 * time.Hour
	key := os.Getenv("KEY_LELANG")

	claims := SessionClaims{
		userId,
		jwt.StandardClaims{
			Issuer:    "api",
			ExpiresAt: time.Now().Add(expAT).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return nil, err
	}

	data := Session{
		Token:     tokenString,
		IsValid:   true,
		ExpiresAt: time.Now().Add(expAT),
		UserId:    userId,
	}

	if _, err := o.Insert(&data); err != nil {
		return nil, err
	} else {
		return &data, nil
	}

}
