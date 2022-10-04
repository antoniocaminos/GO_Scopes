package main

import (
	"myapp/packageone"
)

var myVar = "packagelevel var"

func main() {
	//Variables
	//declarar una variable a nivel package en main
	//llamado myVar

	//declarar una variable de tipo block en la funcion main
	//llamada blockVar

	var blockVar = "blocklevel var"
	//declarar una variable a nivel package en packageone
	//llamada PackageVar

	//crear y exportar una funcion en el packageone
	//llamada PrintMe

	//en la funcion main imrpimir los valores de myVar,
	//blockVar, y PackageVar en una linea usando la funcion PrintMe en packageOne

	packageone.PrintMe(myVar, blockVar)

}
