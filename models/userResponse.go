package models

// import (
// 	"github.com/beego/beego/v2/client/orm"
// )

type UserResponse struct {
	StatusCode int    `orm: "omitempty"`
	User       *Users `orm: "omitempty"`
	StatusDesc string `orm:"size(255)"`
}

// func init() {
// 	orm.RegisterModel(new(UserResponse))
// }
