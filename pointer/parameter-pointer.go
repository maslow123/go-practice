package main
import "fmt"

func main() {
	var number = 4
	fmt.Println("Before :", number) // 4

	change(&number, 10)
	fmt.Println("After :", number) // 10
}

func change(original *int, value int) {
	*original = value
}