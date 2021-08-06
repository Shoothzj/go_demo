package reflecttest

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValueOfInt(t *testing.T) {
	var a int = 56
	value := reflect.ValueOf(a)
	fmt.Println(value.Interface())
	fmt.Println(value.Interface().(int))
}

func TestValueOfStruct(t *testing.T) {
	turbo := &turbo{
		Name: "hzj",
		Age:  1,
	}
	value := reflect.ValueOf(turbo)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			fmt.Printf("field type %v, file value%v\n", field.Type(), field.Interface())
		}
		st := value.FieldByName("Name")
		fmt.Printf("%v\n", st.Interface())
	}
}
