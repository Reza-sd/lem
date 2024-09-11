package logstack

// =========================================
func statwrapper(statusCode uint8, previousStatusCodesSlice []uint8) []uint8 {
	if statusCode == 0 && len(previousStatusCodesSlice) == 0 { //nill or empty
		return nil
	}

	if previousStatusCodesSlice == nil {
		previousStatusCodesSlice = []uint8{}
	}
	return append(previousStatusCodesSlice, statusCode)
}

//========================================
