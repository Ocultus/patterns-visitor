package main

import (
	"Ocultus/patterns-visitor.git/figures"
	"fmt"
)

func main() {
 	circle:=&figures.Circle{
	 	Dot: *&figures.Dot{
			 X: 10,
			 Y: 0,
		},
		Radius:10,
	}

	square:=&figures.Square{
		Dot: *&figures.Dot{
			X: 0,
			Y: 0,
		},
		Side: 20,
	}

	rec:=&figures.Rectangle{
		Dot: *&figures.Dot{
			X: 20,
			Y: 10,
		},
		SideX: 20,
		SideY: 10,
	}

	areaCalc:= figures.New()

	fmt.Println("Circle Area:", areaCalc.VisitForCircle(circle))
	fmt.Println("Square Area:", areaCalc.VisitForSquare(square))
	fmt.Println("Rectangle Area:", areaCalc.VisitForRectangle(rec))
}