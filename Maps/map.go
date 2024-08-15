package maps

type Dictionary map[string]string
type DictionaryError string

const (
	NotFoundError     = DictionaryError("Could not find the word in the given dictionary")
	AlreadyExistError = DictionaryError("Could not add the element because it already exist")
	None              = DictionaryError("")
)

func (d Dictionary) Search(key string) (string, DictionaryError) {
	value, found := d[key]

	if !found {
		return "", NotFoundError
	}
	return value, None
}

func (d Dictionary) Add(key, value string) DictionaryError {
	_, error := d.Search(key)
	if error == None {
		return AlreadyExistError
	}
	d[key] = value
	return None
}
