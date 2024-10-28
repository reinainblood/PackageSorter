package sorter

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Dimensions struct {
	Width  float64
	Height float64
	Length float64
	Mass   float64
}

func Sort(width float64, height float64, length float64, mass float64) string {
	if width <= 0 || height <= 0 || length <= 0 || mass <= 0 {
		return "all dimensions and mass must be greater than zero"
	}
	dimensions := Dimensions{Width: width, Height: height, Length: length, Mass: mass}
	bulky := IsBulky(dimensions)
	heavy := IsHeavy(dimensions)
	if bulky && heavy {
		return "REJECTED"
	} else if bulky && !heavy {
		return "SPECIAL"
	} else if !bulky && heavy {
		return "SPECIAL"
	} else {
		return "STANDARD"
	}
}

func IsHeavy(dimensions Dimensions) bool {
	if dimensions.Mass > 20 {
		return true
	} else {
		return false
	}
}

func IsBulky(dimensions Dimensions) bool {
	volume := dimensions.Height * dimensions.Length * dimensions.Width
	if volume > 1000000 {
		return true
	}
	v := reflect.ValueOf(dimensions)

	for i := 0; i < v.NumField(); i++ {
		if v.NumField() > int(150) {
			return true
		}
	}
	return false

}

func SortThings() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: go run app.go <width> <height> <length> <mass>")
		return
	}

	width, err1 := strconv.ParseFloat(os.Args[1], 64)
	height, err2 := strconv.ParseFloat(os.Args[2], 64)
	length, err3 := strconv.ParseFloat(os.Args[3], 64)
	mass, err4 := strconv.ParseFloat(os.Args[4], 64)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Println("All arguments must be numbers")
		return
	}

	result := Sort(width, height, length, mass)
	fmt.Println(result)
}
