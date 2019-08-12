package main
import(
	"fmt"
	"strings"
)
func main() {
	/* ### CARA PERTAMA ### */
	// yourHobbies("Fadhly","Coding","Badminton","Etc..")

	/* ### CARA KEDUA ### */
	var hobbies = []string {"Coding","Badminton","Etc.."} // slice
	yourHobbies("Fadhly", hobbies...)
}
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies,", ")

	fmt.Printf("Hello my name is %s\n",name)
	fmt.Printf("And my hobbies are %s\n",hobbiesAsString)

}