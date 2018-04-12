package main

import (
	"encoding/json"
)

type JanusResponse struct {
	Janus       string `json:"janus"`
	Transaction string `json:"transaction"`
	Data        `json:"data"`
	// JanusError  `json:"error"`
}

type Data struct {
	ID int `json:"id"`
}

type JanusError struct {
	Code   string `json:"code"`
	Reason string `json:"reason"`
}

func (r *JanusResponse) UnmarshalJSON(data []byte) error {
	var i interface{}
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	m := i.(map[string]interface{})
	r.Janus = m["janus"].(string)

	d := m["data"] //if has key data
	v := d.(map[string]interface{})

	//if has key error

	r.Data.ID = int(v["id"].(float64))
	return nil
}
