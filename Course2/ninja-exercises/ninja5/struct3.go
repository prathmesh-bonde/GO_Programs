package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	truck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourwheel: true,
	}

	sedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(truck)
	fmt.Println(sedan)
}
