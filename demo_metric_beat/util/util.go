package util

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"strconv"
)

func ParseKeyValue(content []byte, separator string, callback func(key, value []byte) error) error {
	sc := bufio.NewScanner(bytes.NewReader(content))
	for sc.Scan() {
		parts := bytes.SplitN(sc.Bytes(), []byte(separator), 2)
		if len(parts) != 2 {
			continue
		}

		if err := callback(parts[0], bytes.TrimSpace(parts[1])); err != nil {
			return err
		}
	}

	return sc.Err()
}

func ParseBytesOrNumber(data []byte) (uint64, error) {
	parts := bytes.Fields(data)

	if len(parts) == 0 {
		return 0, errors.New("empty value")
	}

	num, err := strconv.ParseUint(string(parts[0]), 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse value")
	}

	var multiplier uint64 = 1
	if len(parts) >= 2 {
		switch string(parts[1]) {
		case "kB":
			multiplier = 1024
		default:
			return 0, errors.Errorf("unhandled unit %v", string(parts[1]))
		}
	}

	return num * multiplier, nil
}
