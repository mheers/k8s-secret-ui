package helpers

import "errors"

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// Remove takes a slice and removes a specified value
func Remove(slice []string, val string) ([]string, error) {
	i, found := Find(slice, val)
	if !found {
		return slice, errors.New("value to be removed not found")
	}
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1], nil
}
