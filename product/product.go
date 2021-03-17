package product

const (
	ErrProductNotFound   = Error("product not found")
	ErrProductBalanceOut = Error("product balance out")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
