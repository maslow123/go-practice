package main 

import(
	"encoding/json"
	"net/http"
	"fmt"
)

type student struct { // digunakan sebagai tipe elemen slice sample data
	ID		string
	Name	string
	Grade	int
}

var data = []student { // variabel ini digunakan untuk menampung data dari struct student
	student{"E001", "Fadhly", 17},
	student{"E002", "Rasya", 12},
	student{"E003", "Hilman", 21},
	student{"E004", "Jono", 41},
}

// buat function *users* untuk handle endpoint */users*
func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // menentukan tipe response, yaitu sebagai JSON

	// jika request adalah post, maka data yang di-encode ke JSON dijadikan sebagai response
	if r.Method == "POST" { 
		var result, err = json.Marshal(data) 

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

// buat func *user* untuk menampilkan satu buah data saja, diambil berdasarkan IDnya
func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var id = r.FormValue("id") // digunakan untuk mengambil data form yang dikirim oleh client
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id { // jika ada id yang dicari, maka dikembalikan sebagai response
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}
		// jika tidak ada maka error 404, Bad Request dan dikembalikan dengan pesan User Not Found
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}