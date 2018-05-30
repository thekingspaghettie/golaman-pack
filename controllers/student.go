package controllers

import (
	"io/ioutil"
	"github.com/astaxie/beego"
	"fmt"
	"os"
	"net/http"

	_"golaman-pack/models"
)

var path	 	= "assets/storage/student.txt"
var content 	= ""

func init(){
	beego.Warn("HI")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func ReadStudents(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){

	// GetJSONDisplay(map[string]interface{}{
	// 	"hi": "King Lawrence",
	// 	"POGI": "akop",
	// },r,w)
	
	b, err := ioutil.ReadFile(path)
	if err != nil {
		beego.Warn(err)
	}
	content = string(b)

	// print content
	beego.Warn(string(b))
}

func CreateStudent(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	beego.Warn(ReadJSONPayload(r))
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

	// append the previous content
	err1 := os.Truncate(path, 0)
	if err1 != nil {
		beego.Warn(err1)
	}

	_, err = file.WriteString(content)

	// write some text to file
	_, err = file.WriteString("Id: 3\n")
	_, err = file.WriteString("Lastname: Cabiten\r\n")
	_, err = file.WriteString("Firstname: Ehddver\r\n\n")
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// save changes
	err = file.Sync()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func DeleteStudent(){
	
}

func UpdateStudent(){
	
}