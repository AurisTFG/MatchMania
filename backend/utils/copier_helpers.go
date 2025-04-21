package utils

import (
	"fmt"

	"github.com/jinzhu/copier"
)

func CopyOrPanic[T any](src any) *T {
	var dest T

	if err := copier.Copy(&dest, src); err != nil {
		panic("Failed to copy data from type " + fmt.Sprintf("%T", src) + " to type " + fmt.Sprintf("%T", dest))
	}

	return &dest
}
