package tagerr

import (
	"errors"
	"net/http"

	grpccodes "google.golang.org/grpc/codes"
)

// Root *tagerr.Err errors that hold the most common HTTP and gRPC codes.
var (
	ErrInvalidReq = &Err{
		Err:      errors.New("invalid request"),
		Tag:      "invalid_request",
		HTTPCode: http.StatusBadRequest,
		GRPCCode: grpccodes.InvalidArgument,
	}
	ErrOutOfRange = &Err{
		Err:      Wrap(ErrInvalidReq.Err, errors.New("out of range")),
		Tag:      "out_of_range",
		HTTPCode: http.StatusBadRequest, // Shared with ErrInvalidReq as a case of it
		GRPCCode: grpccodes.OutOfRange,
	}
	ErrNotAuth = &Err{
		Err:      errors.New("not authenticated"),
		Tag:      "not_authenticated",
		HTTPCode: http.StatusUnauthorized,
		GRPCCode: grpccodes.Unauthenticated,
	}
	ErrNotAllowed = &Err{
		Err:      errors.New("not allowed"),
		Tag:      "not_allowed",
		HTTPCode: http.StatusForbidden,
		GRPCCode: grpccodes.PermissionDenied,
	}
	ErrNotFound = &Err{
		Err:      errors.New("not found"),
		Tag:      "not_found",
		HTTPCode: http.StatusNotFound,
		GRPCCode: grpccodes.NotFound,
	}
	ErrReqTimeout = &Err{
		Err:      errors.New("request timeout"),
		Tag:      "request_timeout",
		HTTPCode: http.StatusRequestTimeout,
		GRPCCode: grpccodes.DeadlineExceeded, // Shared with ErrUpstreamTimeout
	}
	ErrAlreadyExists = &Err{
		Err:      errors.New("already exists"),
		Tag:      "already_exists",
		HTTPCode: http.StatusConflict, // Shared with ErrAborted
		GRPCCode: grpccodes.AlreadyExists,
	}
	ErrAborted = &Err{
		Err:      errors.New("aborted"),
		Tag:      "aborted",
		HTTPCode: http.StatusConflict, // Shared with ErrAlreadyExists
		GRPCCode: grpccodes.Aborted,
	}
	ErrFailedPreCond = &Err{
		Err:      errors.New("failed precondition"),
		Tag:      "failed_precondition",
		HTTPCode: http.StatusPreconditionFailed,
		GRPCCode: grpccodes.FailedPrecondition,
	}
	ErrRateLimit = &Err{
		Err:      errors.New("rate limit exceeded"),
		Tag:      "rate_limit_exceeded",
		HTTPCode: http.StatusTooManyRequests,
		GRPCCode: grpccodes.ResourceExhausted,
	}
	ErrCanceled = &Err{
		Err:      errors.New("canceled"),
		Tag:      "canceled",
		HTTPCode: 499, // non-standard status code (e.g. used by Nginx) for client cancellation
		GRPCCode: grpccodes.Canceled,
	}
	ErrInternal = &Err{
		Err:      errors.New("internal error"),
		Tag:      "internal_error",
		HTTPCode: http.StatusInternalServerError,
		GRPCCode: grpccodes.Internal,
	}
	ErrDataLoss = &Err{
		Err:      Wrap(ErrInternal, errors.New("data loss")),
		Tag:      "data_loss",
		HTTPCode: http.StatusInternalServerError, // shared with ErrInternal as a case of it
		GRPCCode: grpccodes.DataLoss,
	}
	ErrNotImpl = &Err{
		Err:      errors.New("not implemented"),
		Tag:      "not_implemented",
		HTTPCode: http.StatusNotImplemented,
		GRPCCode: grpccodes.Unimplemented,
	}
	ErrUnavailable = &Err{
		Err:      errors.New("unavailable"),
		Tag:      "unavailable",
		HTTPCode: http.StatusServiceUnavailable,
		GRPCCode: grpccodes.Unavailable,
	}
	ErrUpstreamTimeout = &Err{
		Err:      errors.New("upstream timeout"),
		Tag:      "upstream_timeout",
		HTTPCode: http.StatusGatewayTimeout,
		GRPCCode: grpccodes.DeadlineExceeded, // Shared with ErrReqTimeout
	}
)
