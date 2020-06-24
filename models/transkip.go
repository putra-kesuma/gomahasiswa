package models

import (
	"fmt"
	"log"

	"github.com/putra-kesuma/gomahasiswa/services"
)

//struct for transkip
type transkip struct {
	nimMahasiswa   string
	idMatakuliah   int
	namaMatakuliah string
	sks            int
}

//function insert Transkip
func InsertTranskip(catchNim string) {
	db, err := services.Connect() //call function connect
	defer db.Close()
	var masukanIdMK int
	DetailMatakuliah()
	fmt.Print("Masukkan Id Matakuliah = ")
	fmt.Scanln(&masukanIdMK)

	_, err = db.Exec(`insert into m_transkip (nim_mahasiswa,id_matakuliah) 
	value (?,?)`, &catchNim, &masukanIdMK)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mata kuliah berhasil ditambahkan")
}

//function menampilkan data mk berdasrkan nim
func MatakuliahByNim(catchNim string) {
	db, err := services.Connect() //call function connect
	defer db.Close()              // kalau sudah open jangan lupa close koneksinya
	data, err := db.Query(`select mhs.nim, mk.id_matakuliah, mk.nama_matakuliah, mk.sks  
	from m_mahasiswa mhs join m_transkip t on mhs.nim = t.nim_mahasiswa
	join m_matakuliah mk on t.id_matakuliah = mk.id_matakuliah where mhs.nim = ?`, &catchNim)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	var result []transkip
	for data.Next() {
		var allTranskip = transkip{}
		var err = data.Scan(&allTranskip.nimMahasiswa, &allTranskip.idMatakuliah, &allTranskip.namaMatakuliah, &allTranskip.sks)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, allTranskip)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}
