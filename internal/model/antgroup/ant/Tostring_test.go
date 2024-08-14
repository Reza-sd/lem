package ant

import (
	
	"testing"
)


func Test_ToString(t *testing.T) {
	t.Skip()
	//------------
	t.Run(`1-Print`, func(t *testing.T) {
		var myAnt1 Ant
		myAnt1.Init(44)
		myAnt1.Print()
	})
}