package models

// NumericValue is an interface for any model that holds a calculable integer value.
// This allows handlers and repositories to work with different types of numbers
// in a generic way.
type NumericValue interface {
	GetValue() int
}
