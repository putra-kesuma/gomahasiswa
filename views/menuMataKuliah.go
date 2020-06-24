package views

import (
	"fmt"
	"os"
)

func MenuMataKuliah() {
	fmt.Println("**************************************")
	fmt.Println("Menu Mata Kuliah")
	fmt.Println("**************************************")
	fmt.Println("1. Tambah Mata Kuliah")
	fmt.Println("2. Hapus Mata Kuliah")
	fmt.Println("3. Ubah Mata Kuliah")
	fmt.Println("4. Lihat data mata kuliah")
	fmt.Println("5. Exit")
	fmt.Println("Pilihan menu dari 1-5")

	var selectedMenuMataKuliah string
	fmt.Scanln(&selectedMenuMataKuliah)
	switch selectedMenuMataKuliah {
	case "1":
		// models.InsertProduct()
		break
	case "2":
		// models.DeleteProduk()
		break
	case "3":
		// models.DetailProduk()
		break
	case "4":
		os.Exit(0)
	default:
		MainMenu()
	}

}
