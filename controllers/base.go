package controllers

import (
	"encoding/json"
	"net/http"
	_"reflect"

	"github.com/astaxie/beego/orm"
	_"github.com/astaxie/beego"

	"golang.org/x/crypto/bcrypt"
	"golaman-pack/lib/modules"
)

//Global variables used around the package of controllers
var ormer = orm.NewOrm()

/*
	Display the object interface{} to JSON structure

	@param load @type interface{}
	@param r @type http.Request
	@param w @type http.ResponseWriter
*/
func GetJSONDisplay(load interface{}, statusCode int, w http.ResponseWriter){
	object, err := json.Marshal(load)
	if err != nil {
		panic(err.Error())
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(object)
}

/*
	Gets the object from Body, which is the Payload and convert it
	to a readable object.

	@param load @type interface{}
	@param r @type http.Request
*/
func ReadJSONPayload(r *http.Request)  (map[string]interface{}, error){
	type PayLoad struct {
		Load map[string]interface{}
	}

	var pay PayLoad
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&pay)
	if err == nil {
		return pay.Load, nil
	}
	return nil, err
}

/*
	Generate a hash key using from plain text 
	see: https://godoc.org/golang.org/x/crypto/bcrypt
	
	@param plainText @type string
	@return string, error
*/
func GenerateBCryptHash(plainText string) (string, error) {
	hashedPlainText, hashedPlainTextErr := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if hashedPlainTextErr == nil {
		return string(hashedPlainText), nil
	}
	return "", hashedPlainTextErr
}

/*
	Check if hash key is matched to the plain text
	see: https://godoc.org/golang.org/x/crypto/bcrypt
	
	@param plainText @type string
	@param cipherText @type string
	@return bool
*/
func IsBCrpytHashMatch(cipherText string, plainText string) bool {
	match := bcrypt.CompareHashAndPassword([]byte(cipherText), []byte(plainText))
	if match == nil {
		return true
	}
	return false
}

func SendEmail() {
	subject := "Get latest Tech News directly to your inbox"
	receiver:= []string{
		"king.caubalejo@valueline.com.ph",
		"caubalejokinglawrence@gmail.com",
		"jeffrey.ver@valueline.com.ph",
		"jeffreyver1521@gmail.com",
		"johnmarkarquitola@gmail.com",
		"johnmark.arquitola@b1g2.io",
	}
	r := mail.NewRequest(receiver, subject)
	r.Send("assets/template/email-body.html", map[string]string{"username": "Conor"})
}