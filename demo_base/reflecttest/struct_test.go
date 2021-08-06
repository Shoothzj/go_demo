package reflecttest

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStruct(t *testing.T) {
	var turbo = &turbo{
		Name: "hzj",
		Age:  1,
	}
	types := reflect.TypeOf(turbo)
	if types.Kind() == reflect.Ptr {
		types = types.Elem()
	}
	for i := 0; i < types.NumField(); i++ {
		tf := types.Field(i)
		fmt.Println(tf.Name, tf.Type)
		fmt.Println(tf.Name, tf.Anonymous)
		fmt.Println(tf.Name, tf.Index)
		if tag, ok := tf.Tag.Lookup("json"); ok {
			fmt.Printf("field %s tag is %s\n", tf.Name, tag)
		}
	}
	if _, ok := types.FieldByName("Class"); !ok {
		fmt.Println("field not exists")
	} else {
		fmt.Printf("field exists")
	}
}
