package logstack

var errCodeDes1 = map[uint16]string{
	10: "10-not happy",
	11: "11-is not ok",
	12: "12-cant handle user",
}

//var sampleLogger2 = BuildNewLogger(pkgName, errCodeDes1, true, true)

var sampleLogger1 = Help_BuildNewLogger(_PKG_NAME, errCodeDes1, true, true)

///

///
