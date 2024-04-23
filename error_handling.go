package main

import "fmt"

type ErrHandling struct {
	err error
}

func (ex *ErrHandling) ShowError() {
	fmt.Println(ex.err)
}

func CreateErrorHandler(err error) *ErrHandling {
	return &ErrHandling{
		err: err,
	}
}
