package gorest

import (
	"fmt"
)

func (res Response) Send(message string) {
	fmt.Fprint(res, message)
}
func (res Response) Status(statusCode int) {
	res.WriteHeader(statusCode)
}
