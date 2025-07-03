package tagerr

import (
	"errors"
	"fmt"

	grpccodes "google.golang.org/grpc/codes"
)

// *tagerr.Err is an error implementation that in a nested wrapping provides the tag of the most inner child and the code of the root Error.
// This is useful when you want the most inner "meaningful" error code while being able to check if error is of one of the outer types.
// The root errors should carry interface codes (e.g. http or grpc codes) to allow easy DTO mapping.
// The most common root errors are defined in this package.
type Err struct {
	Err      error
	Tag      string
	HTTPCode int
	GRPCCode grpccodes.Code
}

func (e *Err) Error() string {
	return e.Err.Error()
}

// Wrap a target error inside this *tagerr.Err.
// If the target error is also a *tagerr.Err, target's tag will be used.
func (e *Err) Wrap(target error) *Err {
	if target, ok := target.(*Err); ok {
		return &Err{
			Err:      fmt.Errorf("%w: %w", e.Err, target.Err),
			Tag:      target.Tag,
			HTTPCode: e.HTTPCode,
			GRPCCode: e.GRPCCode,
		}
	}
	return &Err{
		Err:      fmt.Errorf("%w: %w", e.Err, target),
		Tag:      e.Tag,
		HTTPCode: e.HTTPCode,
		GRPCCode: e.GRPCCode,
	}
}

// Returns true if the errors.Is() on the underlying error returns true.
// If the given error is a *tagerr.Err, its underlying error will be used.
func (e *Err) Is(target error) bool {
	if target == nil {
		return e == nil
	}
	if target, ok := target.(*Err); ok {
		return errors.Is(e.Err, target.Err)
	}
	return errors.Is(e.Err, target)
}

// Wrap a target error inside an err.
// If the wrapped error is a *tagerr.Err, *tagerr.Err.Wrap(target) method will be used.
func Wrap(err, target error) error {
	if err, ok := err.(*Err); ok {
		return err.Wrap(target)
	}
	return fmt.Errorf("%w: %w", err, target)
}

// Checks if the given error wraps or is equal to the target error.
// If the wrapping error is a *tagerr.Err, *tagerr.Err.Is(target) will be used.
func Is(err error, target error) bool {
	if err, ok := err.(*Err); ok {
		return err.Is(target)
	}
	return errors.Is(err, target)
}
