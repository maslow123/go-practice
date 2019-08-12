package main 

import(
	"crypto/sha1"
	"fmt"
	"time"
)

func doHashUsingSalt(text string) (string, string) { // mengembalikan 2 nilai string
	var salt = fmt.Sprintf("%d", time.Now().UnixNano()) // hasilnya akan selalu unix setiap detiknya, karena scope terendah waktu pada fungsi time adalah *nano second*
	var saltedText = fmt.Sprintf("Text: '%s', salt: '%s'", text, salt)
	fmt.Println(saltedText)

	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func main() {
	var text = "My password"
	fmt.Printf("original : %s\n\n",text)
    

	var hashed1, salt1 = doHashUsingSalt(text)
	fmt.Printf("hashed : %s\n\n", hashed1)

	var hashed2, salt2 = doHashUsingSalt(text)
	fmt.Printf("hashed 2 : %s\n\n", hashed2)

	var hashed3, salt3 = doHashUsingSalt(text)
	fmt.Printf("hashed 3 : %s\n\n", hashed3)

	
	_, _, _ = salt1, salt2, salt3 // deklarasi salt 1, 2, 3

}