/*
	Tag biasa dilakukan untuk keperluan encode/decode data json. Informasi tag juga bisa diakses lewat reflect,
*/

package main 
import "fmt"

type person struct {
	name string `tag1`
	age int `tag2`
}

