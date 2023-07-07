// File contains all errors related to not specific data structures.

package dataStructures

import "fmt"

type InvalidDataError struct {
	Err   error
	Data  string
	Field string
}

func (e *InvalidDataError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("Invalid data was given: %v", e.Err)
	}
	return fmt.Sprintf("Invalid data was given for %s field: \"%s\"", e.Field, e.Data)
}
