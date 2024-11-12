package main
import "fmt"


var rappers []string

func main(){
	// declaration of variable
// var taskone="read book"
//declaring the array
var listitems=[]string{"read book","learn golang","learn about monitoring services"}

	// fmt.Println("lizt of my todos")
	// fmt.Println(taskone)
	// fmt.Println("learn golang ")
	// fmt.Println("learn about monitoring servicess ")

	// fmt.Println() //this will print a new line
	// // this is how we interpret a number in go

	// fmt.Println("the number is ",59); // just dont use the double quotes

	fmt.Println(listitems)


	var dds="yo yo honey singh"
	var pro="raftar"
	rappers=[]string{dds,pro}


}


func newMain() {
	fmt.Println("best rappers:",rappers)

    // this is how we can use a for loop
   for index,name:=range rappers{
	fmt.Println(index,name)
	// fmt.Printf(format:"%d. %s\n",index+1,name)
    }
}
