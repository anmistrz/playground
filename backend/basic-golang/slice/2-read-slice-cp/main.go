package main

import "fmt"

// Disini kalian coba untuk membaca nilai dari slice yang diberikan.
// Lalu kalian akan menambahkan kata "Olleh" pada slice tersebut.
func main() {
	slice := []string{"Hello", "World"}
	//beginanswer
	slice = append(slice, "Olleh")
	fmt.Println(slice)
	//endanswer
}
