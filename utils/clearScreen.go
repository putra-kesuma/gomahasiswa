package utils

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/putra-kesuma/gomahasiswa/views"
)

func ClearScreen() { // function untuk membersihkan layar
	osRunning := runtime.GOOS
	if osRunning == "linux" || osRunning == "darwin" {
		cmd := exec.Command("clear") //inisialisasi perintah yang akan di jalankan di command
		cmd.Stdout = os.Stdout
		cmd.Run() // jalankan perintah
	} else if osRunning == "windows" { //jika os nya windows makan menjalankan perintah cls
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	views.MainMenu()
}
