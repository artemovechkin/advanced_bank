package customerror

type Error interface {
	Error() string
	Status() int
}

type CustomError struct {
	State   int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) Status() int {
	return e.State
}
