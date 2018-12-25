package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/go/src/math/rand"
)

// [FIXME]:rename JanusGateway

type JanusGateway struct {
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

func NewJanusGateway(n int) *JanusGateway {
	return &JanusGateway{
		Transaction: randomString(n),
	}
}

func (g *JanusGateway) Send(url string, v *JanusRequest) []byte {
	fmt.Println(v)
	body, _ := json.Marshal(v)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		panic("")
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)

	if err != nil {
		panic("")
	}
	return body
}

func (g *JanusGateway) NewSession() *JanusSession {
	req := &JanusRequest{
		Janus:       "create",
		Transaction: g.Transaction,
	}
	body := g.Send(BaseURL, req)

	// fmt.Println(string(body))
	var response *JanusResponse
	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}

	return &JanusSession{
		ID:      response.Data.ID,
		gateway: g,
	}
}
