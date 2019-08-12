/*
	Nested struct adalah anonymous struct yang di embed kesebuah struct, dan deklarasinya langsung didalam struct
	peng-embed
*/
package main 
import "fmt"


type student struct {
	person struct {
		name string
		age int
	}
	grade int
	hobbies []string
}


