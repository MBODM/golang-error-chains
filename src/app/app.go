package app

type CustomError struct {
	Err error
	Msg string
}

func (e *CustomError) Error() string {
	return e.Msg
}

func (e *CustomError) Unwrap() error {
	return e.Err
}
