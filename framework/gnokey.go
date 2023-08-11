package gnomobile

import "fmt"

type PromiseBlock interface {
	CallResolve(reply string)
	CallReject(error error)
}

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
