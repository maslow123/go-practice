package main
import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	sortWord("osama")
}

func sortWord(word string) {
	arrWord := strings.Split(strings.Replace(word," ","",-1),"")
	vocal := []string {"a","i","u","e","o"}

	var life, dead []string
	var hasil string

	for _,value := range arrWord {
		for i,v := range vocal {
			if value == v {
				life = append(life, value)
				break
			} else if i == len(vocal) - 1{
				dead = append(dead, value)
			}
		}
	}
	sort.Strings(life)
	sort.Strings(dead)
	hasil = strings.Join(life,"") + strings.Join(dead,"")

	fmt.Println(hasil)
}