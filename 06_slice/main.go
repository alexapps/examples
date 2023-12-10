package main

import (
	"fmt"
	"reflect"
)

/*
   implementing an "own" slice of strings values
*/

type MySlice struct {
	length   int
	capacity int
	offset   int
	array    reflect.Value
}

func (s MySlice) Len() int {
	return s.length
}

func (s MySlice) Cap() int {
	return s.capacity
}

func MySliceArray(a any, low, high int) MySlice {
	// Validate that a is not a nil pointer
	ptr := reflect.ValueOf(a)
	if ptr.Kind() != reflect.Pointer || ptr.IsNil() {
		panic("can only slice a non-nil pointer")
	}

	// Check if a points to array of strings
	v := ptr.Elem()
	if v.Kind() != reflect.Array || v.Type().Elem().Kind() != reflect.String {
		panic("can only slice array string")
	}

	// Validate bounds
	if low < 0 || high > v.Len() || low > high {
		panic(fmt.Sprintf("slice bounds out of range [%d:%d]", low, high))
	}
	// Create MySlice struct

	return MySlice{
		array:    v,
		offset:   low,
		length:   high - low,
		capacity: v.Len() - low,
	}

}

func (s MySlice) Get(x int) string {
	if x < 0 || x >= s.length {
		panic(fmt.Sprintf("index out of the range [%d] with length %d", x, s.length))
	}

	// retrieve the element
	return s.array.Index(s.offset + x).String()
}

func (s MySlice) Set(x int, value string) {
	if x < 0 || x >= s.length {
		panic(fmt.Sprintf("index out of the range [%d] with length %d", x, s.length))
	}

	// Set the element
	s.array.Index(s.offset + x).SetString(value)

}

func (s MySlice) String() string {
	out := "["
	for i := 0; i < s.length; i++ {
		if i > 0 {
			// add space between elements
			out += " "
		}
		out += s.Get(i)
	}

	return out + "]"
}

func main() {
	a := [...]string{"a", "b", "c", "d", "e", "f", "g"}

	s := MySliceArray(&a, 0, 3)
	fmt.Printf("result %T, %v", s, s)
}
