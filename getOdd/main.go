package main

import "fmt"

func main() {
	var a [1000]int
	sum := 0
	contador := 1

	for contador <= 1000 {
		sum += contador
		a[contador - 1] = contador
		if contador%2 != 0 {
		   fmt.Println(contador, "is odd")}
		contador++
	
	}
	fmt.Println(sum)

}
