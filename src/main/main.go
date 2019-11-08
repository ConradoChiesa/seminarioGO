package main

import (
	"fmt"
	"main/agencia"
)

func main() {
	auto1 := agencia.NewAuto("123 ABC", "Honda", 2007)
	auto2 := agencia.NewAuto("321 BAC", "BWM", 2008)
	auto3 := agencia.NewAuto("765 FDG", "Lifan", 2018)
	auto1.Acelerar()
	auto1.Acelerar()
	auto1.Acelerar()
	agen := agencia.NewAgencia()
	agen.AddAuto(auto1)
	agen.AddAuto(auto2)
	agen.AddAuto(auto3)

	fmt.Println("Auto", auto1)
	auto1.Frenar()
	fmt.Println("Auto", auto1)

	fmt.Println(agen.GetAutos())
	agen.
}
