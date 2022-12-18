package env

import "strings"

// String able to get environment variables and convert them to string.
var String = EnvironmentType[string]{
	convert: func(value string) (string, error) {
		return value, nil
	},
}

// CSV able to get environment variables and convert them to []string.
var CSV = EnvironmentType[[]string]{
	convert: func(value string) ([]string, error) {
		return strings.Split(value, ","), nil
	},
}
