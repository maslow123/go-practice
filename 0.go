package main

import(
	"fmt"
	"strings"
)

func main() {
	alphabet("omama")
}

func alphabet(word string) {
	words := strings.Replace(word," ","",-1)
	arrWord := strings.Split(words,"")
	vocal := []string{"a","i","u","e","o"}
	
	var arrVocal []string
	var arrConst []string

	for _, value := range arrWord {
		for i,v := range vocal {
			if value == v {
				if len(arrVocal) != 0 {
					for u,y := range arrVocal {
						if value == y {
							break
						} else if u == len(arrVocal) - 1 {
							arrVocal = append(arrVocal, value)
						}
					}
				}else {
					arrVocal = append(arrVocal, value)
				}
				break
			}else if i == len(vocal) -1 {
				arrConst = append(arrConst, value)
			}
		}
	}
	
	fmt.Println("Huruf mati :",len(arrConst))
	fmt.Println("Huruf hidup:",len(arrVocal))
}