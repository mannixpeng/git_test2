package main

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestErrorHandler(t *testing.T) {
	err := A()
	fmt.Printf("%+v\n", err)
	//fmt.Println(err)
}

func A() error{
	return B()
}

func B() error{
	err := fmt.Errorf("err failed")
	//err := errors.New("err failed")
	wErr := errors.Wrap(err, "test error")
	//wErr := errors.WithStack(err)
	//wErr := errors.WithMessage(err, "sdfsss")
	//fmt.Printf("%+v", mErr)
	//fmt.Printf("%+v", wErr)
	return wErr
}
