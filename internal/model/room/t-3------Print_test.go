package room

import "testing"

func Test_Print(t *testing.T){
	t.Run("1-",func(t *testing.T) {
		sample:=SampleRoom{}
		myRoom := sample.Sample_Room_1()
		myRoom.Print()
	})
}