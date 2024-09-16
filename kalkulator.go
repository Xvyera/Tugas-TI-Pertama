package main

import "fmt"

func main() {
	var angka1, angka2 float32
	var operasi string

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&angka1)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&angka2)

	fmt.Print("Masukkan operasinya ( + - * / ): ")
	fmt.Scanln(&operasi)

	switch operasi {

	case "+":
		fmt.Printf("%f %s %f = %f", angka1, operasi, angka2, angka1+angka2)

	case "-":
		fmt.Printf("%f %s %f = %f", angka1, operasi, angka2, angka1-angka2)

	case "*":
		fmt.Printf("%f %s %f = %f", angka1, operasi, angka2, angka1*angka2)

	case "/":
		if angka2 == 0.0 {
			fmt.Println("Situasi saat dibagi oleh nol")
		} else {
			fmt.Printf("%f %s %f = %f", angka1, operasi, angka2, angka1/angka2)
		}
	default:
		fmt.Println("Operasi Tidak Valid")

	}

}
