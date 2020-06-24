package models

import (
	"fmt"
	"log"

	"github.com/putra-kesuma/gomahasiswa/services"
)

//struct for Matakuliah
type matakuliah struct {
	idMatakuliah   int
	namaMatakuliah string
	sks            int
}

//function insert Mahasiswa
func InsertMatakuliah() {
	db, err := services.Connect() //call function connect
	defer db.Close()
	var masukanMk string
	var masukanSks int
	fmt.Print("Masukkan Nama Mata Kuliah = ")
	fmt.Scanln(&masukanMk)
	fmt.Print("Masukkan jumlah sks = ")
	fmt.Scanln(&masukanSks)

	_, err = db.Exec(`insert into m_matakuliah (nama_matakuliah,sks) 
	value (?,?,?)`, &masukanMk, &masukanSks)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
}

//function show all data Matakuliah
func DetailMatakuliah() {
	db, err := services.Connect() //call function connect
	defer db.Close()              // kalau sudah open jangan lupa close koneksinya
	data, err := db.Query("select * from m_matakuliah")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	var result []matakuliah
	for data.Next() {
		var allMatakuliah = matakuliah{}
		var err = data.Scan(&allMatakuliah.idMatakuliah, &allMatakuliah.namaMatakuliah, &allMatakuliah.sks)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, allMatakuliah)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}

//function delete Matakuliah
func DeleteMatakuliah() {
	db, err := services.Connect() //call function connect
	defer db.Close()
	var masukanId int
	DetailMatakuliah()
	fmt.Print("Masukkan Id matakuliah yang ingin di hapus = ")
	fmt.Scanln(&masukanId)

	_, err = db.Exec("delete from m_matakuilah where id_matakuliah = ? ", &masukanId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
}
