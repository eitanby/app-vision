// Package format provides various utilities to
// make it easy to convert values to strings
package format

import "fmt"

func Number(num int) string {
	return fmt.Sprintf("the number is %d", num)
}
