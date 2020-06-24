package main

import (
	"fmt"
	"os"

	"github.com/putra-kesuma/gomahasiswa/utils"
	"github.com/putra-kesuma/gomahasiswa/views"
)

func main() {
	views.MainMenu()
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			views.MenuMahasiswa()
			break
		case "2":
			views.MenuMataKuliah()
			break
		case "3":
			views.MenuDosen()
			break
		case "4":
			views.MenuNilai()
			break
		case "5":
			views.MenuLaporan()
			break
		case "6":
			os.Exit(0)
		default:
			utils.ClearScreen()
		}
	}
}
