package mydict

import "errors"

// Dictionary Type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Cant update non-exisiting word")
	errWordExists = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	// Way 1 using if
	/*
		if err == errNotFound {
			d[word] = def
		} else if err == nil {
			return errWordExists
		}
		return nil
	*/
	// Way 2 using switch
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil

}

// Update a word to the dictionary
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil: //단어가 있다는 뜻
		d[word] = def
	case errCantUpdate:
		return errCantUpdate
	}
	return nil
}

// Delete a word to the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
