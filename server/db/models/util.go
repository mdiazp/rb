package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// ValidationError ...
type ValidationError struct {
	PropertyName string
	Error        string
}

// Model ...
type Model interface {
	Valid() *[]ValidationError
	AddSQLConstrainsts(db *gorm.DB) error
}

/////////////////  Validating Functions ///////////////////////////////////////

func validateString(es *[]ValidationError, value string,
	propertyName string, minSize int, maxSize int) {
	ln := len(value)

	if ln < minSize {
		*es = append(*es, ValidationError{
			PropertyName: propertyName,
			Error: fmt.Sprintf("%s's size can't be less than %d",
				propertyName, minSize),
		})
	} else if maxSize < ln {
		*es = append(*es, ValidationError{
			PropertyName: propertyName,
			Error: fmt.Sprintf("%s's size can't be greater than %d",
				propertyName, maxSize),
		})
	}
}

func validateInt(es *[]ValidationError, value int,
	propertyName string, minv int, maxv int) {

	if value < minv {
		*es = append(*es, ValidationError{
			PropertyName: propertyName,
			Error:        fmt.Sprintf("%s can't be less than %d", propertyName, minv),
		})
	} else if maxv < value {
		*es = append(*es, ValidationError{
			PropertyName: propertyName,
			Error:        fmt.Sprintf("%s can't be grater than %d", propertyName, maxv),
		})
	}
}

func validateTrue(es *[]ValidationError, propertyName string, valid bool) {
	if !valid {
		*es = append(*es, ValidationError{
			PropertyName: propertyName,
			Error:        fmt.Sprintf("%s has an invalid value", propertyName),
		})
	}
}
