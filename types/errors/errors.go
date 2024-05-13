package errors

import "fmt"

//에러처리 패키지

const (
	NotFoundUser = iota
)

var errMessage = map[int]string{
	NotFoundUser: "not Found User",
}

func ErrorOf(code int, args ...any) error {

	if message, ok := errMessage[code]; ok {
		return fmt.Errorf("%s : %v", message, args)
	} else {
		return fmt.Errorf("not found err code")
	}
}
