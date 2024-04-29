package model

import "github.com/astaxie/beego/orm"

type BeegoUser struct {
	Id    int //默认主健为id
	Name  string
	Phone string
}

func init() {
	orm.RegisterModel(new(BeegoUser))
}
