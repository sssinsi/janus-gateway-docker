package main

import (
	"encoding/json"
	"strconv"
)

type JanusSession struct {
	ID      int `json:"id"`
	gateway *JanusGateway
}

func (s *JanusSession) Send(v *JanusRequest) []byte {
	url := BaseURL + "/" + strconv.Itoa(s.ID)
	return s.gateway.Send(url, v)
}

func (s *JanusSession) NewRoom() *JanusHandle {
	plugin := "janus.plugin.videoroom"
	return s.Attach(plugin)
}

func (s *JanusSession) Attach(plugin string) *JanusHandle {
	v := &JanusRequest{
		Janus:       "attach",
		Transaction: randomString(12),
		Plugin:      plugin,
		OpaqueID:    "videoroomtest-" + randomString(12),
	}

	body := s.Send(v)
	var r *JanusResponse
	if err := json.Unmarshal(body, &r); err != nil {
		panic(err)
	}
	return &JanusHandle{
		ID:      r.Data.ID,
		session: s,
	}
}
