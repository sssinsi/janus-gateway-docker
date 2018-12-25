package main

type JanusRequest struct {
	Janus       string `json:"janus"`
	Transaction string `json:"transaction"`
	Plugin      string `json:"plugin"`
	OpaqueID    string `json:"opaque_id"`
}
