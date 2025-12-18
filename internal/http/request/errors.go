package request

type ErrMissingParam struct {
	Key string
}

func (e ErrMissingParam) Error() string {
	return "missing parameter: " + e.Key
}
