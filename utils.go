//File utils.go

package utils

// Contains provides search of work X in the slice A
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
