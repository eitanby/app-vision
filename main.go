package main

import (
	"fmt"
	format "github.com/eitanby/app-vision/do-format"
	"github.com/eitanby/app-vision/math"
)

func main() {
	num := math.Double(4)
	output := format.Number(num)
	fmt.Println(output)
}
