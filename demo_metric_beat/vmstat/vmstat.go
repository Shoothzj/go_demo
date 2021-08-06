package vmstat

import (
	"demo_metric-beat/module"
	"demo_metric-beat/util"
	"github.com/pkg/errors"
	"reflect"
)

// vmstatTagToFieldIndex contains a mapping of json struct tags to struct field indices.
var vmstatTagToFieldIndex = make(map[string]int)

func init() {
	var vmstat module.VMStatInfo
	val := reflect.ValueOf(vmstat)
	typ := reflect.TypeOf(vmstat)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if tag := field.Tag.Get("json"); tag != "" {
			vmstatTagToFieldIndex[tag] = i
		}
	}
}

func parseVMStat(content []byte) (*module.VMStatInfo, error) {
	var vmStat module.VMStatInfo
	refValues := reflect.ValueOf(&vmStat).Elem()

	err := util.ParseKeyValue(content, " ", func(key, value []byte) error {
		// turn our []byte value into an int
		val, err := util.ParseBytesOrNumber(value)
		if err != nil {
			return errors.Wrapf(err, "failed to parse %v value of %v", string(key), string(value))
		}

		idx, ok := vmstatTagToFieldIndex[string(key)]
		if !ok {
			return nil
		}

		sval := refValues.Field(idx)
		if sval.CanSet() {
			sval.SetUint(val)
		}
		return nil
	})

	return &vmStat, err
}
