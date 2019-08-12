package main 
import "fmt"

type personn struct { name string; age int; hobbies []string }

func main() {
	// atau bisa juga seperti ini inisialisasinya
	var p1 = struct { name string; age int } { age: 22, name: "Fadhly" }
	var p2 = struct { name string; age int } { "Rizqi", 23}

	fmt.Println(p1)
	fmt.Println(p2)
}

