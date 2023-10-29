package types

import "fmt"

type integer = int // creating an alias for type 'int'

var coolNumber integer = 3

type distance float64 // creating your own type 'distance'
type distanceKm float64

func (miles distance) ToKm() distanceKm { // methods in Go; (miles distance) is a special argument
	return distanceKm(miles * 1.60934)
}

func Testing() {
	d := distance(4.5)
	inKm := d.ToKm()

	fmt.Printf("Distance is: %v\nDistance in Kilometers is: %v", d, inKm)
}
