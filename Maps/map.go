package maps

import "errors"

type Dictionary map[string]string

var NotFoundError = errors.New("Could not find the word in the given dictionary")

func (d Dictionary) Search(key string) (string, error) {
	value, found := d[key]

	if !found {
		return "", NotFoundError
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
