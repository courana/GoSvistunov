package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func length(l string) {
	k := 3.28
	value, err := strconv.ParseFloat(l, 64)
	if err != nil {
		fmt.Println("Wrong input 1st argument")
		return
	}
	meter := value * k
	feet := value / k
	fmt.Printf("%s feet = %f meters\n", l, meter)
	fmt.Printf("%s meters = %f feet\n", l, feet)
}

func weight(w string) {
	k := 0.45
	value, err := strconv.ParseFloat(w, 64)
	if err != nil {
		fmt.Println("Wrong input 2nd argument")
		return
	}
	pound := value * k
	kg := value / k
	fmt.Printf("%s pounds = %f kgs\n", w, kg)
	fmt.Printf("%s kgs = %f pounds\n", w, pound)
}

func temp(t string) {
	value, err := strconv.ParseFloat(t, 64)
	if err != nil {
		fmt.Println("Wrong argument 1st argument")
		return
	}
	fmt.Printf("%s deg Fahrenheit = %f deg Celsius\n", t, (value-32)*5/9)
	fmt.Printf("%s deg Celsius = %f deg Fahrenheit\n", t, (value+32)/5*9)

	fmt.Printf("%s deg Celsius = %f deg Kelvin\n", t, value-273.15)
	fmt.Printf("%s deg Kelvin = %f deg Celsius\n", t, value+273.15)

	fmt.Printf("%s deg Fahrenheit = %f deg Kelvin\n", t, (value-32)*5/9-273.15)
	fmt.Printf("%s deg Kelvin = %f deg Fahrenheit\n", t, (value+32)/5*9+273.15)
}

func main() {
	var args [2]string
	if len(os.Args) > 1 {
		args[0] = os.Args[1]
		args[1] = os.Args[2]
	} else {
		fmt.Print("Enter a value and a measurement ('l' for length, 'w' for weight, 't' for temperature).\nValue: ")
		fmt.Scan(&args[0])
		fmt.Print("Measurement: ")
		fmt.Scan(&args[1])
	}

	switch strings.ToLower(args[1]) {
	case "l":
		length(args[0])
	case "w":
		weight(args[0])
	case "t":
		temp(args[0])
	default:
		fmt.Println("Wrong input 2nd argument")
	}
}
