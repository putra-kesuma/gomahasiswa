package views

import (
	"fmt"
	"os"
)

func MenuDosen() {
	fmt.Println("**************************************")
	fmt.Println("Menu Dosen")
	fmt.Println("**************************************")
	fmt.Println("1. Tambah Dosen")
	fmt.Println("2. Hapus Dosen")
	fmt.Println("3. Ubah Dosen")
	fmt.Println("4. Lihat data dosen dan mata kuliah yang di ampu")
	fmt.Println("5. Exit")
	fmt.Println("Pilihan menu dari 1-5")

	var selectedMenuDosen string
	fmt.Scanln(&selectedMenuDosen)
	switch selectedMenuDosen {
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
