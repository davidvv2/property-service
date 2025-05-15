package structure

import "unsafe"

func ArrayCast[castTo, castFrom any](array *castFrom) *castTo {
	up := unsafe.Pointer(array)
	return (*castTo)(up)
}
