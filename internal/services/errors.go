package services

type ServiceErrors struct {
	message string
}

func (r ServiceErrors) Error() string {
	return r.message
}
