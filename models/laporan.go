package models

import (
	"fmt"
	"log"

	"github.com/putra-kesuma/gomahasiswa/services"
)

//struct for Laporan
type laporan struct {
	nim           string
	namaMahasiswa string
	jurusan       string
	namaMk        string
	sks           int
	nilai         int
}

type avglaporan struct {
	namaMk string
	sks    int
	nilai  float64
}

func LaporanPerMatakuliah() {
	db, err := services.Connect() //call function connect
	defer db.Close()              // kalau sudah open jangan lupa close koneksinya
	data, err := db.Query("select * from view_penilaian")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	var result []laporan
	for data.Next() {
		var allLaporan = laporan{}
		var err = data.Scan(&allLaporan.nim, &allLaporan.namaMahasiswa, &allLaporan.jurusan, &allLaporan.namaMk, &allLaporan.sks, &allLaporan.nilai)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, allLaporan)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Laporan Mahasiswa Per Matakuliah")
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("NIM     | Nama Mahasiswa | Jurusan | Matakuliah | sks | nilai | ket |")
	fmt.Println("---------------------------------------------------------------------")
	for _, v := range result {
		fmt.Printf("%v", v.nim)
		fmt.Printf("%10v", v.namaMahasiswa)
		fmt.Printf("%15v", v.jurusan)
		fmt.Printf("%15v", v.namaMk)
		fmt.Printf("%8v", v.sks)
		fmt.Printf("%8v", v.nilai)
		//bisa pakai case di db, tapi saya memilih untuk mengolah nya di backend
		//dengan membuat sebuah func grade
		ket := grade(v.nilai)
		fmt.Printf("%4v\n", ket)
	}
}

func AvgLaporan() {
	db, err := services.Connect() //call function connect
	defer db.Close()              // kalau sudah open jangan lupa close koneksinya
	data, err := db.Query(`select mk.nama_matakuliah, mk.sks, avg(n.nilai) as rata_nilai from 
	m_matakuliah mk join m_nilai n on mk.id_matakuliah = n.id_matakuliah group by mk.nama_matakuliah;`)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	var result []avglaporan
	for data.Next() {
		var allAvglaporan = avglaporan{}
		var err = data.Scan(&allAvglaporan.namaMk, &allAvglaporan.sks, &allAvglaporan.nilai)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, allAvglaporan)
	}

	if err = data.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Laporan rata rata nilai Per Matakuliah")
	fmt.Println("-----------------------------")
	fmt.Println(" Matakuliah | sks | nilai |")
	fmt.Println("-----------------------------")
	for _, v := range result {
		fmt.Printf("%8v", v.namaMk)
		fmt.Printf("%8v", v.sks)
		fmt.Printf("%8v\n", v.nilai)
	}
}

func grade(catchNilai int) string {
	var desc string
	if catchNilai > 90 {
		desc = "A"
	} else if catchNilai > 80 {
		desc = "B"
	} else if catchNilai > 70 {
		desc = "C"
	} else if catchNilai > 60 {
		desc = "D"
	} else {
		desc = "E"
	}
	return desc
}
