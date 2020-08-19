/*
-Calculadora simple 
-Contiene suma,resta,multiplicación y división
-Uso fácil en terminal
-Esta calculadora solo contempla dos valores en cualquier operación
*/

package main

import (
   "fmt"
   "os"
   "os/exec"
   "time"
   "strconv"
   "github.com/fatih/color"

)

func main() {

Limpiar()
Banner()
Tiempo()
Calculadora()

}

func Limpiar() {

c := exec.Command("clear")
c.Stdout = os.Stdout
c.Run()

}

func Banner() {

	b := color.New(color.FgBlue).Add(color.Bold)
b.Println("")
b.Println(" .--.        .-.              .-.            .-.                  ")
b.Println(": .--'       : :              : :            : :                  ")
b.Println(": :    .--.  : :   .--. .-..-.: :   .--.   .-' : .--. .--.  .--.  ")
b.Println(": :__ ' .; ; : :_ '  ..': :; :: :_ ' .; ; ' .; :' .; :: ..'' .; ; ")
b.Println("`.__.'`.__,_;`.__;`.__.'`.__.'`.__;`.__,_;`.__.'`.__.':_;  `.__,_;")
b.Println("")

}

func Tiempo() {

a := color.New(color.FgRed).Add(color.Underline)
	ahora := time.Now()
	a.Println("Horario:", ahora)
	año := ahora.Year()
	mes := ahora.Month()
	dia := ahora.Day()
	hora := ahora.Hour()
	minutos := ahora.Minute()
	segundos := ahora.Second()

a.Println("Año:", año)
a.Println("Mes:", mes)
a.Println("Dia:", dia)
a.Println("Hora:", hora)
a.Println("Minutos:", minutos)
a.Println("Segundos:", segundos)
fmt.Println("")
a.Println("∆∆HORARIO FORMATO INTERNACIONAL√√")
fmt.Println("")

}

func Calculadora() {

	c := color.New(color.FgCyan).Add(color.Underline)
 calc := leida("Digite el número de la operación con la que quieras trabajar:\n\n [1]Suma\n\n [2]Resta\n\n [3]Multiplicación\n\n [4]División\n\n \t[<>]Elige una opción: ")
fmt.Println(calc)

if calc == "1" {
  c.Println("\tHAS ELEGIDO SUMA!!\n")
  num1, num2 := numeros()
  result := num1 + num2
  c.Println(fmt.Sprintf("\nEl resultado de la suma es: " +  "%d\n", result))
}else if calc == "2" {
  c.Println("\tHAS ELEGIDO RESTA!\n")
  num1, num2 := numeros()
  result := num1 - num2
  c.Println(fmt.Sprintf("\nEl resultado de la resta es: " + "%d\n", result))
}else if calc == "3" {
  c.Println("\tHAS ELEGIDO MULTIPLICACIÓN!\n")
  num1, num2 := numeros()
  result := num1 * num2
  c.Println(fmt.Sprintf("\nEl resultado de la multiplicación es: " + "%d\n", result))
}else if calc == "4" {
  c.Println("\tHAS ELEGIDO DIVISIÓN!\n")
  num1, num2 := numeros()
  result := float32(num1) / float32(num2)
  c.Println("\nEl resultado de la división es: " + "%d\n", result)
}else {
c.Println("\tEsta opción no existe\n")
}

}


func leida(message string) string {
	d := color.New(color.FgYellow).Add(color.Bold)
d.Print(message)
var input string
fmt.Scanln(&input)
return input

}

func numeros() (int, int) {

  num1String := leida("[+]Ingresa el primer número: ")
  num1, _ := strconv.Atoi(num1String)
  num2String := leida("[+]Ingresa el segundo número: ")
  num2, _ := strconv.Atoi(num2String)
  return num1, num2
}


