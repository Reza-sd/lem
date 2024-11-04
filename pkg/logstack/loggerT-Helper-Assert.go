package logstack

import (
	"fmt"
	"reflect"
	"testing"
)

//--------------------
/*
	%s

This is a verb that formats the first argument as a string.
It's commonly used to print plain text strings.
%+v

This is a verb with a flag +.
The %v part formats the second argument in its default format.
The + flag adds additional information, typically:
For structs: It includes field names along with their values.
For other types: It might provide more detailed output, such as pointers or addresses.
*/
// -------------------
func (f *helperFn[T]) Assert(t testing.TB, got, exp any, inputsArr ...any) {

	t.Helper()

	inputsStr := inputsArrToString(inputsArr)

	if !reflect.DeepEqual(got, exp) {
		t.Errorf("\n>------------------\nInputs:\n%v\nGot=type:(%T) value:(%v)\nExp=type:(%T) value:(%v)\n>------------------\n",
			inputsStr, got, got, exp, exp)
	}

}

// ------------------
func inputsArrToString(inputs []any) string {
	var str string
	var count int
	for _, item := range inputs {
		count++
		str = str + fmt.Sprintf("%d-inp=type:(%T) value:(%v)\n", count, item, item)
	}
	return str
}

//---------------------
