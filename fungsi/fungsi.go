package main

import "fmt"
import "strings"

func main() {
	var names = [] string {"John","Wick"}
	printMessage("Hello",names)
}

func printMessage(message string, arr[]string) {
	var nameString = strings.Join(arr," ") // menggabungkan slice 1 dengan slice 2 dan dipisahkan oleh koma
	fmt.Println(message, nameString)
}