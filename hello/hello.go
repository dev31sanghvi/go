package main

import(
"fmt"
//importing module from different file
"example.com/greetings"
"errors"
)


func main(){

	log.SetPrefix("grettings:")
	log.SetFlags(0)


	
	message,err:=greetings.Hello("bitch")
if err != nil{
	log.Fatal(err)
}
	fmt.Println(message)

}
