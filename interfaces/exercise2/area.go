package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	shapeArgPtr := flag.String("shape", "", "Shape type")
	radiusArgPtr := flag.Float64("radius", -1, "Circle radius")
	heightArgPtr := flag.Float64("height", -1, "Rectangle height")
	widthArgPtr := flag.Float64("width", -1, "Rectangle width")

	flag.Parse()

	if *shapeArgPtr == "" {
		printErrorAndExit("Shape type is not defined")
	}

	var shape Shape

	switch *shapeArgPtr {
	case "circle":
		switch {
		case *radiusArgPtr == -1:
			printErrorAndExit("Radius is not defined")
		case *radiusArgPtr <= 0:
			printErrorAndExit("Radius must be > 0")
		}
		shape = Circle{*radiusArgPtr}
	case "rectangle":
		switch {
		case *heightArgPtr == -1:
			printErrorAndExit("Height is not defined")
		case *heightArgPtr <= 0:
			printErrorAndExit("Height must be > 0")
		case *widthArgPtr == -1:
			printErrorAndExit("Width is not defined")
		case *widthArgPtr <= 0:
			printErrorAndExit("Width must be > 0")
		}
		shape = Rectangle{*heightArgPtr, *widthArgPtr}
	default:
		printErrorAndExit("Incorrect shape type")
	}

	area := shape.Area()
	fmt.Println(area)
}

func printErrorAndExit(message string) {
	fmt.Println(message)
	fmt.Println("Usage of program:")
	flag.PrintDefaults()
	os.Exit(1)
}
