package pigeon

const (
	ErrNoSuchCommand     = Error("[PIGEON] no such command")
	ErrNoSuchEvent       = Error("[PIGEON] no such event")
	ErrCommandDuplicated = Error("[PIGEON] command duplicated")
	ErrEventDuplicated   = Error("[PIGEON] event duplicated")
	ErrAggregateNotFound = Error("[PIGEON] aggregate not found")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
