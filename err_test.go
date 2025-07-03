package tagerr_test

import (
	"errors"
	"testing"

	"github.com/pedramktb/go-tagerr"
)

func Test_Wrap_And_Is(t *testing.T) {
	root := &tagerr.Err{
		Err:      errors.New("Root"),
		Tag:      "Root",
		HTTPCode: 1000,
	}

	level1Part := errors.New("Level1")
	level1 := root.Wrap(&tagerr.Err{
		Err: level1Part,
		Tag: "Level1",
	})

	level2Part := errors.New("Level2")
	level2 := level1.Wrap(level2Part)

	tests := []struct {
		name    string
		err     error
		target  error
		isError bool
	}{
		{
			name:    "equal #1",
			err:     root,
			target:  root,
			isError: true,
		},
		{
			name:    "equal #2",
			err:     level1,
			target:  level1,
			isError: true,
		},
		{
			name:    "equal #3",
			err:     level2,
			target:  level2,
			isError: true,
		},
		{
			name:    "wrapper #1",
			err:     level1,
			target:  level1Part,
			isError: true,
		},
		{
			name:    "wrapper #2",
			err:     level2,
			target:  level2Part,
			isError: true,
		},
		{
			name:    "wrapped #1",
			err:     level1Part,
			target:  level1,
			isError: false,
		},
		{
			name:    "wrapped #2",
			err:     level2Part,
			target:  level2,
			isError: false,
		},
		{
			name:    "as parent #1",
			err:     level1,
			target:  root,
			isError: true,
		},
		{
			name:    "as parent #2",
			err:     level2,
			target:  level1,
			isError: true,
		},
		{
			name:    "as grand parent",
			err:     level2,
			target:  root,
			isError: true,
		},
		{
			name:    "as child #1",
			err:     root,
			target:  level1,
			isError: false,
		},
		{
			name:    "as child #2",
			err:     level1,
			target:  level2,
			isError: false,
		},
		{
			name:    "as child #3",
			err:     root,
			target:  level2,
			isError: false,
		},
		{
			name:    "not related",
			err:     root,
			target:  errors.New("NotRoot"),
			isError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tagerr.Is(tc.err, tc.target); got != tc.isError {
				t.Errorf("IsError() = %v, want %v", got, tc.isError)
			}
		})
	}
}
