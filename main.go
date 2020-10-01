package main

import "fmt"

func main() {

	var dia int
	var mes int
	fmt.Scan(&dia)
	fmt.Scan(&mes)
	if (mes == 1 && dia >= 21 && dia <= 31) || (mes == 2 && dia >= 1 && dia <= 19) {
		fmt.Print("acuario")
	} else if (mes == 2 && dia >= 20 && dia <= 28) || (mes == 3 && dia >= 1 && dia <= 20) {
		fmt.Print("piscis")
	} else if (mes == 3 && dia >= 21 && dia <= 31) || (mes == 4 && dia >= 1 && dia <= 20) {
		fmt.Print("aries")
	} else if (mes == 4 && dia >= 21 && dia <= 30) || (mes == 5 && dia >= 1 && dia <= 21) {
		fmt.Print("tauro")
	} else if (mes == 5 && dia >= 22 && dia <= 31) || (mes == 6 && dia >= 1 && dia <= 21) {
		fmt.Print("geminis")
	} else if (mes == 6 && dia >= 22 && dia <= 30) || (mes == 7 && dia >= 1 && dia <= 23) {
		fmt.Print("cancer")
	} else if (mes == 7 && dia >= 24 && dia <= 31) || (mes == 8 && dia >= 1 && dia <= 23) {
		fmt.Print("leo")
	} else if (mes == 8 && dia >= 24 && dia <= 31) || (mes == 9 && dia >= 1 && dia <= 23) {
		fmt.Print("virgo")
	} else if (mes == 9 && dia >= 24 && dia <= 30) || (mes == 10 && dia >= 1 && dia <= 23) {
		fmt.Print("libra")
	} else if (mes == 10 && dia >= 24 && dia <= 31) || (mes == 11 && dia >= 1 && dia <= 22) {
		fmt.Print("escorpio")
	} else if (mes == 11 && dia >= 23 && dia <= 31) || (mes == 12 && dia >= 1 && dia <= 21) {
		fmt.Print("sagitario")
	} else if (mes == 12 && dia >= 22 && dia <= 31) || (mes == 1 && dia >= 1 && dia <= 20) {
		fmt.Print("capricornio")
	}
}
