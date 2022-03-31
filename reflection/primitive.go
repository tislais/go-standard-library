package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func main() {
	ourTitle := "the go standard library"
	timed := MakeTimedFunction(properTitle).(func(string) string)
	newTitle := timed(ourTitle)
	fmt.Println(newTitle)
	timedToo := MakeTimedFunction(doubleOurNumber).(func(int) int)
	fmt.Println(timedToo(2))
}

func properTitle(input string) string {
	words := strings.Fields(input)
	smallwords := " a an on the to "

	for i, w := range words {
		if strings.Contains(smallwords, " "+w+" ") {
			words[i] = w
		} else {
			words[i] = strings.Title(w)
		}
	}
	return strings.Join(words, " ")
}

func MakeTimedFunction(f interface{}) interface{} {
	rf := reflect.TypeOf(f)

	if rf.Kind() != reflect.Func {
		panic("expects a function")
	}
	vf := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(rf, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := vf.Call(in)
		end := time.Now()
		fmt.Printf("Calling %s took %v\n", runtime.FuncForPC(vf.Pointer()).Name(), end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}

func doubleOurNumber(a int) int {
	time.Sleep(1 * time.Second)
	return a * 2
}

// more on reflection https://medium.com/capital-one-tech/learning-to-use-go-reflection-822a0aed74b7
