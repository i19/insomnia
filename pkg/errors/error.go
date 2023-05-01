package errors

type ErrorCode int

const (
	OK ErrorCode = iota
)
const (
	BadRequest ErrorCode = iota + 400
	Unauthorized
	Forbidden
)
const (
	InternalServerError ErrorCode = iota + 500
)
