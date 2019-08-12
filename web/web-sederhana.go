package main 

import "fmt"
import "html/template"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hai, apakabar !")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){ 
		var data = map[string] string {
			"Name": "M.Fadhly NR",
			"Message": "Have a nice day !",
		}
		var t, err = template.ParseFiles("template-engine.html") // parsing template, mengembalikan 2 data yaitu instance templatenya dan error(jika ada)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		t.Execute(w, data) // menampilkan hasil parsing ke layar web browser
	}) 

	http.HandleFunc("/index", index) // routing index memiliki 2 parameter, 1 tujuan routing, 2 callback

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}