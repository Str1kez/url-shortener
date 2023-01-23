package db

type NoResultFound struct{}

func (err *NoResultFound) Error() string {
	return "Error: No result found for this request"
}

type Timeout struct{}

func (err *Timeout) Error() string {
	return "Error: Timeout is reached"
}
