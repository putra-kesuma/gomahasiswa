package views

import (
	"fmt"
	"os"

	"github.com/putra-kesuma/gomahasiswa/models"
)

func MenuMahasiswa() {
	fmt.Println("**************************************")
	fmt.Println("Menu Mahasiswa")
	fmt.Println("**************************************")
	fmt.Println("1. Tambah Mahasiswa")
	fmt.Println("2. Hapus Mahasiswa")
	fmt.Println("3. Ubah Mahasiswa")
	fmt.Println("4. Lihat data mahasiswa dan mata kuliah yang di ambil")
	fmt.Println("5. Exit")
	fmt.Println("Pilihan menu dari 1-5")

	var selectedMenuMahasiswa string
	fmt.Scanln(&selectedMenuMahasiswa)
	switch selectedMenuMahasiswa {
	case "1":
		models.InsertMahasiswa()
		break
	case "2":
		models.DeleteMahasiswa()
		break
	case "3":
		models.UpdateMahasiswa()
		break
	case "4":
		models.ShowMahasiswa()
		break
	case "5":
		os.Exit(0)
	default:
		MainMenu()
	}

}
