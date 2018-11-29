
package main
 
import "fmt"
 
func main(){
    fmt.Print("Enter number : ")
    var n1 int
    fmt.Scanln(&n1)
       fmt.Print("Enter number : ")
    var n2 int
    fmt.Scanln(&n2)
	fmt.Print("Sum:" , n1+n2)
	fmt.Print("Difference:",n1-n2)
	fmt.Print("Multiplication:",n1*n2)
		if(n2==0){
				fmt.Print("Invalid Operation")
	}else{
	fmt.Print("Division:",n1/n2)
	}
}
