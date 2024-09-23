package reflection

import "reflect"

func walk(x any, fn func(string)) {
	v := getValue(x)

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			walk(v.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			walk(v.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(v.String())
	}
}

func getValue(x any) reflect.Value {
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	return v
}
