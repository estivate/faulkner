package faulkner

import (
	"bytes"
	"regexp"
	"testing"
)

// Test logger creation, option passing and banner printing
func TestBanner(t *testing.T) {
	wants := "--------------------------\nTest\n--------------------------\n"
	message := "Test"
	var buf bytes.Buffer
	log_buffer := SetBuffer(&buf)
	logger, _ := NewLogger(log_buffer)
	logger.PrintBanner(message)
	if buf.String() != wants {
		t.Errorf("Banner wants %s, got %v", wants, buf.String())
	}
}

// Test Info Messages
func TestInfoMessage(t *testing.T) {
	message := "Test"
	want := regexp.MustCompile(`\b` + message + `\b`)
	var buf bytes.Buffer
	log_buffer := SetBuffer(&buf)
	logger, _ := NewLogger(log_buffer)
	logger.Info.Println(message)
	if !want.MatchString(buf.String()) {
		t.Errorf("Info Message wants %#q, got %s", want, buf.String())
	}
}
