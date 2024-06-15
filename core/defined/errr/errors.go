package errr

import (
	"errors"

	"go.temporal.io/sdk/temporal"
)

// Error codes
var (
	ErrNamespaceIdInvalid = errors.New("namespaceId is invalid")
	ErrEntityIdInvalid    = errors.New("entityId is invalid")
	ErrInvalidInput       = errors.New("invalid input")
	ErrInvalidOutput      = errors.New("invalid output")
	ErrInvalidMeta        = errors.New("invalid meta")
	ErrInvalidMetad       = errors.New("invalid metad")
	ErrInvalidEntity      = errors.New("invalid entity")
)

// Io error codes
var (
	ErrIoNamespaceIdIsRequired   = errors.New("namespaceId is required")
	ErrIoTimelineIdIsRequired    = errors.New("timelineId is required")
	ErrIoBeginTimeIsAfterEndTime = errors.New("begin time required after end time")
	ErrIoIdIsRequired            = errors.New("id is required")
)

// ErrRetryables string enum: ['ErrRetryable', 'ErrNonRetryable']
type ErrRetryables string

const (
	ErrRetryable    ErrRetryables = "ErrRetryable"
	ErrNonRetryable ErrRetryables = "ErrNonRetryable"
)

// custom error class
type ErrCustom struct {
	Message   string
	Type      string
	Cause     error
	Retryable ErrRetryables
}

// factory function
func NewErrCustom(message string, err error, retryable ErrRetryables) *ErrCustom {
	return &ErrCustom{
		Message:   message,
		Type:      err.Error(),
		Cause:     err,
		Retryable: retryable,
	}
}

func (e *ErrCustom) Error() string {
	return e.Message
}

func CastTempoError(e error) error {
	if e != nil {
		// check if e is errr.ErrCustom
		if e, ok := e.(*ErrCustom); ok {
			if e.Retryable == ErrNonRetryable {
				return temporal.NewNonRetryableApplicationError(e.Message, e.Type, e.Cause, nil)
			}
		}
		return e
	}
	return nil
}
