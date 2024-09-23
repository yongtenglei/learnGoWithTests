package reflection

import "reflect"

func walk(x any, fn func(string)) {
	v := reflect.ValueOf(x)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
