package main

type JanusHandle struct {
	ID      int `json:"id"`
	session *JanusSession
}
