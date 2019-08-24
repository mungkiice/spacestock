package main
import "fmt"
// Golang Specifics
// 2.1 struct
type Human struct {
	Name string
}

func ChangeName(human Human, to string){
	human.Name = to
}

func ChangeOriginalName(human *Human, to string){
	human.Name = to
}

func main(){
	john := Human{Name: "John Doe"}
	ChangeName(john, "Jane Doe")

	//What does the john.Name prints and why?
    fmt.Printf("%s\n", john.Name)
    // English
	// the program will print "John Doe", 
	// because the ChangeName function did not receive the original variable in the first place,
	// it received only the value of the variable and then create a new variable inside the function with a new memory address to hold the value,
	// so the original variable still has the old value even though we called ChangeName function,
	// to change the original value which located outside the function is by using a pointer as a function parameter and sending the address/pointer of the variable instead,
	// here's an example of how to change the original variable:

	// Bahasa
    // program akan mencetak "John Doe" karena ketika memanggil method ChangeName,
    // method tersebut akan menginisiasi variabel baru untuk menampung nilai dari parameter yang dikirim dengan alamat memori yang berbeda dengan variabel awal,
    // sehingga nilai yang berhasil diubah yaitu nilai dari variabel yang ada didalam method ChangeName,
    // untuk mengubah nilai variabel diluar method diperlukan address/pointer dari variabel tersebut,
    // sehingga method mengubah nilai di alamat memori dari variabel tersebut,
    // berikut contoh cara untuk mengubah variabel aslinya:
    
    ChangeOriginalName(&john, "Jane Doe")
    fmt.Printf("%s\n", john.Name)
}

// 2.2 Please setup a REST service for CRUD
