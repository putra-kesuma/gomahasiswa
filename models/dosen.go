package models

import (
	"fmt"
	"log"

	"github.com/putra-kesuma/gomahasiswa/services"
)

//struct for dosen
type dosen struct {
	idDosen   int
	namaDosen string
	namaMk    string
}

//function insert Dosen
func InsertDosen() {
	db, err := services.Connect() //call function connect
	defer db.Close()
	var masukanNama string
	var masukanIdMK int
	fmt.Print("Masukkan Nama Dosen = ")
	fmt.Scanln(&masukanNama)
	DetailMatakuliah()
	fmt.Print("Masukkan id MK = ")
	fmt.Scanln(&masukanIdMK)
	//tiga validasi ketika memasukkan data Dosen
	if masukanNama == "" {
		fmt.Println("Nama Dosen tidak boleh kosong")
	} else if len(masukanNama) < 3 {
		fmt.Println("panjang nama dosen tidak boleh kurang dari tiga")
	} else if masukanIdMK == 0 {
		fmt.Println("Id matakuliah tidak boleh kosong")
	} else {
		_, err = db.Exec(`insert into m_dosen (nama_dosen,id_matakuliah) 
	value (?,?)`, &masukanNama, &masukanIdMK)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("success")
	}
}

func UpdateDosen() {
	db, err := services.Connect() //call function connect
	defer db.Close()
	var masukanId, masukanIdMK int
	var masukanNama string
	DetailDosen()
	fmt.Print("Masukkan Id dosen yang ingin di edit = ")
	fmt.Scanln(&masukanId)
	fmt.Print("Masukkan Nama Dosen = ")
	fmt.Scanln(&masukanNama)
	DetailMatakuliah()
	fmt.Print("Masukkan id MK = ")
	fmt.Scanln(&masukanIdMK)

	_, err = db.Exec(`update m_dosen set nama_dosen = ?, id_matakuliah = ? where id_dosen = ?;`, &masukanNama, &masukanIdMK, &masukanId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
}

//function show all data Dosen
func DetailDosen() {
	db, err := services.Connect() //call function connect
	defer db.Close()              // kalau sudah open jangan lupa close koneksinya
	data, err := db.Query(`select d.id_dosen,d.nama_dosen, mk.nama_matakuliah  from 
	m_dosen d join m_matakuliah mk on d.id_matakuliah = mk.id_matakuliah;`)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	var result []dosen
	for data.Next() {
		var allDosen = dosen{}
		var err = data.Scan(&allDosen.idDosen, &allDosen.namaDosen, &allDosen.namaMk)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, allDosen)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}

//function delete Dosen
func DeleteDosen() {
	db, err := services.Connect() //call function connect
	defer db.Close()
	var masukanId int
	DetailDosen()
	fmt.Print("Masukkan Id Dosen yang ingin di hapus = ")
	fmt.Scanln(&masukanId)

	_, err = db.Exec("delete from m_dosen where id_dosen = ? ", &masukanId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
}
