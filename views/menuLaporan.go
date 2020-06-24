package views

import (
	"fmt"
	"os"
)

func MenuLaporan() {
	fmt.Println("**************************************")
	fmt.Println("Menu Laporan")
	fmt.Println("**************************************")
	fmt.Println("1. Laporan mahasiswa per mata kuliah")
	fmt.Println("2. Laporan rata rata nilai per mata kuliah")
	fmt.Println("3. Exit")
	fmt.Println("Pilihan menu dari 1-3")

	var selectedMenuLaporan string
	fmt.Scanln(&selectedMenuLaporan)
	switch selectedMenuLaporan {
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
