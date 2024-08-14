package modelpk

import (
	"testing"
)

func Test_Print(t *testing.T) {

	t.Run(`1-Print`, func(t *testing.T) {

		model1 := Sample_Model1()
		model1.Print()
	})
}
