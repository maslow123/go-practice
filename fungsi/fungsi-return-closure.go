package main
import "fmt"

func main(){
	var max = 3
	var numbers = []int{2,3,0,4,3,2,0,4,2,0,3}
	var howMany, getNumbers = findMax(numbers,max)
	var theNumbers = getNumbers()

	fmt.Println("Numbers\t:", numbers)
	fmt.Printf("Find\t: %d\n\n", max)

	fmt.Println("Found\t:", howMany) // 9
	fmt.Println("Value\t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}

func findMax(numbers []int, max int)(int,func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}