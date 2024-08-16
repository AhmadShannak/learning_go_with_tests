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

func (d Dictionary) Update(key, newValue string) DictionaryError {
	_, er := d.Search(key)
	if er != None {
		return er
	}

	d[key] = newValue
	return None
}

func (d Dictionary) Delete(key string) DictionaryError {
	_, er := d.Search(key)
	if er != None {
		return er
	}

	delete(d, key)
	return None
}
