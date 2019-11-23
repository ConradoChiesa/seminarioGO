package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"main/agencia"
	"net/http"
)

var db *gorm.DB

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
	agen.DeleteAuto(1)
	fmt.Println(agen.GetAutos())

	fmt.Printf("PUNTEROS\n")
	a := 5
	b := 10

	var c *int
	c = &b
	fmt.Println("a", a)
	fmt.Println("&a", &a)
	fmt.Println("&b", &b)
	fmt.Println("c", c)
	fmt.Println("&c", &c)

	fmt.Printf("Otro ejemplo\n")
	myString := "Hola Goland"
	fmt.Println(myString)

	// resultado, _ :=

	var d []byte
	d, _ = json.Marshal(myString)
	fmt.Println(a)

	var otroString string
	_ = json.Unmarshal(d, &otroString)
	fmt.Println("otroString", otroString)

	miAuto := agencia.Auto{
		Patente: "GOS235",
		Marca:   "Honda",
		Modelo:  2007,
	}

	//b, _ =json.Marshal(miAuto)
	var miAutoMarshelled []byte
	miAutoMarshelled, _ = json.Marshal(miAuto)

	fmt.Println(miAutoMarshelled)

	router := gin.Default()

	router.GET("/hello", GetHello)
	router.GET("/cars", GetCars)
	router.POST("/cars", AddCar)
	if err := router.Run(); err != nil {
		fmt.Println(err)
	}

	db, err := gorm.Open("mysql", "root:root@/db_agencia")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&agencia.Auto{})

}

func AddCar(ctx *gin.Context) {
	var car agencia.Auto
	ctx.BindJSON(&car)
	db.Create(car)
	ctx.JSON(http.StatusOK, gin.H{
		"auto": car,
	})
}

func GetHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bienvenido a Microservicios GoLand",
	})
}

func GetCars(ctx *gin.Context) {
	var cars []agencia.Auto
	db.Find(&cars)
	fmt.Println(cars)
	/*	miAuto := agencia.Auto{
			Patente: "GOS235",
			Marca: "Honda",
			Modelo: 2007,

		}
	*/
	ctx.JSON(http.StatusOK, gin.H{
		"resultado": cars,
	})
}
