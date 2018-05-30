package controllers

import (
	"net/http"
	_"github.com/astaxie/beego"
	"golaman-pack/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	payLoad, errPayLoad := ReadJSONPayload(r)

	if errPayLoad != nil {	
		GetJSONDisplay(map[string]interface{}{
			"dev-message": errPayLoad,
		}, 500, w)
		return
	}
	var users models.Users
	users.Lastname		= payLoad["last_name"].(string)
	users.Firstname		= payLoad["first_name"].(string)
	users.Middlename	= payLoad["middle_name"].(string)

	GetJSONDisplay(map[string]interface{}{
		"dev": "Hi developer",
		"pay": payLoad,
	}, 200, w)
	return
}