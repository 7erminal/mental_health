package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Authentication struct {
	Password string `orm:"size(255)"`
	Email    string `orm:"size(255)"`
}

func init() {
	orm.RegisterModel(new(Authentication))
}
