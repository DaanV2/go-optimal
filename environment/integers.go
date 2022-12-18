package env

import "strconv"

var Int64 = environment[int64]{
	convert: func(value string) (int64, error) {
		return strconv.ParseInt(value, 10, 64)
	},
}

var Int32 = environment[int32]{
	convert: func(value string) (int32, error) {
		v, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return 0, err
		}

		return int32(v), nil
	},
}

