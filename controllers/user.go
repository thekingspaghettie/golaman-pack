package controllers

import (
	"net/http"
	_"net/smtp"
	_"github.com/astaxie/beego"
	_"github.com/astaxie/beego/orm"

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
	users.LastName		= payLoad["last_name"].(string)
	users.FirstName		= payLoad["first_name"].(string)
	users.MiddleName	= payLoad["middle_name"].(string)

	transaction := ormer.Begin()

	if transaction != nil {
		transaction = ormer.Rollback()
		GetJSONDisplay(map[string]interface{}{
			"dev": transaction,
		}, 500, w)
		return
	}

	lastInsertedId, errUser := ormer.Insert(&users)

	if errUser != nil {
		GetJSONDisplay(map[string]interface{}{
			"dev": errUser,
		}, 500, w)
		return
	}

	transaction = ormer.Commit()
	GetJSONDisplay(map[string]interface{}{
		"dev": "Hi developer",
		"pay": payLoad,
		"users": lastInsertedId,
	}, 200, w)
	return
}

func CreateGuest(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Method == "POST" {
		payLoad, errPayLoad := ReadJSONPayload(r)

		if errPayLoad == nil {
	
			email 		:= payLoad["email"].(string)
			password 	:= payLoad["password"].(string)
			
			hashedPassword, hashedPassErr := GenerateBCryptHash(password)
			if hashedPassErr != nil {
				GetJSONDisplay(map[string]interface{}{
					"message": hashedPassErr,
				}, 500, w)
				return
			}
	
			var guest models.Guest
			guest.Email			= email
			guest.Password		= hashedPassword
	
			transaction := ormer.Begin()
	
			if transaction != nil {
				transaction = ormer.Rollback()
				GetJSONDisplay(map[string]interface{}{
					"message": transaction,
				}, 500, w)
				return
			}
	
			lastInsertedId, errUser := ormer.Insert(&guest)
	
			if errUser != nil {
				GetJSONDisplay(map[string]interface{}{
					"message": errUser,
				}, 500, w)
				return
			}
	
			transaction = ormer.Commit()
			GetJSONDisplay(map[string]interface{}{
				"message": "Hi developer",
				"pay": payLoad,
				"users": lastInsertedId,
			}, 200, w)
			return	
		}
	
		GetJSONDisplay(map[string]interface{}{
			"message": errPayLoad,
		}, 500, w)
		return
	}
	GetJSONDisplay(map[string]interface{}{
		"message": MethodMessage("WM", r.Method),
	}, 500, w)
	return
}

func ReadGuest(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if r.Method == "POST" {
		
		SendEmail()
		payLoad, errPayLoad := ReadJSONPayload(r)
		GetJSONDisplay(map[string]interface{}{
			"message": payLoad,
		}, 200, w)
		return
		var guest models.Guest
	
		if errPayLoad == nil {
	
			email 		:= payLoad["email"].(string)
			password 	:= payLoad["password"].(string)
	
			errGuest := ormer.QueryTable("guest").Filter("email", email).One(&guest, "Password")
			
			if errGuest == nil {
	
				successMessage := "Successfully logged in"
				if !IsBCrpytHashMatch(guest.Password, password) {
					successMessage = "Incorrect username/password."
				}
	
				GetJSONDisplay(map[string]interface{}{
					"message": successMessage,
				}, 200, w)
				return
			}
	
			GetJSONDisplay(map[string]interface{}{
				"message": "Incorrect username/password.",
			}, 500, w)
			return
		}
	
		GetJSONDisplay(map[string]interface{}{
			"message": errPayLoad,
		}, 500, w)
		return	
	}

	GetJSONDisplay(map[string]interface{}{
		"message": MethodMessage("WM", r.Method),
	}, 500, w)
	return
}