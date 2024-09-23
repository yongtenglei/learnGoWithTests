package reflection

import "reflect"

func walk(x any, fn func(string)) {
	v := getValue(x)

	walkValue := func(v reflect.Value) {
		walk(v.Interface(), fn)
	}

	numOfValues := 0
	var getFieldFunc func(int) reflect.Value

	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			walkValue(v.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			walkValue(v.Index(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			walkValue(v.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := v.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.String:
		fn(v.String())
	}

	for i := 0; i < numOfValues; i++ {
		walk(getFieldFunc(i).Interface(), fn)
	}
}

func getValue(x any) reflect.Value {
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	return v
}
