package main

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	if v, ok := d[word]; !ok {
		return "", ErrNotFound
	} else {
		return v, nil
	}
}
