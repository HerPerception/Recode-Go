package main
import (
	"fmt"
	"os"
	"strconv"
)
func main(){
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	} 
	num, _ := strconv.Atoi(os.Args[1])
	if num < 0 {
		fmt.Println(0)
		return
	}
	number := 0
	for num > 0{
		 if num%2 != 0  {
			number = num
		 }
		}
	// 	}else if i%3 != 0  {
	// 		number += i
	// 	}else if i%5 != 0  {
	// 		number += i
	// 	}else {
	// 		continue
	// 	}
	// }
	fmt.Println(number)
	fmt.Println(4%2)
}