package logstack

// =========================================
func statWrapper(statusCode errT, previousStatusCodesSlice []errT) []errT {
	if statusCode == 0 && len(previousStatusCodesSlice) == 0 { //nill or empty
		return nil
	}

	if previousStatusCodesSlice == nil {
		previousStatusCodesSlice = []errT{}
	}
	return append(previousStatusCodesSlice, statusCode)
}

//========================================
