//File utils.go

package utils

// Contains provides search the word X in the slice A
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// ContainsInt provides search the digit X in the slice of ints A
func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
