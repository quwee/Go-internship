package hello

import "testing"

func TestHello(t *testing.T) {
	expect := "Hello World!!!"
	got := Hello()

	if expect != got {
		t.Fatalf("Error: expect: %s, got: %s", expect, got)
	}
}
