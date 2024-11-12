package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string)(string,errors) {
	if name == ""{
		return "", errors.New("empty name")
		//errors.New is a function 

}

message:= fmt.Sprintf("yo bitch",name)
return message,nil
