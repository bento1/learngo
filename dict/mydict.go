package mydict

import "errors"

//Dictionary dictionary 장점은 method를 구현할수 있다.
type Dictionary map[string]string

type Money int

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("Already Existed")

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	//a못찾으면 에러를 발생시킴
	//에러를 처리한다.
	//map이 해줌
	// i,ok:=
	value, exist := d[word]
	if exist {
		return value, nil
	}
	return "", errNotFound
}
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil

}
func (d Dictionary) Update(word string, def string) error {
	_, err := d.Search(word)
	if err == nil {
		d[word] = def
	} else if err != nil {
		return errNotFound
	}
	return nil
}
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == nil {
		delete(d, word)
	} else if err != nil {
		return errNotFound
	}
	return nil
}
