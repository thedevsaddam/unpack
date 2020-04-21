package unpack

import (
	"reflect"
)

// Do perform a destructuring on slice values
func Do(src interface{}, dest ...interface{}) {
	srcRV := reflect.ValueOf(src)
	dstRV := reflect.ValueOf(dest)

	dstLen := dstRV.Len()

	if dstLen == 0 {
		return
	}

	if srcRV.Kind() == reflect.Slice ||
		srcRV.Kind() == reflect.Array {
		for i := 0; i < srcRV.Len(); i++ {

			if dstLen == 0 {
				return
			}

			val := srcRV.Index(i)
			v := dstRV.Index(i).Elem().Elem()
			if v.CanSet() {
				v.Set(val)
			}

			dstLen--
		}
	}
}
