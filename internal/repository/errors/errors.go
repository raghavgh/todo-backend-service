package errors

type RepoErrors struct {
	Message string
}

func (r RepoErrors) Error() string {
	return r.Message
}
