package main

import "fmt"

func main() {
	session := NewJanusSession(12)
	fmt.Println(session)

	r := session.New()
	fmt.Println(r)
}
