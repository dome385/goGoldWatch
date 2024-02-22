package main

import (
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var jsonToReturn = `
{
	"ts": 1708625784823,
	"tsj": 1708625776013,
	"date": "Feb 22nd 2024, 01:16:16 pm NY",
	"items": [
	  {
		"curr": "USD",
		"xauPrice": 2020.6,
		"xagPrice": 22.7755,
		"chgXau": -5.22,
		"chgXag": -0.137,
		"pcXau": -0.2577,
		"pcXag": -0.5979,
		"xauClose": 2025.82,
		"xagClose": 22.9125
	  }
	]
  }
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
