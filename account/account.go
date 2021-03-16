package account

const (
	ErrAccountNotFound    = Error("account not found")
	ErrEmailAlreadyExists = Error("email already exists")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
