package telegram

type ApiError interface {
	Code() int
	Error() string
}

type apiError struct {
	code    int
	message string
}

func (e *apiError) Code() int {
	return e.code
}

func (e *apiError) Error() string {
	return e.message
}
