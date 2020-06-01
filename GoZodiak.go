// Se debe digitar día y mes en números.

package main

import "fmt"

var rta string
var dia, mes int

func main() {

	fmt.Println("*****************************")
	fmt.Println("****  ¡Conoce tu signo!  ****")
	fmt.Println("*****************************")
	fmt.Println("")
	fmt.Println("Digita el día de nacimiento.")
	fmt.Scanf("%d\n", &dia)
	fmt.Println("Digita el mes de nacimiento (numérico).")
	fmt.Scanf("%d\n", &mes)
	fmt.Println("")

	if dia >= 20 && mes == 1 {
		rta = "* Acuario *"
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia <= 18 && mes == 2 {
		rta = "* Acuario * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia >= 19 && mes == 2 {
		rta = "* Piscis * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia <= 20 && mes == 3 {
		rta = "* Piscis * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia >= 21 && mes == 3 {
		rta = "* Aries * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia <= 19 && mes == 4 {
		rta = "* Aries * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia >= 20 && mes == 4 {
		rta = "* Tauro * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia <= 20 && mes == 5 {
		rta = "* Tauro * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia >= 21 && mes == 5 {
		rta = "* Géminis * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia <= 20 && mes == 6 {
		rta = "* Géminis * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia >= 21 && mes == 6 {
		rta = "* Cáncer * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia <= 22 && mes == 7 {
		rta = "* Cáncer * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia >= 23 && mes == 7 {
		rta = "* Leo * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia <= 22 && mes == 8 {
		rta = "* Leo * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia >= 23 && mes == 8 {
		rta = "* Virgo * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia <= 22 && mes == 9 {
		rta = "* Virgo * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia >= 23 && mes == 9 {
		rta = "* Libra * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia <= 22 && mes == 10 {
		rta = "* Libra * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia >= 23 && mes == 10 {
		rta = "* Escorpio * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia <= 21 && mes == 11 {
		rta = "* Escorpio * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia >= 22 && mes == 11 {
		rta = "* Sagitario * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia <= 21 && mes == 12 {
		rta = "* Sagitario * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia >= 22 && mes == 12 {
		rta = "* Capricornio * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else if dia >= 19 && mes == 1 {
		rta = "* Capricornio * "
		fmt.Println("Tu signo sodiacal es: \n", rta)
	} else {
		fmt.Println("Algunos datos son incorrectos.")
	}
}
