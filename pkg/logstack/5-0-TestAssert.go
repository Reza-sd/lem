package logstack

import (
	"fmt"
	"reflect"
	"testing"
)

//-------------------
func (l *loggerT) Assert(t testing.TB, got, exp any, inputsArr ...any) {
	t.Helper()
	inputsStr := inputsArrToString(inputsArr)

	if !reflect.DeepEqual(got, exp) {
		t.Errorf("\n>------------------\nInputs:\n%v\nGot=type:(%T) value:(%v)\nExp=type:(%T) value:(%v)\n>------------------\n",
			inputsStr, got, got, exp, exp)
	}
}

//------------------
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
