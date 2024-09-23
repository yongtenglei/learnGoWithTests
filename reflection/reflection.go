package reflection

import "reflect"

func walk(x any, fn func(string)) {
	v := getValue(x)

	numOfValues := 0
	var getFieldFunc func(int) reflect.Value

	switch v.Kind() {
	case reflect.Struct:
		numOfValues = v.NumField()
		getFieldFunc = v.Field
	case reflect.Slice:
		numOfValues = v.Len()
		getFieldFunc = v.Index
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
