package env

import "strconv"

// Int64 able to get environment variables and convert them to int64.
var Int64 = EnvironmentType[int64]{
	convert: func(value string) (int64, error) {
		return strconv.ParseInt(value, 10, 64)
	},
}

// Int32 able to get environment variables and convert them to int32.
var Int32 = EnvironmentType[int32]{
	convert: func(value string) (int32, error) {
		v, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return 0, err
		}

		return int32(v), nil
	},
}

// Int16 able to get environment variables and convert them to int16.
var Int16 = EnvironmentType[int16]{
	convert: func(value string) (int16, error) {
		v, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return 0, err
		}

		return int16(v), nil
	},
}

// Int8 able to get environment variables and convert them to int8.
var Int8 = EnvironmentType[int8]{
	convert: func(value string) (int8, error) {
		v, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return 0, err
		}

		return int8(v), nil
	},
}

// Int able to get environment variables and convert them to int.
var Int = EnvironmentType[int]{
	convert: func(value string) (int, error) {
		v, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return 0, err
		}

		return int(v), nil
	},
}
