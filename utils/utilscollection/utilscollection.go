package utilscollection

// IsEmptyInterfaceArray : method to check array interface is empty
func IsEmptyInterfaceArray(data []interface{}) bool {
	return len(data) == 0
}

// IsNotEmptyInterfaceArray : method to check array interface is not empty
func IsNotEmptyInterfaceArray(data []interface{}) bool {
	return !IsEmptyInterfaceArray(data)
}
