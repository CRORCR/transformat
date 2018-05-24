package main

const (
	ErrSuccess          = 10000
	ErrInvalidParameter = 10001
	ErrServerBusy       = 10002
)

func getMessage(code int) (msg string) {
	switch code {
	case ErrSuccess:
		msg = "success"
	case ErrInvalidParameter:
		msg = "invalid parameter"
	case ErrServerBusy:
		msg = "server busy"
	default:
		msg = "unknown error"
	}
	return
}
