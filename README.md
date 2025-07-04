# tagerr

tagerr is a go package that provides a hierarchical error type that allows to store the tag of the most inner child of its type and the code of the root error. Additionally, by using `WithStack()` method you can add a stack trace to the error which can be later accessed via `Stack()`.

This is useful when you want the most inner "meaningful" error code while being able to check if error is of one of the outer types.

This package does not use any non-standard error implementations and solely relies on fmt for error wrapping.

## Examples
```go
var ErrProductInUse = tagerr.ErrInvalidReq.Wrap(&tagerr.Err{
    Err: errors.New("product in use by at least one shop"),
    Tag: "product_in_use",
})
```
This yields the following:
```go
&tagerr.Err{
    Err: "invalid request: product in use by at least one shop",
    Tag: "product_in_use",
    HTTPCode: 401,
}
```

You can also wrap other error types inside:
```go
// If we have:
var ErrDBNotHandled = tagerr.ErrInternal.Wrap(&tagerr.Err{
    Err:    errors.New("unhandled database error"),
    Tag:    "unhandled_db_error",
})
// And somewhere in DB layer we have:
return ErrDBNotHandled.Wrap(err).WithStack() // this inner err can be a postgres driver error
```
We then get:
```go
&tagerr.Err{
    Err: "internal server error: unhandled database error: unsupported data type Address..",
    Tag: "unhandled_db_error",
    HTTPCode: 500,
}
```
And `err.Stack()` will return something like this:
```log
github.com/pedramktb/go-tagerr.(*Err).WithStack
	.../db.go:123 -> where we had the return ErrDBNotHandled.Wrap(err).WithStack() call
...
```