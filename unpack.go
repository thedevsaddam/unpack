package unpack

import (
	"errors"
	"fmt"
	"reflect"
)

// Do perform a destructuring on slice/array values
func Do(src interface{}, dest ...interface{}) error {
	srcRV := reflect.ValueOf(src)
	dstRV := reflect.ValueOf(dest)

	dstLen := dstRV.Len()

	if dstLen == 0 {
		return errors.New("unpack: destination can't be empty")
	}

	if srcRV.Kind() == reflect.Slice ||
		srcRV.Kind() == reflect.Array {
		for i := 0; i < srcRV.Len(); i++ {

			if dstLen == 0 {
				return nil
			}

			if dstRV.Index(i).Elem().Kind() != reflect.Ptr {
				return errors.New("unpack: the destination is not pointers")
			}

			val := srcRV.Index(i)
			v := dstRV.Index(i).Elem().Elem()
			if val.Kind() != v.Kind() {
				return fmt.Errorf("unpack: can't assign value from source type %v to destination type %v", val.Kind(), v.Kind())
			}
			if v.CanSet() {
				v.Set(val)
			}

			dstLen--
		}
	}

	return nil
}
