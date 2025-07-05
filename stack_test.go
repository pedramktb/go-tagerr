package tagerr_test

import (
	"errors"
	"testing"

	"github.com/pedramktb/go-tagerr"
	"github.com/stretchr/testify/assert"
)

var stacklessError = &tagerr.Err{
	Err:      errors.New("some error"),
	Tag:      "some_error",
	HTTPCode: 1000,
}

func stacklessErrorReturningFunction() error {
	return stacklessError
}

func stackfullErrorReturningFunction() error {
	return stacklessError.WithStack()
}

func Test_Stack(t *testing.T) {
	err := stacklessErrorReturningFunction()
	assert.Empty(t, err.(*tagerr.Err).Stack())

	err = stackfullErrorReturningFunction()
	assert.NotContains(t, err.(*tagerr.Err).Stack(), "github.com/pedramktb/go-tagerr.(*Err).WithStack")
	assert.NotContains(t, err.(*tagerr.Err).Stack(), "/err.go:")
	assert.Contains(t, err.(*tagerr.Err).Stack(), "github.com/pedramktb/go-tagerr_test.Test_Stack")
	assert.Contains(t, err.(*tagerr.Err).Stack(), "github.com/pedramktb/go-tagerr_test.stackfullErrorReturningFunction")
	assert.Contains(t, err.(*tagerr.Err).Stack(), "/stack_test.go:")
}
