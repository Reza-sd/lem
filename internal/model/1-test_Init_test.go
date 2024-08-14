package modelpk

import (
	"testing"
)

func Test_Init(t *testing.T) {

	t.Run(`1-Print`, func(t *testing.T) {
		myLem := Sample_Lem1()

		var model1 Model
		model1.Init(&myLem)
		model1.Print()
	})
}
