/*
mylib is homemade library.
*/
package mylib

// Average returns the average of series of numbers
func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
