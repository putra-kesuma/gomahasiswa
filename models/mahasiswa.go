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
	var masukanNama, masukanNim, masukanJurusan, option, suboption string
	fmt.Print("Masukkan NIM = ")
	fmt.Scanln(&masukanNim)
	fmt.Print("Masukkan Nama Mahasiswa = ")
	fmt.Scanln(&masukanNama)
	fmt.Print("Masukkan Jurusan = ")
	fmt.Scanln(&masukanJurusan)

	//tiga validasi ketika memasukkan data mahasiswa
	//selanjutnya setiap validasi harus dibuatkan functionnya
	if masukanNim == "" {
		fmt.Println("Nim tidak boleh kosong")
	} else if len(masukanNama) < 3 {
		fmt.Println("panjang Nama karakter tidak boleh kurang dari tiga")
	} else if masukanJurusan == "" {
		fmt.Println("Jurusan tidak boleh kosong")
	} else {

		fmt.Print("tekan y jika ingin menambahkan mata kuliah yang diambil = ")
		fmt.Scanln(&option)
		for {
			if option == "y" {
				InsertTranskip(masukanNim)
				fmt.Print("tekan y jika ingin menambahkan mata kuliah lagi = ")
				fmt.Scanln(&suboption)
				if suboption != "y" {
					break
				}
			}
		}

		_, err = db.Exec(`insert into m_mahasiswa (nim,nama_mahasiswa,jurusan) 
	value (?,?,?)`, &masukanNim, &masukanNama, &masukanJurusan)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Data mahasiswa success disimpan")
	}

}

//function Update Mahasiswa
func UpdateMahasiswa() {
	db, err := services.Connect() //call function connect
	defer db.Close()
	var masukanId int
	var masukanNama, masukanNim, masukanJurusan string
	DetailMahasiswa()
	fmt.Print("Masukkan Id Mahasiswa yang ingin di edit = ")
	fmt.Scanln(&masukanId)
	fmt.Print("Masukkan NIM = ")
	fmt.Scanln(&masukanNim)
	fmt.Print("Masukkan Nama Mahasiswa = ")
	fmt.Scanln(&masukanNama)
	fmt.Print("Masukkan Jurusan = ")
	fmt.Scanln(&masukanJurusan)

	_, err = db.Exec(`update m_mahasiswa set nim = ?, nama_mahasiswa = ?, jurusan = ? where id_mahasiswa = ?;`, &masukanNim, &masukanNama, &masukanJurusan, &masukanId)
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

//function show all data Mahasiswa
func ShowMahasiswa() {
	var option string
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
	fmt.Print("untuk melihat data mata kuliah yang diambil tekan y = ")
	fmt.Scanln(&option)
	if option == "y" {
		var tempnim string
		fmt.Print("Masukkan nim mahasiswa = ")
		fmt.Scanln(&tempnim)
		MatakuliahByNim(tempnim)
	}
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
