package main

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	NotFoundErr         = DictionaryErr("Word not available in this Dictionary")
	WordExistsErr       = DictionaryErr("Word already exists in this Dictionary")
	WordDoesNotExistErr = DictionaryErr("Word does not exists in this Dictionary")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", NotFoundErr
	}

	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case NotFoundErr:
		d[word] = definition
	case nil:
		return WordExistsErr
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case NotFoundErr:
		return WordDoesNotExistErr
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case NotFoundErr:
		return err
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
