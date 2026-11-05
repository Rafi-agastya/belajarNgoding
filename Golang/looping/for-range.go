package main

import "fmt"

/* var xs = "123"
func main() {
	for i, v := range xs {
		fmt.Println("index ke-", i, "value: ", v)
	}
} */

/* var ys = [5]int{10,20,30,40,50} //array
 func main() {
	for _, v := range ys{
		fmt.Println("value: ", v)
	}
} */

/* var zs = ys[0:2] //slice
func main(){
	for _, v := range zs{
		fmt.Println("value: ", v)
	}
} */

// var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} "map"
/* func main() {
	/*for k, v := range kvs{
	fmt.Println("key:", k, "value: ", v)
	}*/

	/*for range kvs{
		fmt.Println("Done")
	} */

	/* for i := range 5{
		fmt.Print(i)
	/* 
	} 
} */

// break & continue
/* func main() {
	for i := 1; i <= 10; i++{
		if i%2 == 1 {
			continue
		} 

		if i > 8 {
			break
		}

		fmt.Println("angka ke-", i)
	}
} */

// bersarang
func main() {
	/* for i := 0; i < 10; i++{
		for j := i; j < 5; j++{
			fmt.Print(j, " ")
		}

	fmt.Println()
	} */

	outerloop: 
	for i := 0; i < 5; i++{
		for j := i; j < 5; j++{
			if i == 3 {
				break outerloop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}


