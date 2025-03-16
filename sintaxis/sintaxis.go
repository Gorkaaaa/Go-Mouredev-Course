package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Esta es la páginas de sintaxis.")

	// DEFINIR VARIABLE CON VALOR
	var mystring string = "Esto es una cadena de texto." // Variable String.
	fmt.Println(mystring)                                // Así hacemos print de una variable.

	// VALORES INT
	var myInt int = 7
	fmt.Println(myInt) // 7

	// SUMAR UN VALOR
	myInt = myInt + 4
	fmt.Println(myInt) // 11

	// SUMAR DENTRO DEL PRINT
	fmt.Println(myInt - 1) // 10

	// REPRESENTAR DOS TIPOS DE VALORES DISTINTOS.
	fmt.Println(mystring, myInt) // Esto es una cadena de texto. 11

	// Ver los datos con los que estamos trabajando.
	fmt.Println(reflect.TypeOf(myInt)) // int

	// Tipos de datos decimales
	var myFloat float64 = 6.5             // Definimos la variable.
	fmt.Println(myFloat)                  // 6.5
	fmt.Println(float64(myInt) + myFloat) //17,5 ::: Transformamos myInt a float y le sumamo el float original.

	// Tipos de datos boleanos.
	var myBool bool = true // Creamos la variable en falso
	myBool = false         // Le cambiamos el valor
	fmt.Println(myBool)    // Hacemos print de la variable.

	// Variable declarada de forma automatica e inicializada de manera abreviada.
	myString3 := "Esto es una cadena de texto" // Se nos define automáticamente.
	fmt.Println(myString3)                     // Esto es una cadena de texto

	// Constantes
	// Las podemos declarar sin utilizarlas.
	const myConst = "Esto es una constante"
	fmt.Println(myConst) // Esto es una constante.

	// Control de flujo
	// Función clasica con else if
	if myInt == 11 {
		fmt.Println("El valor es 11")
	} else if myInt == 10 {
		fmt.Println("El valor es 10")
	} else {
		fmt.Println("El valor no es 11")
	}

	// Agregas más condiciones.
	// AND
	var myString2 string = "Hola"
	if myInt == 11 && myString2 == "Hola" {
		fmt.Println("El valor es 11")
	} else if myInt == 10 {
		fmt.Println("El valor es 10")
	} else {
		fmt.Println("El valor no es 11")
	}

	// OR
	var myString4 string = "Hola"
	if myInt == 11 || myString4 == "Hola" {
		fmt.Println("El valor es 11")
	} else if myInt == 10 {
		fmt.Println("El valor es 10")
	} else {
		fmt.Println("El valor no es 11")
	}

	// ARRAY
	var myArray [3]int      // Definimos el nombre, numero de objetos y tipo de datos.
	myArray[0] = 1          // Le damos valor 1 a la posición 0.
	myArray[1] = 2          // Le damos valor 1 a la posición 1.
	myArray[2] = 3          // Le damos valor 1 a la posición 2.
	fmt.Println(myArray)    // [1 2 3]
	fmt.Println(myArray[1]) // 2

	// Map
	myMap := make(map[string]int)
	myMap["Brais"] = 36
	myMap["crais01"] = 35
	myMap["hoz98"] = 24
	fmt.Println(myMap["Brais"]) // 36
	fmt.Println(myMap)          // map[Brais:36 crais01:35 hoz98:24]

	// Map creando y inicializando
	myMap1 := map[string]int{"Brais": 36, "crais01": 35, "hoz98": 24}
	println(myMap1["Brais"]) // 36

	// Lista
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList.Back().Value)

	// Bucles
	// Iterar en los valores de un array.
	for index := 0; index < len(myArray); index++ {
		fmt.Println(myArray[index]) // 1 \n    2 \n   3 \n
	}

	// Llamar a una función
	myFunction()

	// Llamar a lo que retorna una función
	fmt.Println(Función_Return())

	// Estructura
	type MyStruct struct {
		name string
		age  int
	}

	myStruct := MyStruct{"Brais", 36}
	fmt.Println(myStruct)
}

// Función

func myFunction() {
	fmt.Println("Mi función")
}

func Función_Return() string { // Le indicamos el tipo de valor que queremos retornar.
	return "Mi función." // {Brais 36}
}
