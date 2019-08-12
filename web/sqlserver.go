package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql" // (_) karena kita tidak langsung berinteraksi dengan driver tersebut

type student struct {
	id		string
	name	string
	age		int
	grade	int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar_golang") // koneksi ke database belajar_golang
	// user		 => root
	// password  => 
	// host		 => 127.0.0.1 atau localhost
	// port		 => 3306
	// dbname	 => belajar_golang

	if err != nil {
		return nil, err
	}

	return db, nil
}

// buat func sqlQuery untuk proses pembacaan data dari database
func sqlQuery(age int) {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, grade FROM tb_student WHERE age= ?", age) // func db.Query menghasilkan instance bertipe (sql.*Rows) yang perlu di close ketika sudah tidak digunakan
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student  // tipe elemen struct *student* disiapkan dalam variable *result*, nantinya hasil query ditampung pada variable ini
	
	// lakukan perulangan dengan acuan kondisi *rows.Next()*, perulangan ini dilakukan sebanyak record yang ada pada db
	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade) // method *scan()* milik *sql.Rows* fungsinya untuk mengambil nilai record yang sedang diiterasi, untuk disimpan pada variable pointer

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// record yang didapat kemudian di *append* ke slice *result*
		result = append(result, each) 
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, each := range result {
		fmt.Println(each)
		fmt.Println("===================================================")
	}
}

/*			Membaca 1 record data dengan menggunakan *queryRow()*			*/

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = student{}
	var id = "B001"
	err = db.
		QueryRow("SELECT name, grade, age FROM tb_student WHERE id=?", id).
		Scan(&result.name, &result.grade, &result.age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Name: %s\nGrade: %d\nAge: %d\n", result.name, result.grade, result.age)
	fmt.Println("===================================================")
}

/*			Eksekusi Query dengan menggunakan Prepare()			*/
func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT name, grade, age FROM tb_student WHERE id= ?")
	/*
		Method *prepare()* digunakan untuk deklarasi query, yang mengembalikan objek bertipe sql.*Stmt, dari objek tersebut dipanggil method *queryRow()* beberapa kali dengan isi value *id* berbeda-beda untuk tiap pemanggilannya
	*/
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	
	var result1 = student{}
	stmt.QueryRow("B001").Scan(&result1.name, &result1.grade, &result1.age)
	fmt.Printf("Name: %s\nGrade: %d\nAge: %d\n\n", result1.name, result1.grade, result1.age)

	var result2 = student{}
	stmt.QueryRow("B002").Scan(&result2.name, &result2.grade, &result2.age)
	fmt.Printf("Name: %s\nGrade: %d\nAge: %d\n\n", result2.name, result2.grade, result2.age)
	fmt.Println("===================================================\n")
}

/* Create, Update, dan Delete dengan menggunakan Exec() */
func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO tb_student VALUES (?, ?, ?, ?)", "B006", "Budi Sudarsono", 24, 5)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Insert success !!")

	_, err = db.Exec("UPDATE tb_student SET age= ? WHERE id= ?", 23, "B006")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Update success !!")

	_, err = db.Exec("DELETE from tb_student WHERE id= ?", "B006")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Delete success !!")
}

func main() {
	sqlQuery(18)
	sqlQueryRow()
	sqlPrepare()
	sqlExec()
}