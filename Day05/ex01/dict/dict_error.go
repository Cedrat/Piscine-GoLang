package dict

type DictError string

// func New(text string) error {
// 	return &DictError{text}
// }

const (
	ErrorNotFound     = DictError("Word Not Found")
	ErrorAlreadyExist = DictError("Word Already Exist")
	ErrorDoesNotExist = DictError("Word Does Not Exist")
	ErrorNoRemovable  = DictError("Word Cannot Be Removed")
)

func (e DictError) Error() string {

	return string(e)
}
