package mask

import (
	"reflect"
)

func Mask(i interface{}) interface{} {
	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Struct:
		return maskStruct(i).Elem().Interface()
	}
	return i
}

func maskStruct(s interface{}) reflect.Value {
	rt := reflect.TypeOf(s)
	if rt.Kind() == reflect.Ptr {
		panic("arg is pointer")
	}

	rv := reflect.ValueOf(&s)
	newRv := reflect.New(rt)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fieldValue := rv.Elem().Elem().FieldByName(field.Name)
		//fmt.Printf("fieldname: %s, value: %#+v\n", field.Name, fieldValue)
		var newFieldValue reflect.Value

		// if field has `mask` tag, masking field value
		if _, ok := field.Tag.Lookup("mask"); ok {
			// TODO: when kind is ptr
			switch fieldValue.Kind() {
			case reflect.Ptr:
				if fieldValue.IsNil() {
					continue
				}
				newFieldValue = maskPtr(fieldValue)
			case reflect.Struct:
				newFieldValue = maskStruct(fieldValue.Interface()).Elem()
			default:
				// mask literals
				newFieldValue = maskLiterals(fieldValue).Elem()
			}
		} else {
			newFieldValue = rv.FieldByName(field.Name)
		}

		newRv.Elem().FieldByName(field.Name).Set(newFieldValue)
	}
	return newRv
}

func maskPtr(fieldValue reflect.Value) reflect.Value {
	switch fieldValue.Kind() {
	case reflect.Struct:
		return maskStruct(fieldValue.Interface())
	case reflect.Ptr:
		return maskPtr(fieldValue.Elem())
	}
	return maskLiterals(fieldValue)
}

// return riterals pointer
func maskLiterals(fieldValue reflect.Value) reflect.Value {
	switch fieldValue.Kind() {
	case reflect.String:
		s := "💩💩💩💩💩"
		return reflect.ValueOf(&s)
	case reflect.Int:
		i := int(-1)
		return reflect.ValueOf(&i)
	case reflect.Int8:
		i := int8(-1)
		return reflect.ValueOf(&i)
	case reflect.Int16:
		i := int16(-1)
		return reflect.ValueOf(&i)
	case reflect.Int32:
		i := int32(-1)
		return reflect.ValueOf(&i)
	case reflect.Int64:
		i := int64(-1)
		return reflect.ValueOf(&i)
	case reflect.Float32:
		f := float32(-1.0)
		return reflect.ValueOf(&f)
	case reflect.Float64:
		f := float64(-1.0)
		return reflect.ValueOf(&f)
	case reflect.Uint:
		ui := uint(1)
		return reflect.ValueOf(&ui)
	case reflect.Uint8:
		ui := uint8(1)
		return reflect.ValueOf(&ui)
	case reflect.Uint16:
		ui := uint16(1)
		return reflect.ValueOf(&ui)
	case reflect.Uint32:
		ui := uint32(1)
		return reflect.ValueOf(&ui)
	case reflect.Uint64:
		ui := uint64(1)
		return reflect.ValueOf(&ui)
	}
	return fieldValue
}
