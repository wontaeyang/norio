package errs

import "fmt"

func Panic(message string, err error) {
	fmt.Printf("Error: %s\n", message)
	panic(err)
}
