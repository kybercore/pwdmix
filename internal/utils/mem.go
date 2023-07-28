package utils

// OverwriteMemory overwrites the memory of the password with 0s
func OverwriteMemory(password *[]byte) {
	b := []byte(*password) // convert to mutable []byte
	for i := range b {
		b[i] = 0 // overwrite bytes
	}
	*password = b // convert back to string
}
