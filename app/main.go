package main

import "fmt"

const BaseURL = "http://localhost:8088/janus"

func main() {
	g := NewJanusGateway(12)

	s := g.NewSession()

	h := s.NewRoom()
	fmt.Println(h.ID)
}
