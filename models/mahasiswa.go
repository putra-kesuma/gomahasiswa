package models

import (
	"fmt"
	"log"

	"github.com/putra-kesuma/gomahasiswa/services"
)

//struct for product
type mahasiswa struct {
	idMahasiswa   int
	nim           string
	namaMahasiswa string
	jurusan       string
}

//function insert Mahasiswa
func InsertMahasiswa() {
	db, err := services.Connect() //call function connect
	defer db.Close()
	var masukanNama, masukanNim, masukanJurusan string
	fmt.Print("Masukkan NIM = ")
	fmt.Scanln(&masukanNim)
	fmt.Print("Masukkan Nama Mahasiswa = ")
	fmt.Scanln(&masukanNama)
	fmt.Print("Masukkan Jurusan = ")
	fmt.Scanln(&masukanJurusan)

	_, err = db.Exec(`insert into m_mahasiswa (nim,nama_mahasiswa,jurusan) 
	value (?,?,?)`, &masukanNim, &masukanNama, &masukanJurusan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
}

//function show all data Mahasiswa
func DetailMahasiswa() {
	db, err := services.Connect() //call function connect
	defer db.Close()              // kalau sudah open jangan lupa close koneksinya
	data, err := db.Query("select * from m_mahasiswa")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	var result []mahasiswa
	for data.Next() {
		var allMahasiswa = mahasiswa{}
		var err = data.Scan(&allMahasiswa.idMahasiswa, &allMahasiswa.nim, &allMahasiswa.namaMahasiswa, &allMahasiswa.jurusan)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, allMahasiswa)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}

//function delete Mahasiswa
func DeleteMahasiswa() {
	db, err := services.Connect() //call function connect
	defer db.Close()
	var masukanId int
	DetailMahasiswa()
	fmt.Print("Masukkan Id mahasiswa yang ingin di hapus = ")
	fmt.Scanln(&masukanId)

	_, err = db.Exec("delete from m_mahasiswa where id_mahasiswa = ? ", &masukanId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
}
