// Package eway graphs estimated travelling times along Singapore expressways from Land Transport Authority data.
package eway

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type estTravelTime struct {
	URL   string        `json:"odata.metadata"`
	Value []segmentTime `json:"value"`
}

type segmentTime struct {
	ExpresswayName string      `json:"Name"`
	Direction      json.Number `json:"Direction"`
	EndPoint       string      `json:"FarEndPoint"`
	Start          string      `json:"StartPoint"`
	End            string      `json:"EndPoint"`
	TimeMinutes    json.Number `json:"EstTime"`
}

func (e estTravelTime) String() string {
	s := "Estimated Travel Times along Singapore Expressways:\n"
	for _, v := range e.Value {
		s += fmt.Sprintf("%s\n", v)
	}
	return s
}

func (s segmentTime) String() string {
	return fmt.Sprintf("%s towards %s: %s to %s: %v minutes", s.ExpresswayName, s.EndPoint, s.Start, s.End, s.TimeMinutes)
}

func load(r io.Reader) estTravelTime {
	en := json.NewDecoder(r)
	var tTime estTravelTime
	if err := en.Decode(&tTime); err != nil {
		log.Fatal(err)
	}
	return tTime
}
