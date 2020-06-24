package views

import (
	"fmt"
	"os"

	"github.com/putra-kesuma/gomahasiswa/models"
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
		models.LaporanPerMatakuliah()
		break
	case "2":
		models.AvgLaporan()
		break
	case "3":
		os.Exit(0)
	default:
		MainMenu()
	}

}
