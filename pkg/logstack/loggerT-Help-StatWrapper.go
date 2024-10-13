package logstack

// =========================================
func (f *helperFn[T]) StatWrapper(statusCode T, previousStatusCodesSlice []T) []T {
	if statusCode == 0 && len(previousStatusCodesSlice) == 0 { //nill or empty
		return nil
	}

	if previousStatusCodesSlice == nil {
		previousStatusCodesSlice = []T{}
	}
	return append(previousStatusCodesSlice, statusCode)
}

//========================================
