package main

import (
	"fmt"
	"net/http"
)


var shortGolang = " Go crash course"
var fullGolang = "add Go to my skillset"
var rewardDessert = "Reward myself with a donut"

var taskItems = []string{shortGolang, fullGolang, rewardDessert}
func main(){
	fmt.Println("My Task List")

	http.HandleFunc()
	http.ListenAndServe(":8080", nil)



}