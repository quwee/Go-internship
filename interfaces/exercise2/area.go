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
		if *radiusArgPtr < 0 {
			printErrorAndExit("Radius is not defined")
		}
		shape = Circle{*radiusArgPtr}
	case "rectangle":
		if *heightArgPtr < 0 {
			printErrorAndExit("Height is not defined")
		}
		if *widthArgPtr < 0 {
			printErrorAndExit("Width is not defined")
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
