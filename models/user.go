package models

import (
	_"github.com/astaxie/beego/orm"
)

type Users struct{
	Id 			int
	LastName	string
	FirstName 	string
	MiddleName 	string
}

type Guest struct {
	Id 			int
	Email		string
	Password 	string
}