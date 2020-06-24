package models

import (
	"fmt"
	"log"

	"github.com/putra-kesuma/gomahasiswa/services"
)

//struct for Nilai
type nilai struct {
	nim           string
	namaMahasiswa string
	jurusan       string
	namaMk        string
	sks           int
	nilai         int
}

//function insert Nilai
func InsertNilai() {
	db, err := services.Connect() //call function connect
	defer db.Close()
	var masukanIdMahasiswa, masukanIdMK, nilai int
	DetailMahasiswa()
	fmt.Print("Masukkan Id Mahasiswa yang ingin di nilai = ")
	fmt.Scanln(&masukanIdMahasiswa)
	//menampilkan data mata kuliah yang di ambil
	var nim string
	err = db.QueryRow("select nim from m_mahasiswa where id_mahasiswa = ?", &masukanIdMahasiswa).Scan(&nim)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Berikut list Mata kuliah yang sudah di ambil")
	MatakuliahByNim(nim)
	fmt.Print("Masukkan Id Matakuliah = ")
	fmt.Scanln(&masukanIdMK)
	fmt.Print("Masukkan Nilai = ")
	fmt.Scanln(&nilai)

	_, err = db.Exec(`insert into m_nilai (id_mahasiswa,id_matakuliah,nilai) 
	value (?,?,?)`, &masukanIdMahasiswa, &masukanIdMK, &nilai)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
}

//function show all data nilai
func DetailNilai() {
	db, err := services.Connect() //call function connect
	defer db.Close()              // kalau sudah open jangan lupa close koneksinya
	data, err := db.Query("select * from view_penilaian")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	var result []nilai
	for data.Next() {
		var allNilai = nilai{}
		var err = data.Scan(&allNilai.nim, &allNilai.namaMahasiswa, &allNilai.jurusan, &allNilai.namaMk, &allNilai.sks, &allNilai.nilai)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, allNilai)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}

//function delete Nilai
func DeleteNilai() {
	db, err := services.Connect() //call function connect
	defer db.Close()
	var masukanId int
	DetailNilai()
	fmt.Print("Masukkan Id nilai yang ingin di hapus = ")
	fmt.Scanln(&masukanId)

	_, err = db.Exec("delete from m_nilai where id_nilai = ? ", &masukanId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
}
