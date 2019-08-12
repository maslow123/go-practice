package main 

import "fmt"
import "net/url"

func main() {
	var urlString = "http://kalipare.com:80/hello?name=mfadhly nr&age=17"
	var u, e=  url.Parse(urlString) // parsing string ke bentuk url, mengembalikan 2 data yaitu *variable objek bertipe url.URL* dan *error(jika ada)*

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url : %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme) // http
	fmt.Printf("host: %s\n", u.Host) // kalipare.com:80
	fmt.Printf("path: %s\n", u.Path) // /hello
	// secara otomatis query pada url akan di parsing juga menjadi bentuk *map[string][]string*
	var name = u.Query()["name"][0] // mfadhly nr
	var age = u.Query()["age"][0] // 17
	fmt.Printf("name: %s, age: %s\n", name, age)
}