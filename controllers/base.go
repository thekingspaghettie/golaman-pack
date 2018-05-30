package controllers

import (
	"encoding/json"
	"net/http"
	_"github.com/astaxie/beego"
)

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