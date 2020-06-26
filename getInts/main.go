package main
 
import (
    "fmt"
)
 
func main() {

		
	fmt.Println("Valor inicial")
	var  aInt int
	fmt.Scanf("%d", &aInt)
	fmt.Println(aInt)
	fmt.Println("Valor final")
	var bInt int
	fmt.Scanf("%d", &bInt)
	fmt.Println(bInt)


	n := aInt
	
	suma :=0

	for  n <= bInt {
		suma += n
		if n%2 != 0 {
		   fmt.Println(n, "is odd")
		}
		n++
    }

	fmt.Println("El ejercicio se aplica desde el número de inicio" , aInt , "hasta el numero final", bInt)
	fmt.Println(suma)
}