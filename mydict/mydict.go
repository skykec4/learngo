package mydict

import (
	"errors"
	"fmt"
)

//Dictionary s
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("단어가 없어서 바꾸지 못해")
	errCantDelete = errors.New("단어가 없어서 delete 못해")
	errWordExists = errors.New("이미 있는 단어다")
)

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}

	return "", errNotFound
}

//Add add
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}

	return nil
}

//GetDic getDic
func (d Dictionary) GetDic() {

	// fmt.Println("dic ", d)

	for i, a := range d {
		fmt.Println(i, " : ", a)
	}
}

//Update update words
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = def

	case errNotFound:
		return errCantUpdate

	}

	return nil

}

//Delete delete words
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)

	case errNotFound:
		return errCantDelete

	}

	return nil

}
