package main

import (
	"fmt"
	"reflect"
)

func typeDetector(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan interface{}:
		return "chan"
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			return "chan"
		}
		return "unknown"
	}
}

func main() {
	var chInt chan int
	var chStr chan string
	var chEmpty chan interface{}

	testValues := []interface{}{
		1,
		"foo",
		false,
		chInt,
		chStr,
		chEmpty,
		// для теста на неизвестный тип
		3.14,
		[3]int{3, 5, 7},
	}
	for _, v := range testValues {
		types := typeDetector(v)
		fmt.Printf("Значение: %v → Тип: %s\n", v, types)
	}
}
