package be

import "go/types"

type Response struct {
	success bool
	message string
	result  types.Object
}
