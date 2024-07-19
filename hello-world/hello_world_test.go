package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "Hello, World!"
	observed := HelloWorld()

	if observed != expected {
		t.Fatalf("HelloWorld() = %v, want %v", observed, expected)
	}
}
