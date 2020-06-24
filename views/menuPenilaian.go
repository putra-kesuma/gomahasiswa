package views

import (
	"fmt"
	"os"

	"github.com/putra-kesuma/gomahasiswa/models"
)

func MenuNilai() {
	fmt.Println("**************************************")
	fmt.Println("Menu Penilaian")
	fmt.Println("**************************************")
	fmt.Println("1. Tambah Nilai")
	fmt.Println("2. Hapus Nilai")
	fmt.Println("3. Lihat data Penilaian")
	fmt.Println("4. Exit")
	fmt.Println("Pilihan menu dari 1-4")

	var selectedMenuNilai string
	fmt.Scanln(&selectedMenuNilai)
	switch selectedMenuNilai {
	case "1":
		models.InsertNilai()
		break
	case "2":
		models.DeleteNilai()
		break
	case "3":
		models.DetailNilai()
		break
	case "4":
		os.Exit(0)
	default:
		MainMenu()
	}

}
