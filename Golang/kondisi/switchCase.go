package main 

import "fmt"

/* func main(){
	var point = 1

	switch point {
	case 8:
		fmt.Printf("Perfect")

	case 7: 
		fmt.Printf("awosome")
	
	default: 
		fmt.Printf("not bad")
	}
} */

/* func main(){
	var point =  10

	switch point {
	case 8:
		fmt.Printf("Perfect")

	case 7,6,5,4: // banyak sekaaligus
		fmt.Printf("awosome")
	
	default: 
		fmt.Printf("not bad")
	}
} */

var point = 10
	func main() {
		if point > 7 {
			switch point {
			case 10: 
				fmt.Println("perfect");
			
			default: 
				fmt.Println("nice")
			}
		} else {
			if point == 5 {
				fmt.Println("not bad")
			} else if point == 3 {
				fmt.Println("keep trying")
			} else {
				fmt.Println("you can do it")
				if point == 0 {
					fmt.Println("try harder")
				}
			}
		}
}