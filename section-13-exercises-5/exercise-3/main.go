package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	aTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Red",
		},
		fourWheel: true,
	}
	fmt.Println(aTruck)
	fmt.Println(aTruck.vehicle.color)
	aSedan := sedan{
		luxury: false,
		vehicle: vehicle{
			doors: 5,
			color: "blue",
		},
	}
	fmt.Println(aSedan)
	fmt.Println(aSedan.vehicle.color)
}
