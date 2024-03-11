package main

import (
	"fmt"
	apputil "github.com/eitanby/app-util"
)

func main() {
	result := apputil.SpecialAdd(5, 5)
	fmt.Printf("result is %d", result)
}
