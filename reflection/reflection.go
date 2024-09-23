package reflection

import "reflect"

func walk(x any, f func(string)) {
	v := reflect.ValueOf(x)

	for i := 0; i < v.NumField(); i++ {
		filed := v.Field(i)
		f(filed.String())

	}
}
