//File utils.go

package utils

// InSlice provides search the word X in the slice A
func InSlice(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// InSliceInt provides search the digit X in the slice of ints A
func InSliceInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
