package main

import (
	"fmt"
	"github.com/eitanby/app-vision/internal/do-format"
	"github.com/eitanby/app-vision/internal/math"
)

func main() {
	num := math.Double(4)
	output := format.Number(num)
	fmt.Println(output)
}
