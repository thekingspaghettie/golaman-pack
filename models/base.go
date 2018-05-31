package models

import (
	"github.com/astaxie/beego/orm"
)

/*
	Notes:
		If the table fields in your database are named using "snake case"
		like < my_sample_field > , you'll name it in your model like this
		"MySampleField".
*/


//Initialization of models
func init() {
	orm.RegisterDataBase("default", "mysql", "root:admin@/golaman?charset=utf8", 30)

	orm.RegisterModel(
		new(Users), 
		new(Guest),
	)
}