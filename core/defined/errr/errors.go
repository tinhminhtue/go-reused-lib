package errr

import "errors"

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
