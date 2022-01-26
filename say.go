package RebornGoProject

import "fmt"

// say Hi to someone
func SayHi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

func SayHello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func SayNo(name string) string {
	return fmt.Sprintf("No, %s", name)
}
