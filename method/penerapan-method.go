/*
	Method adalah fungsi yang menempel pada type(bisa struct ataupun yang lainnya). Method bisa diakses lewat variable
	objek. Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct hingga level private
*/

package main 
import(
	"fmt"
	"strings"
)

type students struct {
	name string
	grade int
}

func main() {
	var s1 = students{"M Fadhly NR",17}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("Nama panggilan :", name)
}

func (s students) sayHello() {
	fmt.Println("Halo", s.name)
}

func (s students) getNameAt(i int) string {
	return strings.Split(s.name," ")[i-1]
}