package dict

type DictError struct {
	msg string
}

func New(text string) error {
	return &DictError{text}
}

var (
	ErrorNotFound     = New("Word Not Found")
	ErrorAlreadyExist = New("Word Already Exist")
	ErrorDoesNotExist = New("Word Does Not Exist")
	ErrorNoRemovable  = New("Word Cannot Be Removed")
)

func (e *DictError) Error() string {

	return e.msg
}
