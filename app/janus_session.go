package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/go/src/math/rand"
)

type JanusSession struct {
	Transaction string
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewJanusSession(n int) *JanusSession {
	return &JanusSession{
		Transaction: randomString(n),
	}
}

func (s *JanusSession) New() *JanusResponse {
	req := &JanusRequest{
		Janus:       "create",
		Transaction: s.Transaction,
	}
	jsonValue, _ := json.Marshal(req)
	res, err := http.Post(
		"http://localhost:8088/janus",
		"application/json",
		bytes.NewBuffer(jsonValue))

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	var response *JanusResponse
	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}
	return response
}
