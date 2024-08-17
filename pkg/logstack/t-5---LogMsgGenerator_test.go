package logstack

import (
	"testing"
)

func Test_Init(t *testing.T) {

	t.Run(`1-Print`, func(t *testing.T) {
	SampleLogger.ErrLog("FuncName","OpName","err","OpDes")
	})
}

