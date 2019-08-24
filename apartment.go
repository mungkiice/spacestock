package main
import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

type Apartment struct {
	ID int
	Name string
	Address string
}

var db *sql.DB

func init(){
	tmpDB, err := sql.Open("postgres", "user=postgres password=secret dbname=spacestock")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}
}

func getApartment() ([]Apartment){
	apartments := []Apartment{}
	rows, err := db.Query("SELECT * FROM apartments")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		var address string
		err = rows.Scan(&id, &name, &address)
		if err != nil{
			currentApartment := Apartment{ID: id, Name:name, Address:address}
			log.Fatal(err)
			apartments = append(apartments, currentApartment)
		}
	}
	return apartments
}

func saveApartment(name, address string) (int){
	var apartmentID int
	err := db.QueryRow(`INSERT INTO apartments(name, address) VALUES($1, $2) RETURNING id`, name, address).Scan(&apartmentID)

	if err != nil {
		return 0
	}

	fmt.Printf("Last inserted ID: %v\n", apartmentID)
	return apartmentID
}

func updateApartment(id int, name, address string) (int) {
	res, err := db.Exec(`UPDATE books set name=$1, address=$2 where id=$3 RETURNING id`, name, address, id)
	if err != nil {
		return 0
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0
	}

	return int(rowsUpdated)
}
func removeApartment(apartmentID int) (int){
	res, err := db.Exec(`delete from apartments where id = $1`, apartmentID)
	if err != nil {
		return 0
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0
	}

	return int(rowsDeleted)
}