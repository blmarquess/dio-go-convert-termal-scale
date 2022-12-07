package main

import (
	"fmt"
)

func main() {
	var selectedOption int
	var degree float64
	fmt.Println(" 1. Convert kelvin to Celsius\n 2. Convert kelvin to Fahrenheit\n 3. Convert Celsius to Fahrenheit")
	fmt.Println(" 4. Convert Celsius to Kelvin\n 5. Convert Fahrenheit to Celsius\n 6. Convert Fahrenheit to Kelvin")
	fmt.Print("Select an option: ")
	fmt.Scan(&selectedOption)
	fmt.Print("Enter the degree: ")
	fmt.Scan(&degree)
	switch selectedOption {
	case 1:
		fmt.Printf("%.2f Kelvin = %.2f Celsius\n", degree, convertKelvinToCelsius(degree))
	case 2:
		fmt.Printf("%.2f Kelvin = %.2f Fahrenheit\n", degree, convertKelvinToFahrenheit(degree))
	case 3:
		fmt.Printf("%.2f Celsius = %.2f Fahrenheit\n", degree, convertCelsiusToFahrenheit(degree))
	case 4:
		fmt.Printf("%.2f Celsius = %.2f Kelvin\n", degree, convertCelsiusToKelvin(degree))
	case 5:
		fmt.Printf("%.2f Fahrenheit = %.2f Celsius\n", degree, convertFahrenheitToCelsius(degree))
	case 6:
		fmt.Printf("%.2f Fahrenheit = %.2f Kelvin\n", degree, convertFahrenheitToKelvin(degree))
	default:
		fmt.Println("Invalid option")
	}
}

func convertKelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273
}

func convertKelvinToFahrenheit(kelvin float64) float64 {
	return (kelvin-273)*9/5 + 32
}

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func convertCelsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func convertFahrenheitToKelvin(fahrenheit float64) float64 {
	return (fahrenheit-32)*5/9 + 273
}
