package modelpk

import (
	"testing"
)

func Test_Init(t *testing.T) {

	t.Run(`1-Print`, func(t *testing.T) {
		myLem := Sample_Lem_1ant_2room()

		var model1 Model
		model1.Init(&myLem)
		model1.Print()
	})
	t.Run(`2-Print`, func(t *testing.T) {
		myLem := Sample_Lem_2ant_2room()

		var model1 Model
		model1.Init(&myLem)
		model1.Print()
	})
	t.Run(`2-Print`, func(t *testing.T) {
		myLem := Sample_Lem_2ant_3room()

		var model1 Model
		model1.Init(&myLem)
		model1.Print()
	})
}
