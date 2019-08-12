package main 

import(
	"fmt"
	"io"
	"os"
)

var path = `hei.txt`

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

// membuat file baru
func createFile() {
	// deteksi apakah file sudah ada atau belum
	var _, err = os.Stat(path) // os.Stat mengembalikan 2 data yaitu info tentang path yang dicari dan error(jika ada)
	
	// buat file baru jika belum ada, maka fungsi err bernilai true
	if os.IsNotExist(err) { // 
		var file, err = os.Create(path) // secara default statusnya adalah Open, maka perlu di close
		if isError(err) { return }
		defer file.Close() // statement akan dijalankan di akhir
	} else { 
		fmt.Println("File sudah ada, maaf gabisa")
		return
	}
	
	fmt.Println("===> file berhasil dibuat", path)
}

// edit isi file
func writeFile() {
	// buka file dengan level akses READ & WRITE
	var file, err = os.OpenFile(path, os.O_RDWR, 0644) // file dibuka dengan level akses read & write dengan kode permission 0644
	if isError(err) { return }
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString("halo\n")
	if isError(err) { return }
	_, err = file.WriteString("mari belajar golang\n")
	if isError(err) { return }

	//simpan perubahan
	err = file.Sync()
	if isError(err) { return }

	fmt.Println("===> file berhasil diisi")
}

// membaca file
func readFile() {
	// buka file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
						//     (path yang akan dibuka, level akses, kode permission)
	if isError(err) { return }
	defer file.Close()

	// baca file
	var text = make([]byte, 1024) // bertipe slice byte dengan alokasi elemen 1024, dan bertugas untuk menampung data hasil statement file.Read(), dan 
	
	// proses pembacaan akan dilakukan terus menerus, berurutan dari baris pertama hingga terakhir
	for {
		n, err := file.Read(text)
		// ketika errornya adalah selain io.EOF, maka proses baca file akan terus berlanjut. EOF(End Of File || baris terakhir)
		if err != io.EOF {
			if isError(err) { break }
		}
		fmt.Println(n)
		if n == 0 {
			break
		}
	}
	if isError(err) { return }

	fmt.Println("===> file berhasil dibaca")
	fmt.Println(string(text))
}

// menghapus file
func deleteFile() {
	var err = os.Remove(path)
	if isError(err) { return }

	fmt.Println("===> file berhasil dihapus")
}

func main() {
	createFile()
	writeFile()
	readFile()
	deleteFile()
}