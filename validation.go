package validation

import (
	"reflect"

	"github.com/asaskevich/govalidator"
)

func init() {
	// arrayValidator validator array's sub array struct
	govalidator.CustomTypeTagMap.Set("arrayStructValidator", func(i, o interface{}) bool {
		typeof := reflect.TypeOf(i)
		values := reflect.ValueOf(i)
		switch typeof.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < values.Len(); i++ {
				ok, err := govalidator.ValidateStruct(values.Index(i))
				if err != nil {
					panic(err)
				}
				if !ok {
					return false
				}
			}
			return true
		}
		return false
	})

	// booleanValidator validate bool value
	govalidator.CustomTypeTagMap.Set("booleanRequiredValidator", func(i, o interface{}) bool {
		switch i.(type) {
		case bool:
			return true
		case *bool:
			return true
		}
		return false
	})
}
