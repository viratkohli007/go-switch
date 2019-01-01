package main;
import "fmt"
import "time"

func main()  {
	 a:=1
	fmt.Println("print ",a," as ")
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
switch time.Now().Weekday(){
case time.Saturday, time.Sunday:
	  fmt.Println("This is weekend")
default:
	 fmt.Println("This is weekday")

}

}