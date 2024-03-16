package main

import "testing"

func TestConfig_getHoldings(t *testing.T) {
	_, err := testApp.currentHoldings()
	if err != nil {
		t.Error("failed to get current holdings from database:", err)
	}
}
