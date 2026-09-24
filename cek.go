package main

import "fmt"

func main() {
	var number int
	var pilihan string

	// Looping
		for {
		fmt.Print("Masukkan angka : ")
		fmt.Scan(&number)

	// Konfirmasi mau ulang atau tidak
		fmt.Print("Cek lagi? (y/t) : ")
		fmt.Scan(&pilihan)

	// kalau user ketik 't', Looping akan berhenti
		if pilihan == "t" {
			break
		}
	}
}
