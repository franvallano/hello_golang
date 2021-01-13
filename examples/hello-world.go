package main

import (
	"fmt" 
	"time"
	"os"
)
import "io/ioutil"

// Constant example
const year int = 2020

// Model example
type Person struct {
	name string
	age int
	country string
}

func main() {
	fmt.Println()
	fmt.Println("Hello World")

	time.Sleep(time.Second * 1)
	
	fmt.Println()
	fmt.Println("// ***************************************** //")
	fmt.Println()

	var name string = "Fran"

	country := "Spain"

	var decimal float32 = 1.45

	fmt.Print("Decimal: ")
	fmt.Println(decimal)

	fmt.Print("Year (Constant): ")
	fmt.Println(year)

	fmt.Println("Hello " + name + ", from " + country)

	fmt.Println()
	fmt.Println("// ***************************************** //")
	fmt.Println()

	var person1 = Person {name : "Fran", age : 29,	country : "Spain"}

	var person2 = Person {"Pepe", 30, "France"}

	fmt.Print("Person 1: ")
	fmt.Println(person1)

	fmt.Print("Person 1 - Name: ")
	fmt.Println(person1.name)

	fmt.Print("Person 2: ")
	fmt.Println(person2)

	fmt.Print("Person 2 - Age: ")
	fmt.Println(person2.age)

	fmt.Println()
	fmt.Println("// ***************************************** //")
	fmt.Println()

	fmt.Println(operation(1, 2, "+"))
	fmt.Println(operation(4, 20, "*"))

	fmt.Println()
	fmt.Println("// ***************************************** //")
	fmt.Println()

	fmt.Print("Closure example: ")
	fmt.Println(closure("Fran", "Delgado"))

	fmt.Println()
	fmt.Println("// ***************************************** //")
	fmt.Println()

	fmt.Print("Closure Params example: ")
	fmt.Println(closureParams("Fran", "Delgado", "Vallano"))

	fmt.Println()
	fmt.Println("// ***************************************** //")
	fmt.Println()

	// Arrays
	cars := [3]string {"Mercedes", "Ford", "Fiat"}
	
	fmt.Print("Array Cars: ")
	fmt.Println(cars)
	fmt.Println()

	for _,car := range cars {
		fmt.Println(car)
	} 

	fmt.Println()

	// Dynamic Array (Slice)
	cities := []string {"Sevilla", "Huelva", "Cádiz"}

	fmt.Print("Array Cities: ")
	fmt.Println(cities)

	cities = append(cities, "Córdoba")
	cities = append(cities, "Málaga")

	fmt.Print("Array Cities: ")
	fmt.Println(cities)

	// Dimensional Arrays 
	var supers [3][2]string
	supers[0][0] = "Batman"
	supers[0][1] = "DC"
	supers[1][0] = "Superman"
	supers[1][1] = "DC"
	supers[2][0] = "Spiderman"
	supers[2][1] = "Marvel"
	
	fmt.Println()
	fmt.Print("Array Supers: ")
	fmt.Println(supers)
	fmt.Print("Array Supers length: ")
	fmt.Println(len(supers))
	fmt.Print("Array Supers filter: ")
	fmt.Println(supers[0:1])

	fmt.Println()
	fmt.Println("// ***************************************** //")
	fmt.Println()


	fmt.Print("Welcome ")
	if(len(os.Args)>1){
		fmt.Println(os.Args[1]) // os.Args[0] exec route
	}else{
		fmt.Println("User")
	}

	fmt.Println()
	fmt.Println("// ***************************************** //")
	fmt.Println()
	
	fmt.Println("Read file: ")
	readFile("file.txt")

	fmt.Println("// ***************************************** //")
	fmt.Println()
	
	fmt.Println("Write file: ")
	writeFile(string(os.Args[2]) + "\n", "file.txt")
	readFile("file.txt")

	fmt.Println("// ***************************************** //")
	fmt.Println()
}


// Function example
func operation (num1 float32, num2 float32, op string) (string, float32) {
	var result float32
	
	switch op {
		case "+":
			result = num1 + num2
		case "-":
			result =  num1 - num2
		case "*":
			result =  num1 * num2	
		case "/":
			result =  num1 / num2			
	}
	return "Result :", result
}

// Closure example

func closure (name string, surname string) string{

	completeName := func() string {
		return name + " " + surname
	}

	return completeName()

}

// Closure param

func closureParams (params ...string) string{

	completeName := func() string {
		var completeChar string
		// _ dont pass vars
		for _, param := range params {
			completeChar = completeChar + param + " "
		}
		return completeChar
	}

	return completeName()

}


// Read & Write 

func readFile(fileName string){
	texto, err := ioutil.ReadFile(fileName)
	showError(err)
	fmt.Println(string(texto))
}

func writeFile(newText string, fileName string){
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0777)
	showError(err)

	text, err := file.WriteString(newText)
	showError(err)
	fmt.Println(text)
}


func showError(e error){
	if e != nil {  // nil = without error
		panic(e)
	}
}