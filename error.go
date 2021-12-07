package be

type InvalidParamsError struct {
	Message string
}

func (e InvalidParamsError) String() string {
	return e.Message
}

func (e InvalidParamsError) Error() string{
	return e.String()
}


