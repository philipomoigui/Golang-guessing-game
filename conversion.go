package main

import (
	"fmt"
)

type feet float64
type centimeter float64
type minutes int
type seconds float64
type Fahrenheit float64
type celsius float64
type degrees float64
type radian float64
type pounds float64
type kilogram float64
type miles float64
type kilometer float64

type converter struct{}

const pi = 3.14159

func main() {
	cvr := converter{}
	// calling our method and printing the result at the same time
	fmt.Println(cvr.centimeterToFeet(5))

	// calling our method and printing the result at the same time
	fmt.Println(cvr.feetToCentimeter(45))

	// calling our method and printing the result at the same time
	fmt.Println(cvr.minutesToSeconds(4))

	// calling our method and printing the result at the same time
	fmt.Println(cvr.secondsToMinutes(300))

	//calling our method and printing the result at the same time
	fmt.Println(cvr.celsiusToFahrenheit(54))

	// calling our method and printing the result at the same time
	fmt.Println(cvr.fahrenheitToCelsius(85))

	// calling our method and printing the result at the same time
	fmt.Println(cvr.degreesToRadian(90))

	// calling our method and printing the result at the same time
	fmt.Println(cvr.radianToDegrees(4))

	// calling our method and printing the result at the same time
	fmt.Println(cvr.kilogramToPounds(5))

	// calling our method and printing the result at the same time
	fmt.Println(cvr.poundsToKilogram(3))

	// calling our method and printing the result at the same time
	fmt.Println(cvr.milesToKilometer(10))

	// calling our method and printing the result at the same time
	fmt.Println(cvr.KilometerToMiles(85))

}

func (cvr converter) centimeterToFeet(c centimeter) feet {
	//conversion code
	return feet(c * 30.48)

}

func (cvr converter) feetToCentimeter(f feet) centimeter {
	//conversion code
	return centimeter(f / 30.48)
}

func (cvr converter) minutesToSeconds(m minutes) seconds {
	// Conversion code
	return seconds(m * 60)
}
func (cvr converter) secondsToMinutes(s seconds) minutes {
	// Conversion code
	return minutes(s / 60)
}
func (cvr converter) celsiusToFahrenheit(c celsius) Fahrenheit {
	// Conversion code
	return Fahrenheit(c*1.8 + 32)
}
func (cvr converter) fahrenheitToCelsius(f Fahrenheit) celsius {
	// Conversion code
	return celsius((f - 32) / 1.8)
}
func (cvr converter) degreesToRadian(d degrees) radian {
	// Conversion code
	return radian(d * pi / 180)
}

func (cvr converter) radianToDegrees(r radian) degrees {
	// Conversion code
	return degrees(r * 180 / pi)
}

func (cvr converter) kilogramToPounds(k kilogram) pounds {
	// Conversion code
	return pounds(k / 0.45359237)
}

func (cvr converter) poundsToKilogram(p pounds) kilogram {
	// Conversion code
	return kilogram(p * 0.45359237)
}

func (cvr converter) milesToKilometer(m miles) kilometer {
	// Conversion code
	return kilometer(m * 1.609344)
}

func (cvr converter) KilometerToMiles(k kilometer) miles {
	// Conversion code
	return miles(k / 1.609344)
}
