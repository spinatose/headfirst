package main

import "fmt"

type Liters float64
func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Gallons float64
func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

func main () {
	carfuel := Gallons(1.2) // converts float to gallons
	busfuel := Liters(4.5) // converts float to liters
	carfuel += GallonsFromLiters(40.0) 
	busfuel += LitersFromGallons(30.0) 
 
	fmt.Printf("Car fuel: %s\n", carfuel)
	fmt.Printf("Bus fuel: %s\n", busfuel)

	fmt.Printf("Car fuel in liters: %s\n", carfuel.ToLiters())
	fmt.Printf("Bus fuel in gallons: %s\n", busfuel.ToGallons())
}

func GallonsFromLiters(l Liters) Gallons {
	return Gallons(l * .264)
}

func LitersFromGallons(g Gallons) Liters {
	return Liters(g * 3.785)
}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * .264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}