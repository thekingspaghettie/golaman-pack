package controllers

import (
	
)

func MethodMessage(param string, method string) (string) {
	switch param {
		case "WM":
			return "Wrong method, your method is " + method
	}
	return "";
}