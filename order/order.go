package order

const (
	ErrOrderNotFound = Error("order not found")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
