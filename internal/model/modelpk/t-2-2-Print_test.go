package modelpk

import (
	"testing"
)

func Test_Print(t *testing.T) {
	t.Skip()
	t.Run(`1-Print`, func(t *testing.T) {

		model1 := mySampleModel.Sample_Model1()
		model1.Print()
	})
}
