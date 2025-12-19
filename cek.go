package main

import "fmt"

func main() {
	var number int
	var pilihan string

	// Looping
		for {
		fmt.Print("Masukkan angka : ")
		fmt.Scan(&number)

	// Branching untuk cek genap atau ganjil
		if number%2 == 0 {
		fmt.Println("itu bilangan Genap")
		} else {
		fmt.Println("itu bilangan Ganjil")
		}

	// Konfirmasi mau ulang atau tidak
		fmt.Print("Cek lagi? (y/t) : ")
		fmt.Scan(&pilihan)

	// kalau user ketik 't', Looping akan berhenti
		if pilihan == "t" {
			break
		}
	}
}
