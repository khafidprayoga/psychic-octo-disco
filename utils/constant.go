package utils

// Serrver Error
const (
	DatabaseError = iota + 5001
)

// User error
const (
	HTTPRequestErr = iota + 4001
	NotFoundErr
)
