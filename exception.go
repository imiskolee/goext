package goext

//Exception a exception impl
type Exception struct {
	Code    int
	Message string
}

func Throw(e *Exception) {
	panic(e)
}
