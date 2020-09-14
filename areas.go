package main

import (
	"fmt"
	"math"
)

var opc int

func main() {
	var opc int
	menu := ("menu \n [0] salir \n [1] área cuadrado \n [2] área triangulo \n [3] área circulo \n [4] Farenheit a Celsius\n opcion:")
	fmt.Println(menu)
	fmt.Scan(&opc)
	switch opc {
	case 1:
		a_cuadro()
		break
	case 2:
		a_triangulo()
		break
	case 3:
		a_circulo()
		break
	case 4:
		fahrenheit_Celsius()
		break
	default:
		println("Opción no valida")
		break
	case 0:
		println("salir")
	}
}

func a_cuadro() {
	var lado float32
	fmt.Println("Área de un cuadro")
	fmt.Println("Ingrese el tamaño de uno de los lados del cuadrado en cm:")
	fmt.Scan(&lado)
	fmt.Println("El área es de ", lado*lado, "cm2")
}

func a_circulo() {
	var radio float64
	pi := math.Pi
	fmt.Println("Área de un circulo")
	fmt.Println("Ingrese el tamaño del radio en cm:")
	fmt.Scan(&radio)
	resultado := radio * radio
	resultado = resultado * pi
	fmt.Println("El área es de ", resultado)
}

func a_triangulo() {
	var base float64
	var altura float64
	fmt.Println("Área de un triangulo")
	fmt.Println("Ingrese la altura del triangulo")
	fmt.Scan(&altura)
	fmt.Println("Ingrese el tamaño de la base")
	fmt.Scan(&base)
	resulado := ((base * altura) / 2)
	fmt.Printf("El área es de: %f cm^2", resulado)
}

func fahrenheit_Celsius() {
	var gradosF float64
	fmt.Println("De fahrenheit a Celsius")
	fmt.Println("Ingrese la temperatura en grados farenheit")
	fmt.Scan(&gradosF)
	resulado := ((gradosF - 32) * 5) / 9
	fmt.Printf("La temperatura es de: %f C°", resulado)
}
