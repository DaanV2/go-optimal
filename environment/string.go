package env

var String = environment[string]{
    convert: func(value string) (string, error) {
        return value, nil
    },
}
