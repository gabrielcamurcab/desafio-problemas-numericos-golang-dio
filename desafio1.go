package main

import "fmt"

func desafio1() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func desafio2() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PinPan")
		} else if i%3 == 0 {
			fmt.Println("Pin")
		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("Desafio 1")
	desafio1()
	fmt.Print("\n")
	fmt.Println("Desafio 2")
	desafio2()
}
